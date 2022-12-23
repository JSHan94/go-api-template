package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	"github.com/initia-labs/initia-apis/router"
	"github.com/stretchr/testify/require"
)

var engine *gin.Engine

func init() {
	config := gconfig.Config()
	gin.SetMode(gin.TestMode)
	engine, _ = router.SetupRouter(config)
}

func TestGetBlock(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/1", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	res := w.Body.String()

	var block = &gmodel.CollectedBlock{}
	err := json.Unmarshal([]byte(res), block)
	require.NoError(t, err)
	require.NotEqual(t, block, &gmodel.CollectedBlock{})
}

func TestGetBlockWithWorngHeight(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/block/0", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
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

func TestGetTxByHash(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/tx/D0A96C825A79F0552E4D8CFF7F1A64796E636BB8FD846C5A9D7F57953E21B278", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetTxsByAccount(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/account/init1hk0asaef9nxvnj7gjwawv0zz0yd7adcysktpqu", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetTxsByHeight(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/height/2518", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetTxsByOffset(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/txs/offset/1", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}
