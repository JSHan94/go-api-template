package handler

import (
	"github.com/gin-gonic/gin"
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
)

func GetTxVolume(c *gin.Context, indexName string) (*gmodel.TxVolume, error) {
	panic("not implemented yet")
	// client := gdatabase.GetClient()
	// query, err := GetTxVolumeQuery(c)
	// if err != nil {
	// 	return nil, err
	// }
	// hits, err := client.Search(indexName, query)
	// if err != nil {
	// 	return nil, err
	// }

	// // find all events with uinit and filter them
	// txVolume := &gmodel.TxVolume{}
	// txVolume.Denom = gconfig.NATIVE_TOKEN_SYMBOL
	// txVolume.From = c.Param("from")
	// txVolume.To = c.Param("to")

	// return txVolume, nil
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
	txCount.From = c.Param("from")
	txCount.To = c.Param("to")
	txCount.Value = len(hits)

	return txCount, nil
}
