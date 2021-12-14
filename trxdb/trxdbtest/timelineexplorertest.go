// Copyright 2021 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trxdbtest

import (
	"context"
	"testing"
	"time"

	"github.com/streamingfast/kvdb"
	ct "github.com/streamingfast/sf-ethereum/codec/testing"
	pbcodec "github.com/streamingfast/sf-ethereum/pb/sf/ethereum/codec/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var timelineExplorerTests = []testFunc{
	TestBlockIDAt,
	TestBlockIDAfter,
	TestBlockIDBefore,
}

func TestAllTimelineExplorer(t *testing.T, driverName string, driverFactory DriverFactory) {
	for _, rt := range timelineExplorerTests {
		t.Run(driverName+"/"+getFunctionName(rt), func(t *testing.T) {
			rt(t, driverFactory)
		})
	}
}

func TestBlockIDAt(t *testing.T, driverFactory DriverFactory) {
	noon := time.Date(2020, time.February, 02, 12, 0, 0, 0, time.UTC)
	twopm := time.Date(2020, time.February, 02, 14, 0, 0, 0, time.UTC)
	fourpm := time.Date(2020, time.February, 02, 16, 0, 0, 0, time.UTC)

	tests := []struct {
		name            string
		blocks          []*pbcodec.Block
		time            time.Time
		expectBlockHash string
		expectErr       error
	}{
		{
			name: "sunny path",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			time:            noon,
			expectBlockHash: "00000008aa",
		},
		{
			name: "no block that matches",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			time:      fourpm,
			expectErr: kvdb.ErrNotFound,
		},
		{
			name:      "no blocks",
			blocks:    []*pbcodec.Block{},
			time:      fourpm,
			expectErr: kvdb.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ctx = context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, blk := range test.blocks {
				require.NoError(t, db.PutBlock(ctx, blk))
				require.NoError(t, db.UpdateNowIrreversibleBlock(ctx, blk))
			}
			require.NoError(t, db.Flush(ctx))

			id, err := db.BlockIDAt(ctx, test.time)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectBlockHash, id)
			}
		})
	}
}

func TestBlockIDAfter(t *testing.T, driverFactory DriverFactory) {
	noon := time.Date(2020, time.February, 02, 12, 0, 0, 0, time.UTC)
	twopm := time.Date(2020, time.February, 02, 14, 0, 0, 0, time.UTC)
	fourpm := time.Date(2020, time.February, 02, 16, 0, 0, 0, time.UTC)

	tests := []struct {
		name            string
		blocks          []*pbcodec.Block
		time            time.Time
		inclusive       bool
		expectBlockHash string
		expectTime      time.Time
		expectErr       error
	}{
		{
			name: "sunny path",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(fourpm),
					},
				},
			},
			time:            twopm,
			expectTime:      fourpm,
			expectBlockHash: "00000003aa",
		},
		{
			name: "no block that matches",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			time:      fourpm,
			expectErr: kvdb.ErrNotFound,
		},
		{
			name: "should not match block when not inclusive",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			inclusive: false,
			time:      twopm,
			expectErr: kvdb.ErrNotFound,
		},
		{
			name: "should  match block when inclusive",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			inclusive:       true,
			time:            twopm,
			expectTime:      twopm,
			expectBlockHash: "00000003aa",
		},
		{
			name:      "no blocks",
			blocks:    []*pbcodec.Block{},
			time:      fourpm,
			expectErr: kvdb.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ctx = context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, blk := range test.blocks {
				require.NoError(t, db.PutBlock(ctx, blk))
				require.NoError(t, db.UpdateNowIrreversibleBlock(ctx, blk))
			}
			require.NoError(t, db.Flush(ctx))

			id, foundTime, err := db.BlockIDAfter(ctx, test.time, test.inclusive)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectBlockHash, id)
				assert.Equal(t, test.expectTime, foundTime.UTC())
			}
		})
	}
}

func TestBlockIDBefore(t *testing.T, driverFactory DriverFactory) {
	noon := time.Date(2020, time.February, 02, 12, 0, 0, 0, time.UTC)
	twopm := time.Date(2020, time.February, 02, 14, 0, 0, 0, time.UTC)
	fourpm := time.Date(2020, time.February, 02, 16, 0, 0, 0, time.UTC)

	tests := []struct {
		name            string
		blocks          []*pbcodec.Block
		time            time.Time
		inclusive       bool
		expectBlockHash string
		expectTime      time.Time
		expectErr       error
	}{
		{
			name: "no block that matches",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(fourpm),
					},
				},
			},
			time:      noon,
			expectErr: kvdb.ErrNotFound,
		},
		{
			name: "sunny path",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(fourpm),
					},
				},
			},
			time:            twopm,
			expectTime:      noon,
			expectBlockHash: "00000008aa",
		},
		{
			name: "should not match block when not inclusive",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000009aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			inclusive: false,
			time:      noon,
			expectErr: kvdb.ErrNotFound,
		},
		{
			name: "should match block when inclusive",
			blocks: []*pbcodec.Block{
				{
					Hash:   ct.Hash("00000008aa").Bytes(t),
					Number: 8,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(noon),
					},
				},
				{
					Hash:   ct.Hash("00000003aa").Bytes(t),
					Number: 3,
					Header: &pbcodec.BlockHeader{
						Timestamp: ct.ToTimestamp(twopm),
					},
				},
			},
			inclusive:       true,
			time:            noon,
			expectTime:      noon,
			expectBlockHash: "00000008aa",
		},
		{
			name:      "no blocks",
			blocks:    []*pbcodec.Block{},
			time:      fourpm,
			expectErr: kvdb.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ctx = context.Background()
			db, clean := driverFactory()
			defer clean()

			for _, blk := range test.blocks {
				require.NoError(t, db.PutBlock(ctx, blk))
				require.NoError(t, db.UpdateNowIrreversibleBlock(ctx, blk))
			}
			require.NoError(t, db.Flush(ctx))

			id, foundTime, err := db.BlockIDBefore(ctx, test.time, test.inclusive)

			if test.expectErr != nil {
				assert.Equal(t, test.expectErr, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectBlockHash, id)
				assert.Equal(t, test.expectTime, foundTime.UTC())
			}
		})
	}
}