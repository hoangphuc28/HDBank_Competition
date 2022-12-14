package otherapi

import (
	"bytes"
	"encoding/json"
	usermodel "github.com/Zhoangp/HDBank_Competition/internal/user/model"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

type BodyRequest struct {
	Data    DataRequest `json:"data"`
	Request Request     `json:"request"`
}
type DataRequest struct {
	Credential     string `json:"credential"`
	Email          string `json:"email"`
	FullName       string `json:"fullName"`
	IdentityNumber string `json:"identityNumber"`
	Key            string `json:"key"`
	Phone          string `json:"phone"`
}
type Request struct {
	RequestID   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
}

type DataResponse struct {
	Response struct {
		ResponseID      string `json:"responseId"`
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		ResponseTime    string `json:"responseTime"`
	} `json:"response"`
	Data struct {
		UserID string `json:"userId"`
	} `json:"data"`
}

func Register(user usermodel.User) DataResponse {
	a := DataRequest{Credential: HashRsa(user.UserName, user.Password),
		Email:          user.Email,
		FullName:       user.FullName,
		IdentityNumber: user.IdentityNumber,
		Key:            GetPublicKey(),
		Phone:          user.Phone,
	}
	newUUID, _ := exec.Command("uuidgen").Output()
	b := Request{RequestID: string(newUUID),
		RequestTime: time.Now().Format("20060102150405")}
	bodyData := BodyRequest{a, b}
	data, _ := json.Marshal(&bodyData)
	responseBody := bytes.NewBuffer(data)
	response, _ := http.NewRequest(http.MethodPost, "https://7ucpp7lkyl.execute-api.ap-southeast-1.amazonaws.com/dev/register", responseBody)
	SetHeader(response)
	res, _ := http.DefaultClient.Do(response)
	dataRespone, _ := ioutil.ReadAll(res.Body)
	var x DataResponse
	json.Unmarshal(dataRespone, &x)
	return x
}
