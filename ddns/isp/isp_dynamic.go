package isp

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/jnan806/ddns-golang/ddns/conf"
	"github.com/jnan806/ddns-golang/ddns/model"
	tencentCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func InitIspClient(ispConf map[string]*conf.IspConf) {
	model.IspClientCache = make(map[string]model.IspClient, len(ispConf))
	InitAliyunClients(&model.IspClientCache, ispConf)
	InitTencentClients(&model.IspClientCache, ispConf)
}

func DeterminedIspClientInterface(ispClient model.IspClient) IspActionInterface {
	var ispActionInterface IspActionInterface
	if _, ok := ispClient.Client.(*alidns20150109.Client); ok {
		ispActionInterface = new(IspActionAliyun)
	}

	if _, ok := ispClient.Client.(*tencentCommon.Client); ok {
		ispActionInterface = new(IspActionTencent)
	}

	return ispActionInterface
}
