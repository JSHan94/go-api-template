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

func DecodeOne(hit interface{}, result interface{}) error {
	return Decode(hit.(map[string]interface{})["_source"], result)
}

func DecodeMany(hits []interface{}, result interface{}) error {
	for _, hit := range hits {
		err := Decode(hit.(map[string]interface{})["_source"], result)
		if err != nil {
			return err
		}
	}
	return Decode(hits, result)
}
