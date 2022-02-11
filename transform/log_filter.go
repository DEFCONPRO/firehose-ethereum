package transform

import (
	"bytes"
	"fmt"
	"github.com/streamingfast/dstore"

	"github.com/streamingfast/eth-go"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/streamingfast/bstream"
	"github.com/streamingfast/bstream/transform"
	pbcodec "github.com/streamingfast/sf-ethereum/pb/sf/ethereum/codec/v1"
	pbtransforms "github.com/streamingfast/sf-ethereum/pb/sf/ethereum/transforms/v1"
	"google.golang.org/protobuf/proto"
)

var LogFilterMessageName = proto.MessageName(&pbtransforms.BasicLogFilter{})
var LogFilterIndexStore dstore.Store

func BasicLogFilterFactory(indexStore dstore.Store, possibleIndexSizes []uint64) *transform.Factory {
	return &transform.Factory{
		Obj: &pbtransforms.BasicLogFilter{},
		NewFunc: func(message *anypb.Any) (transform.Transform, error) {
			mname := message.MessageName()
			if mname != LogFilterMessageName {
				return nil, fmt.Errorf("expected type url %q, recevied %q ", LogFilterMessageName, message.TypeUrl)
			}

			filter := &pbtransforms.BasicLogFilter{}
			err := proto.Unmarshal(message.Value, filter)
			if err != nil {
				return nil, fmt.Errorf("unexpected unmarshall error: %w", err)
			}

			if len(filter.Addresses) == 0 && len(filter.EventSignatures) == 0 {
				return nil, fmt.Errorf("a log filter transform requires at-least one address or one event signature")
			}

			f := &BasicLogFilter{
				indexStore:         indexStore,
				possibleIndexSizes: possibleIndexSizes,
			}

			for _, addr := range filter.Addresses {
				f.Addresses = append(f.Addresses, addr)
			}
			for _, sig := range filter.EventSignatures {
				f.EventSigntures = append(f.EventSigntures, sig)
			}

			return f, nil
		},
	}
}

type BasicLogFilter struct {
	Addresses      []eth.Address
	EventSigntures []eth.Hash

	indexStore         dstore.Store
	possibleIndexSizes []uint64
}

func (p *BasicLogFilter) matchAddress(src eth.Address) bool {
	if len(p.Addresses) == 0 {
		return true
	}
	for _, addr := range p.Addresses {
		if bytes.Equal(addr, src) {
			return true
		}
	}
	return false
}

func (p *BasicLogFilter) matchEventSignature(src eth.Hash) bool {
	if len(p.EventSigntures) == 0 {
		return true
	}
	for _, topic := range p.EventSigntures {
		if bytes.Equal(topic, src) {
			return true
		}
	}
	return false
}

func (p *BasicLogFilter) Transform(readOnlyBlk *bstream.Block, in transform.Input) (transform.Output, error) {
	ethBlock := readOnlyBlk.ToProtocol().(*pbcodec.Block)
	traces := []*pbcodec.TransactionTrace{}
	for _, trace := range ethBlock.TransactionTraces {
		match := false
		for _, log := range trace.Receipt.Logs {
			if p.matchAddress(log.Address) && p.matchEventSignature(log.Topics[0]) {
				match = true
				break
			}
		}
		if match {
			traces = append(traces, trace)
		}
	}
	ethBlock.TransactionTraces = traces
	return ethBlock, nil
}

// GetIndexProvider will instantiate a new LogAddressIndex conforming to the bstream.BlockIndexProvider interface
func (p *BasicLogFilter) GetIndexProvider() bstream.BlockIndexProvider {
	if p.indexStore == nil {
		return nil
	}

	if len(p.Addresses) == 0 && len(p.EventSigntures) == 0 {
		return nil
	}

	return NewLogAddressIndexProvider(
		p.indexStore,
		p.Addresses,
		p.EventSigntures,
		p.possibleIndexSizes,
	)
}
