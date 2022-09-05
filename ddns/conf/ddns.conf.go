package conf

import (
	"gopkg.in/ini.v1"
	"reflect"
)

const confPath string = "conf/ddns.conf"

type IspConf struct {
	IspType         string `ini:"ispType"`
	RegionId        string `ini:"regionId"`
	AccessKeyId     string `ini:"accessKeyId"`
	AccessKeySecret string `ini:"accessKeySecret"`
}

type DomainRecordConfItem struct {
	IspId         string   `ini:"ispId"`
	Domain        string   `ini:"domain"`
	RecordMapping []string `ini:"recordMapping"`
	IsUnique      bool     `ini:"isUnique"`
}

type DdnsConf struct {
	Isps          map[string]*IspConf
	DomainRecords map[string]*DomainRecordConfItem
}

var ddnsConf = DdnsConf{}

func LoadConf() DdnsConf {
	if !reflect.DeepEqual(ddnsConf, DdnsConf{}) {
		return ddnsConf
	}

	confFile, err := ini.Load(confPath)
	if err != nil {
		println(err)
		return ddnsConf
	}

	confFile.MapTo(ddnsConf)

	sections := confFile.Sections()
	for _, section := range sections {
		ispType := section.Key("ispType").Value()
		ispId := section.Key("ispId").Value()

		if ispType != "" {
			loadIsp(section)
		}

		if ispId != "" {
			loadDomainRecord(confFile, section)
		}
	}

	return ddnsConf
}

func loadIsp(section *ini.Section) {
	tempIsp := new(IspConf)
	section.MapTo(tempIsp)

	if ddnsConf.Isps == nil {
		ddnsConf.Isps = make(map[string]*IspConf)
	}
	ddnsConf.Isps[section.Name()] = tempIsp
}

func loadDomainRecord(confFile *ini.File, section *ini.Section) {
	tempDomainRecordItem := new(DomainRecordConfItem)
	section.MapTo(tempDomainRecordItem)
	ispId := tempDomainRecordItem.IspId
	isp := ddnsConf.Isps[ispId]
	if reflect.DeepEqual(isp, IspConf{}) {
		loadIsp(section)
		isp = ddnsConf.Isps[ispId]
		if !reflect.DeepEqual(isp, IspConf{}) {
			loadDomainRecord(confFile, section)
		}
	}

	if ddnsConf.DomainRecords == nil {
		ddnsConf.DomainRecords = make(map[string]*DomainRecordConfItem)
	}
	ddnsConf.DomainRecords[section.Name()] = tempDomainRecordItem

}
