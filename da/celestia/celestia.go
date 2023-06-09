package celestia

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/celestiaorg/go-cnc"
	"github.com/gridironOne/gridmint/da"
	"github.com/gridironOne/gridmint/log"
	"github.com/gridironOne/gridmint/store"
	"github.com/gridironOne/gridmint/types"
	pb "github.com/gridironOne/gridmint/types/pb/gridmint"
)

// DataAvailabilityLayerClient use celestia-node public API.
type DataAvailabilityLayerClient struct {
	client *cnc.Client

	config Config
	logger log.Logger
}

var _ da.DataAvailabilityLayerClient = &DataAvailabilityLayerClient{}
var _ da.BatchRetriever = &DataAvailabilityLayerClient{}

// Config stores Celestia DALC configuration parameters.
type Config struct {
	BaseURL     string        `json:"base_url"`
	Timeout     time.Duration `json:"timeout"`
	Fee         int64         `json:"fee"`
	GasLimit    uint64        `json:"gas_limit"`
	NamespaceID [8]byte       `json:"namespace_id"`
}

// Init initializes DataAvailabilityLayerClient instance.
func (c *DataAvailabilityLayerClient) Init(config []byte, kvStore store.KVStore, logger log.Logger) error {
	c.logger = logger

	if len(config) > 0 {
		return json.Unmarshal(config, &c.config)
	}

	return nil
}

// Start prepares DataAvailabilityLayerClient to work.
func (c *DataAvailabilityLayerClient) Start() error {
	c.logger.Info("starting Celestia Data Availability Layer Client", "baseURL", c.config.BaseURL)
	var err error
	//TODO: add * time.Second for the timeout
	c.client, err = cnc.NewClient(c.config.BaseURL, cnc.WithTimeout(c.config.Timeout))
	return err
}

// Stop stops DataAvailabilityLayerClient.
func (c *DataAvailabilityLayerClient) Stop() error {
	c.logger.Info("stopping Celestia Data Availability Layer Client")
	return nil
}

// GetClientType returns client type.
func (c *DataAvailabilityLayerClient) GetClientType() da.Client {
	return da.Celestia
}

// SubmitBatch submits a block to DA layer.
func (c *DataAvailabilityLayerClient) SubmitBatch(batch *types.Batch) da.ResultSubmitBatch {
	blob, err := batch.MarshalBinary()
	if err != nil {
		return da.ResultSubmitBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
			},
		}
	}
	c.logger.Debug("Submitting to da blob with size", "size", len(blob))
	txResponse, err := c.client.SubmitPFD(context.TODO(), c.config.NamespaceID, blob, c.config.Fee, c.config.GasLimit)
	c.logger.Debug("DA submit batch response", "txResponse", txResponse, "error", err)
	if err != nil {
		return da.ResultSubmitBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
			},
		}
	}

	if txResponse.Code != 0 {
		return da.ResultSubmitBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: fmt.Sprintf("Codespace: '%s', Code: %d, Message: %s", txResponse.Codespace, txResponse.Code, txResponse.RawLog),
			},
		}
	}

	return da.ResultSubmitBatch{
		BaseResult: da.BaseResult{
			Code:     da.StatusSuccess,
			Message:  "tx hash: " + txResponse.TxHash,
			DAHeight: uint64(txResponse.Height),
		},
	}
}

// CheckBatchAvailability queries DA layer to check data availability of block at given height.
func (c *DataAvailabilityLayerClient) CheckBatchAvailability(dataLayerHeight uint64) da.ResultCheckBatch {
	shares, err := c.client.NamespacedShares(context.TODO(), c.config.NamespaceID, dataLayerHeight)
	if err != nil {
		return da.ResultCheckBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
			},
		}
	}

	return da.ResultCheckBatch{
		BaseResult: da.BaseResult{
			Code:     da.StatusSuccess,
			DAHeight: dataLayerHeight,
		},
		DataAvailable: len(shares) > 0,
	}
}

// RetrieveBatches gets a batch of blocks from DA layer.
func (c *DataAvailabilityLayerClient) RetrieveBatches(dataLayerHeight uint64) da.ResultRetrieveBatch {
	data, err := c.client.NamespacedData(context.TODO(), c.config.NamespaceID, dataLayerHeight)
	if err != nil {
		return da.ResultRetrieveBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
			},
		}
	}

	batches := make([]*types.Batch, len(data))
	for i, msg := range data {
		var batch pb.Batch
		err = proto.Unmarshal(msg, &batch)
		if err != nil {
			c.logger.Error("failed to unmarshal batch", "daHeight", dataLayerHeight, "position", i, "error", err)
			continue
		}
		batches[i] = new(types.Batch)
		err := batches[i].FromProto(&batch)
		if err != nil {
			return da.ResultRetrieveBatch{
				BaseResult: da.BaseResult{
					Code:    da.StatusError,
					Message: err.Error(),
				},
			}
		}
	}

	return da.ResultRetrieveBatch{
		BaseResult: da.BaseResult{
			Code:     da.StatusSuccess,
			DAHeight: dataLayerHeight,
		},
		Batches: batches,
	}
}
