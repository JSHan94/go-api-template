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
const ACCOUNT = "init1hk0asaef9nxvnj7gjwawv0zz0yd7adcysktpqu"

const HEIGHT_WITH_TX = "2518"
const HEIGHT_WITHOUT_TX = "1"

func init() {
	config := gconfig.Config()
	gin.SetMode(gin.TestMode)
	engine, _ = router.SetupRouter(config)
}

// block.go
func TestGetBlockByHeight(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/height/1", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetBlockByHeightWithWorngHeight(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/height/0", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBlockLatest(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/latest", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetBlockByTime(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/time/"+TIME, nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetBlockByHash(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/hash/"+BLOCK_HASH, nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetBlockAvgTime(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/avgtime?height=100", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetBlocks(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/blocks?from=1&to=100", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

// tx.go
func TestGetTxByHash(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/tx/"+TX_HASH, nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetTxsByAccount(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/account/"+ACCOUNT, nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetTxsByHeight(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/height/"+HEIGHT_WITH_TX, nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetTxsByHeightWithoutTxs(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/height/"+HEIGHT_WITHOUT_TX, nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTxsByOffset(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/offset/1", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

// analysis.go
// func TestGetTxVolume(t *testing.T) {
// 	t.Parallel()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/v1/analysis/tx-volume?from=100&to=200", nil)
// 	engine.ServeHTTP(w, req)
// 	require.Equal(t, http.StatusOK, w.Code)
// }

func TestGetTxCount(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/analysis/tx-count?from=100&to=200", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}
