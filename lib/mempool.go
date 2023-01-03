package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	"github.com/sirupsen/logrus"
)

func GetChainURLAndID(c *gin.Context) (string, string, error) {
	cfg := gconfig.GetConfig().Chain
	if c.Query("chainid") == cfg.MainnetChainID || c.Query("chainid") == "" {
		return cfg.PublicMainnetURL, cfg.MainnetChainID, nil
	}
	if c.Query("chainid") == cfg.TestnetChainID {
		return cfg.PublicTestnetURL, cfg.TestnetChainID, nil
	}
	return "", "", ErrWrongChainID
}

func GetUnconfirmedTxs(chainURL string) (*gmodel.UnconfirmedTxsResult, error) {
	response, err := http.Get(chainURL + ":26657/unconfirmed_txs?limit=10000")
	if err != nil {
		logrus.Info(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Info(err)
		return nil, err
	}

	unconfirmedTxs := &gmodel.UnconfirmedTxsResult{}
	json.Unmarshal(body, &unconfirmedTxs)
	return unconfirmedTxs, nil
}
