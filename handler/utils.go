package handler

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func Decode(source interface{}, result interface{}) error {
	sourceString, err := json.Marshal(source)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return json.Unmarshal(sourceString, &result)
}

func DecodeOne(hits []interface{}, result interface{}) error {
	return Decode(hits[0].(map[string]interface{})["_source"], result)
}

func DecodeMany(hits []interface{}, result interface{}) error {
	return Decode(hits, result)
}
