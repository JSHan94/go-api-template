package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	glib "github.com/initia-labs/initia-apis/lib"
)

func GetTxVolume(c *gin.Context, indexName string) (*gmodel.TxVolume, error) {
	client := gdatabase.GetClient()
	query, err := GetTxVolumeQuery(c)
	if err != nil {
		return nil, err
	}
	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	txVolume := &gmodel.TxVolume{}
	txVolume.Denom = gconfig.NATIVE_TOKEN_SYMBOL
	txVolume.Start = c.Param("start")
	txVolume.End = c.Param("end")

	for _, hit := range hits {
		collectedEvents := &gmodel.CollectedEvents{}
		err = glib.Decode(hit, collectedEvents)
		if err != nil {
			return nil, err
		}

		for _, attr := range collectedEvents.Attributes {
			if attr.Key == "amount" && strings.HasSuffix(attr.Value, "uinit") {
				volume, err := strconv.Atoi(strings.TrimSuffix(attr.Value, "uinit"))
				if err != nil {
					return nil, err
				}
				txVolume.Value += volume
			}
		}
	}
	return txVolume, nil
}

func GetTxCount(c *gin.Context, indexName string) (*gmodel.TxCount, error) {
	client := gdatabase.GetClient()
	query, err := GetTxVolumeQuery(c)
	if err != nil {
		return nil, err
	}

	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	txCount := &gmodel.TxCount{}
	txCount.Start = c.Param("start")
	txCount.End = c.Param("end")
	txCount.Value = len(hits)

	return txCount, nil
}
