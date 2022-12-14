package otherapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const Url = "https://7ucpp7lkyl.execute-api.ap-southeast-1.amazonaws.com"

type TokenType struct {
	Id_token     string `json:"id_token"`
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
}

func SetHeader(response *http.Request) {
	response.Header.Add("x-api-key", "hutech_hackathon@123456")
	response.Header.Add("Content-Type", "application/json")
	response.Header.Add("access-token", GetToken())
}

const refeshString = "client_id=sikcnei4t2h3ntkqj5d49ltvr&grant_type=refresh_token&refresh_token=eyJjdHkiOiJKV1QiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiUlNBLU9BRVAifQ.gnwzw0WAWkYcecAWA33f00uhLTk-03RrWvnV5VSEdp2r44T9i8Nuz1CPKN_pd8all9xRbPVGXZo39a203_u1r24PeF-DIA_XTWejqVxTrKuNdv_PAm__lYiQo3MZBtc8NYc58HIlLrMcd9kZTWs_V4eOFbGUqB-1wwzf0FSYpgr72DCOs7JW2rJ3lZcPBQNa2dPPjbXltrml-VHsLoY4STIgNLFncd23Un4TzbaAf1e0CkCvoQMRJzFNM81HfNd6-p4XY2915V4B4RjmdbHD0stnvDGqrJiwHlJir4RTP28CJyqKXvFmDghMJszRmibD2YdzWapMptRi3o0soDJJeA.dhHzMaxEzsozhjUH.Zhhij2Akme88Lmw9xeK04avnNVDVzozf623Wk3ZbbWfJ8boL6Fru5xxuDzquGK0XDnVIX0J_G5mV5lzay1r82RKp2yqzMyup63zu1W22X1u1OdPm9M2Dr5L6QIU3YiThZ1kTCFMJkpD4o0dmzQH8S3vgFCnviKySHsB0sFlh_Wy6DOAGAjAdun9WfKEwcXmGR5H_2-Y-lEmGIXrcAgmbtItxA_C9TbK0bblEj8Pt3-gwqWRcRtiLnSwacj5NJlTgHICafqLGXCC06zKblU_3x8XG91auPukZySYkVzWZ4FiWa5XG1iNrVW3SFiVegAS1-MUkKxSC-qej1CjdrCGT7hQYD9SbpZVQxZkc2rmiUoM_dzXdPKcnRxiYYWgySYU3FGUX8gIwVswSjUIB6NMRuEWs38XJzyPGPKzzdXtsbu0xWI0X5LXBDO5auOuyiojsrY6zGhHfvNdEt1f3_T0S93RdXhvHoXo8cxj0kKrNLZjjJxtx4VE9gdIlS6c2TR7h8lMUwclYVNs26zg2JnjmqbqlpXfoQtojZjbtGLShGihzGaI5y-0uX3RFYF0HZQYINIV8iIOckPPr77T26S3cil2p4CU49zm-XEvmNiu_tbsfHdbbKHldb05HjUUcFmJJ7EA79i7kgE2Epo2U3o12KfsaM61qIjxK4KT0uVv-5FOYL-_1vRJ1886STSt9sHlhzxCV9HjwfGVRaGeVns7GTuoWeb0S2twY-9b0UidCNkMC7t8Y6JWj8mTu4fBhGAecbxAUZ2UOXtpu7h134HRpQThu1ZgAyy6q5djV5Zc2ArYHoyIbU4ZHsUTvboWM8sYtG97w8_dhPX1svJ0beKYGUpqgYM-Gw4Q3uiozSiUC3v5OF3T-4GsuJgnqKio7xoRPuQUJ-Yy26SITFk_KyTpbGwyZ_hAqJm9nPI65LCdNquQg-wCVUvicqvNi1d_qqG6sPTvtaNLjUMexuCxnUchhylhyNRoyHQW4OUFsp9lYvR8mdTMO8X_iXbZokZemv4k274ZceufXgzazuDZlM6XrQWfhCHL5LhXpW7Y5yO6BdiUb-21zuzbvye_w0Upbb9daOrqTd2yMWwS99A8fgchG4nKNVWULhZGL7lZp2_JqVF_QEe4KbWe9pVnnEhypbBUEpDvp819LFOD1PhkCoNcTBYh_2tjjZqpN6SP-BvGkYw7eWLqxNJ6v5JXUpxfEGf44tFhdNGgnJt74eGzjWWr59kJiq1IJeke-Fx6zBG-ey5CUBtLvvzd9IbRxc0ni0BCGYAnfS3D_T_hKnO0_mg1PiqhYKpVTMiDrGKOu_V2qhju1UE-9n3CmkT6aHN14hsX-Uw-Qf-U27g.y_Mgn_4qI_VW_K7jZoK87w"

func GetToken() (token string) {
	response, _ := http.NewRequest(http.MethodPost, "https://hdbank-hackathon.auth.ap-southeast-1.amazoncognito.com/oauth2/token", bytes.NewBuffer([]byte(refeshString)))
	response.Header.Set("accept", "application/json")
	response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(response)
	dataRespone, _ := ioutil.ReadAll(res.Body)
	var a TokenType
	json.Unmarshal(dataRespone, &a)
	token = a.Id_token
	return
}
