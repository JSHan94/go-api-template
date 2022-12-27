package lib

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func decode(source interface{}, result interface{}) error {
	sourceString, err := json.Marshal(source)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return json.Unmarshal(sourceString, &result)
}

func Decode(hit interface{}, result interface{}) error {
	source := hit.(map[string]interface{})["_source"]
	if len(source.(map[string]interface{})) == 0 {
		return nil
	}
	return decode(source, result)
}
