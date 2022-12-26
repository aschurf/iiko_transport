package iikotransport

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const iikourl = "https://api-ru.iiko.services/api/"

type IikoRequest struct {
	EndPoint string
	Method   string
	Bearer   string
	Payload  map[string]interface{}
}

func (req IikoRequest) Send() (map[string]interface{}, error) {

	client := &http.Client{}

	postBody, _ := json.Marshal(req.Payload)

	resp, err := http.NewRequest(req.Method, iikourl+req.EndPoint, bytes.NewBuffer(postBody))
	resp.Header.Set("Content-Type", "application/json")
	if req.Bearer != "" {
		resp.Header.Add("Authorization", "Bearer "+req.Bearer)
	}

	res, err := client.Do(resp)
	//Handle Error
	if err != nil {
		return make(map[string]interface{}), err
	}
	defer res.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return make(map[string]interface{}), err
	}

	resBytes := []byte(body)
	var jsonRes map[string]interface{}
	_ = json.Unmarshal(resBytes, &jsonRes)

	return jsonRes, nil
}
