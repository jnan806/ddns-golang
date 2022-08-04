package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PublicIpResp struct {
	Ip string `json:"ip"`
}

const jsonIpUrl = "https://jsonip.com/"

func GetJsonIp() string {
	var ip string
	resp, err := http.Get(jsonIpUrl)
	if err != nil {
		println(err)
	} else {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println(err)
		} else {
			var publicIpResp PublicIpResp
			json.Unmarshal(bytes, &publicIpResp)
			ip = publicIpResp.Ip
		}
	}
	return ip
}
