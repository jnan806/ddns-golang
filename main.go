package main

import (
	"fmt"
	"github.com/jnan806/ddns-golang/ddns/conf"
	"github.com/jnan806/ddns-golang/ddns/isp"
	"github.com/jnan806/ddns-golang/ddns/model"
	ddnsUtil "github.com/jnan806/ddns-golang/ddns/util"
	"reflect"
)

func main() {
	// 读取配置文件
	ddnsConf := conf.LoadConf()

	if reflect.DeepEqual(ddnsConf, conf.DdnsConf{}) {
		fmt.Println("配置文件中的配置内容为空...")
		return
	}

	// 初始化客户端缓存
	isp.InitIspClient(ddnsConf.Isps)

	run(&ddnsConf.DomainRecords)

}

func run(domainRecordConfItems *map[string]*conf.DomainRecordConfItem) {
	// 获取公网 ip
	publicIp := ddnsUtil.GetJsonIp()
	// 遍历 domainRecordConfItem
	for _, domainRecordConfItem := range *domainRecordConfItems {
		ispClient := model.GetIspClient(domainRecordConfItem.IspId)
		mappings := model.ParseRecordMapping(publicIp, *domainRecordConfItem)
		isp.Publish(ispClient, *domainRecordConfItem, mappings)
	}
}
