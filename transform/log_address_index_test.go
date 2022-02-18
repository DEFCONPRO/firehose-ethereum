package transform

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/streamingfast/dstore"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/streamingfast/bstream"
	"github.com/streamingfast/eth-go"
	pbcodec "github.com/streamingfast/sf-ethereum/pb/sf/ethereum/codec/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogAddressIndexer(t *testing.T) {
	tests := []struct {
		name                       string
		blocks                     []*pbcodec.Block
		indexSize                  uint64
		shouldWriteFile            bool
		shouldReadFile             bool
		expectAddressesAfterWrite  map[string][]uint64
		expectSignaturesAfterWrite map[string][]uint64
		expectAddressesAfterRead   map[string][]uint64
		expectSignaturesAfterRead  map[string][]uint64
	}{
		{
			name:            "sunny within bounds",
			indexSize:       10,
			shouldWriteFile: false,
			shouldReadFile:  false,
			blocks:          testEthBlocks(t, 2),
			expectAddressesAfterWrite: map[string][]uint64{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {10, 11},
				"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {10, 11},
				"cccccccccccccccccccccccccccccccccccccccc": {10},
				"dddddddddddddddddddddddddddddddddddddddd": {11},
			},
			expectSignaturesAfterWrite: map[string][]uint64{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {10, 11},
				"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {10, 11},
				"cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc": {10},
				"dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd": {11},
			},
		},
		{
			name:            "sunny and we wrote an index",
			indexSize:       2,
			shouldWriteFile: true,
			shouldReadFile:  true,
			blocks:          testEthBlocks(t, 3),
			expectAddressesAfterWrite: map[string][]uint64{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {12},
				"1111111111111111111111111111111111111111": {12},
				"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {12},
			},
			expectSignaturesAfterWrite: map[string][]uint64{
				"0000000000000000000000000000000000000000000000000000000000000000": {12},
				"1111111111111111111111111111111111111111111111111111111111111111": {12},
				"2222222222222222222222222222222222222222222222222222222222222222": {12},
			},
			expectAddressesAfterRead: map[string][]uint64{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {10, 11},
				"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {10, 11},
				"cccccccccccccccccccccccccccccccccccccccc": {10},
				"dddddddddddddddddddddddddddddddddddddddd": {11},
			},
			expectSignaturesAfterRead: map[string][]uint64{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {10, 11},
				"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {10, 11},
				"cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc": {10},
				"dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd": {11},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			results := make(map[string][]byte)

			// spawn an indexStore which will populate the results
			indexStore := dstore.NewMockStore(func(base string, f io.Reader) error {
				if test.shouldWriteFile {
					content, err := ioutil.ReadAll(f)
					require.NoError(t, err)
					results[base] = content
				}
				return nil
			})

			// spawn an indexer with our mock indexStore
			indexer := NewLogAddressIndexer(indexStore, test.indexSize)
			for _, blk := range test.blocks {
				// feed the indexer
				err := indexer.ProcessEthBlock(blk)
				require.NoError(t, err)
			}

			// check our addrs
			require.Equal(t, len(test.expectAddressesAfterWrite), len(indexer.currentIndex.addrs))
			for addr, expectMatches := range test.expectAddressesAfterWrite {
				m, ok := indexer.currentIndex.addrs[addr]
				require.True(t, ok)
				arr := m.ToArray()
				require.Equal(t, expectMatches, arr)
			}

			// check our sigs
			assert.Equal(t, len(test.expectSignaturesAfterWrite), len(indexer.currentIndex.eventSigs))
			for sig, expectMatches := range test.expectSignaturesAfterWrite {
				m, ok := indexer.currentIndex.eventSigs[sig]
				require.True(t, ok)
				arr := m.ToArray()
				require.Equal(t, expectMatches, arr)
			}

			if test.shouldReadFile {
				// populate a new indexStore with the prior results
				indexStore = dstore.NewMockStore(nil)
				for indexName, indexContents := range results {
					indexStore.SetFile(indexName, indexContents)
				}

				// spawn a new indexer with the new indexStore
				indexer = NewLogAddressIndexer(indexStore, test.indexSize)
				for indexName, _ := range results {
					// attempt to read back the indexes
					err := indexer.readIndex(indexName)
					require.NoError(t, err)

					// check our addrs
					require.Equal(t, len(test.expectAddressesAfterRead), len(indexer.currentIndex.addrs))
					for addr, expectMatches := range test.expectAddressesAfterRead {
						m, ok := indexer.currentIndex.addrs[addr]
						require.True(t, ok)
						arr := m.ToArray()
						require.Equal(t, expectMatches, arr)
					}

					// check our sigs
					require.Equal(t, len(test.expectSignaturesAfterRead), len(indexer.currentIndex.eventSigs))
					for sig, expectMatches := range test.expectSignaturesAfterRead {
						m, ok := indexer.currentIndex.eventSigs[sig]
						require.True(t, ok)
						arr := m.ToArray()
						require.Equal(t, expectMatches, arr)
					}
				}
			}
		})
	}
}

