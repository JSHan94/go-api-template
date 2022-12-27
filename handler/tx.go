package handler

import (
	"github.com/gin-gonic/gin"
	gdatabase "github.com/initia-labs/initia-apis/database"
	gmodel "github.com/initia-labs/initia-apis/database/model"
	glib "github.com/initia-labs/initia-apis/lib"
)

func GetTxs(indexName string, query string) (*gmodel.CollectedTxs, error) {
	client := gdatabase.GetClient()
	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}
	collectedTxs := &gmodel.CollectedTxs{}
	for _, hit := range hits {
		collectedTx := &gmodel.CollectedTx{}
		err = glib.Decode(hit, collectedTx)
		if err != nil {
			return nil, err
		}
		collectedTxs.Txs = append(collectedTxs.Txs, collectedTx)
	}

	return collectedTxs, err
}

func GetTxByHash(c *gin.Context, indexName string) (*gmodel.CollectedTx, error) {
	client := gdatabase.GetClient()
	query, err := GetTxByHashQuery(c)
	if err != nil {
		return nil, err
	}
	hits, err := client.Search(indexName, query)
	if err != nil {
		return nil, err
	}

	collectedTx := &gmodel.CollectedTx{}
	err = glib.Decode(hits[0], collectedTx)

	return collectedTx, err
}

func GetTxsByAccount(c *gin.Context, indexName string) (*gmodel.CollectedTxs, error) {
	query, err := GetTxsByAccountQuery(c)
	if err != nil {
		return nil, err
	}

	txs, err := GetTxs(indexName, query)
	if err != nil {
		return nil, err
	}
	txs.Limit = c.Param("limit")
	return txs, err
}

func GetTxsByHeight(c *gin.Context, indexName string) (*gmodel.CollectedTxs, error) {
	query, err := GetTxsByHeightQuery(c)
	if err != nil {
		return nil, err
	}

	txs, err := GetTxs(indexName, query)
	if err != nil {
		return nil, err
	}
	txs.Limit = c.Param("limit")
	return txs, err
}

func GetTxsByOffset(c *gin.Context, indexName string) (*gmodel.CollectedTxs, error) {
	query, err := GetTxsByOffsetQuery(c)

	if err != nil {
		return nil, err
	}
	txs, err := GetTxs(indexName, query)
	if err != nil {
		return nil, err
	}
	txs.Limit = c.Param("limit")
	txs.Offset = c.Param("offset")
	return txs, err
}
