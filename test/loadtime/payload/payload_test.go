package payload_test

import (
	"bytes"
	"testing"

	"github.com/gridironOne/gridmint/test/loadtime/payload"
	"github.com/gridironOne/gridmint/test/pb/loadtime"
	"github.com/google/uuid"
)

const payloadSizeTarget = 1024 // 1kb

func TestSize(t *testing.T) {
	s, err := payload.MaxUnpaddedSize()
	if err != nil {
		t.Fatalf("calculating max unpadded size %s", err)
	}
	if s > payloadSizeTarget {
		t.Fatalf("unpadded payload size %d exceeds target %d", s, payloadSizeTarget)
	}
}

func TestRoundTrip(t *testing.T) {
	const (
		testConns = 512
		testRate  = 4
	)
	testID := [16]byte(uuid.New())
	b, err := payload.NewBytes(&loadtime.Payload{
		Size_:       payloadSizeTarget,
		Connections: testConns,
		Rate:        testRate,
		Id:          testID[:],
	})
	if err != nil {
		t.Fatalf("generating payload: %s", err)
	}
	if len(b) < payloadSizeTarget {
		t.Fatalf("payload size in bytes %d less than expected %d", len(b), payloadSizeTarget)
	}
	p, err := payload.FromBytes(b)
	if err != nil {
		t.Fatalf("reading payload %s", err)
	}
	if p.GetSize_() != payloadSizeTarget {
		t.Fatalf("payload size value %d does not match expected %d", p.GetSize_(), payloadSizeTarget)
	}
	if p.Connections != testConns {
		t.Fatalf("payload connections value %d does not match expected %d", p.Connections, testConns)
	}
	if p.Rate != testRate {
		t.Fatalf("payload rate value %d does not match expected %d", p.Rate, testRate)
	}
	if !bytes.Equal(p.Id, testID[:]) {
		t.Fatalf("payload ID value %d does not match expected %d", p.Id, testID)
	}
}
