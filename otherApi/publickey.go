package otherapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DataPublicKey struct {
	Data Data `json:"data"`
}
type Data struct {
	Key string `json:"key"`
}

func GetPublicKey() string {
	response, _ := http.NewRequest(http.MethodGet, Url+"/dev/get_key", nil)
	SetHeader(response)

	res, _ := http.DefaultClient.Do(response)
	responseData, _ := ioutil.ReadAll(res.Body)
	var pubkey DataPublicKey
	json.Unmarshal(responseData, &pubkey)
	return pubkey.Data.Key
}

func HashRsa(username string, pass string) string {
	object, _ := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Key      string `json:"key"`
	}{
		Username: username,
		Password: pass,
		Key:      GetPublicKey(),
	})
	responseBody := bytes.NewBuffer(object)
	response, _ := http.NewRequest(http.MethodPost, "http://rsa.somee.com/api/rsa", responseBody)
	response.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(response)
	dataRespone, _ := ioutil.ReadAll(res.Body)
	return string(dataRespone)
}