func TestRoaring_SaveLoad(t *testing.T) {
	r := roaring64.NewBitmap()
	r.Add(1000)
	r.Add(2000)
	r.Add(2005)
	r.Add(20000005)
	r.Add(530000005)

	short, err := r.ToBase64()
	require.NoError(t, err)

	bts, err := r.ToBytes()
	require.NoError(t, err)

	r2 := roaring64.NewBitmap()
	r2.UnmarshalBinary(bts)

	short2, err := r2.ToBase64()
	require.NoError(t, err)

	assert.Equal(t, short, short2)
}

func TestLogAddressIndex_Matching(t *testing.T) {
	tests := []struct {
		name          string
		reqAddresses  []string
		reqSignatures []string
		expectBlocks  []uint64
	}{
		{
			name: "single address",
			reqAddresses: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			expectBlocks: []uint64{10, 11},
		},
		{
			name: "single address single block",
			reqAddresses: []string{
				"cccccccccccccccccccccccccccccccccccccccc",
			},
			expectBlocks: []uint64{10},
		},
		{
			name: "two addresses",
			reqAddresses: []string{
				"cccccccccccccccccccccccccccccccccccccccc",
				"dddddddddddddddddddddddddddddddddddddddd",
			},
			expectBlocks: []uint64{10, 11},
		},
		{
			name: "duplicate address match",
			reqAddresses: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
				"cccccccccccccccccccccccccccccccccccccccc",
				"dddddddddddddddddddddddddddddddddddddddd",
			},
			expectBlocks: []uint64{10, 11},
		},
		{
			name: "addr and sig",
			reqAddresses: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			reqSignatures: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			expectBlocks: []uint64{10, 11},
		},
		{
			name: "addr and restrictive sig",
			reqAddresses: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			reqSignatures: []string{
				"dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
			},
			expectBlocks: []uint64{11},
		},
		{
			name: "restrictive addr and good sig",
			reqAddresses: []string{
				"dddddddddddddddddddddddddddddddddddddddd",
			},
			reqSignatures: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			expectBlocks: []uint64{11},
		},
		{
			name: "no joining match",
			reqAddresses: []string{
				"cccccccccccccccccccccccccccccccccccccccc",
			},
			reqSignatures: []string{
				"dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
			},
			expectBlocks: nil,
		},
		{
			name: "nothing matches",
			reqAddresses: []string{
				"ff00ffffffffffffffffffffffffffffffffffff",
			},
			reqSignatures: []string{
				"ff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			},
			expectBlocks: nil,
		},
		{
			name: "only signature",
			reqSignatures: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			expectBlocks: []uint64{10, 11},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			accountIndexStore := dstore.NewMockStore(func(base string, f io.Reader) error {
				return nil
			})
			indexer := NewLogAddressIndexer(accountIndexStore, 10)

			testBlocks := []*bstream.Block{
				testBlockFromFiles(t, "blk10.json"),
				testBlockFromFiles(t, "blk11.json"),
			}
			for _, blk := range testBlocks {
				indexer.ProcessEthBlock(blk.ToProtocol().(*pbcodec.Block))
			}

			var reqAddresses []eth.Address
			for _, addr := range test.reqAddresses {
				reqAddresses = append(reqAddresses, eth.MustNewAddress("0x"+addr))
			}

			var reqSignatures []eth.Hash
			for _, sig := range test.reqSignatures {
				reqSignatures = append(reqSignatures, eth.MustNewHash("0x"+sig))
			}

			matching := indexer.currentIndex.matchingBlocks(reqAddresses, reqSignatures)

			assert.Equal(t, test.expectBlocks, matching)
		})
	}

}