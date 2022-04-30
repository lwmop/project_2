package client

import (
	"Project_2/models"
	"bytes"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func RpcService(TextInput models.TextInput) (outPut []models.WordFrequencyStruct) {
	url := "http://127.0.0.1:8080/api/text"
	body, _ := jsoniter.Marshal(TextInput)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		logrus.Errorf("RpcService: error in calling service: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logrus.Errorf("RpcService: error in calling service: %v", err)
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("RpcService: error in closing response body: %v", err)
		return
	}

	outputDataINJSONData := new(models.WordFrequencyStructData)
	err = json.Unmarshal(body, &outputDataINJSONData)
	return outputDataINJSONData.Dat

}
