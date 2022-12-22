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
	req, _ := http.NewRequest("GET", "/v1/block?height=1", nil)
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
	req, _ := http.NewRequest("GET", "/v1/block?height=0", nil)
	engine.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}
