package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	"github.com/initia-labs/initia-apis/router"
	"github.com/stretchr/testify/require"
)

var engine *gin.Engine

const TX_HASH = "D0A96C825A79F0552E4D8CFF7F1A64796E636BB8FD846C5A9D7F57953E21B278"

const BLOCK_HASH = "18DDA37E666EA7F9E39B32072652C4AF453EFD08BDD7FEE19E5D942B94D4DAF0"
const TIME = "2022-11-08T02:47:10"
const START_TIME = "0001-01-01T00:00:00"
const END_TIME = "2022-11-18T02:47:10"
const ACCOUNT = "init1hk0asaef9nxvnj7gjwawv0zz0yd7adcysktpqu"

const HEIGHT_WITH_TX = "2518"
const HEIGHT_WITHOUT_TX = "1"

type testCase struct {
	name       string
	uri        string
	statusCode int
}

func init() {
	config := gconfig.Config()
	gin.SetMode(gin.TestMode)
	engine, _ = router.SetupRouter(config)
}

// block.go
func TestGetBlock(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		// Test GetBlockByHeight
		{"get block height", "/v1/block/height/1", http.StatusOK},
		{"get 0 block height", "/v1/block/height/0", http.StatusBadRequest},
		{"get -1 block height", "/v1/block/height/-1", http.StatusBadRequest},
		{"get wrong block height", "/v1/block/height/abc", http.StatusBadRequest},
		// Test GetBlockLatest
		{"get latest block", "/v1/block/latest", http.StatusOK},
		// Test GetBlockByTime
		{"get block by time", "/v1/block/time/" + TIME, http.StatusOK},
		// Test GetBlockByHash
		{"get block by hash", "/v1/block/hash/" + BLOCK_HASH, http.StatusOK},
		// Test GetBlockAvgTime
		{"get block avgtime", "/v1/block/avgtime/100", http.StatusOK},
		// Test GetBlocks
		{"get blocks with range", "/v1/blocks/1/100", http.StatusOK},
		{"get blocks with worng range", "/v1/blocks/100/1", http.StatusOK},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tc.uri, nil)
		engine.ServeHTTP(w, req)
		require.Equal(t, tc.statusCode, w.Code, "failed to test : "+tc.name)
	}
}

// tx.go
func TestGetTx(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		// Test GetTxByHash
		{"get tx by hash", "/v1/tx/" + TX_HASH, http.StatusOK},
		// Test GetTxsByAccount
		{"get txs by account", "/v1/txs/account/" + ACCOUNT, http.StatusOK},
		// Test GetTxsByHeight
		{"get txs by height", "/v1/txs/height/" + HEIGHT_WITH_TX, http.StatusOK},
		{"get txs by height without txs", "/v1/txs/height/" + HEIGHT_WITHOUT_TX, http.StatusOK},
		// Test GetTxsByOffset
		{"get txs by offset 3", "/v1/txs/offset/3", http.StatusOK},
		{"get txs by offset 0", "/v1/txs/offset/0", http.StatusOK},
		{"get mempool txs", "/v1/txs/mempool", http.StatusOK},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tc.uri, nil)
		engine.ServeHTTP(w, req)
		require.Equal(t, tc.statusCode, w.Code, "failed to test : "+tc.name)
	}
}

// analysis.go
func TestGetAnalysis(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		// Test GetTxVolume
		{"get tx volume", "/v1/analysis/tx-volume/" + START_TIME + "/" + END_TIME, http.StatusOK},
		// Test GetTxCount
		{"get tx count", "/v1/analysis/tx-count/" + START_TIME + "/" + END_TIME, http.StatusOK},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tc.uri, nil)
		engine.ServeHTTP(w, req)
		require.Equal(t, tc.statusCode, w.Code, "failed to test : "+tc.name)
	}
}
