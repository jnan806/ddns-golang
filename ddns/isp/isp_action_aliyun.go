package isp

import (
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/jnan806/ddns-golang/ddns/conf"
	"github.com/jnan806/ddns-golang/ddns/model"
	"reflect"
)

type IspActionAliyun struct {
}

const ISP_NAME_ALIYUN = "aliyun"

func InitAliyunClients(IspClientCache *map[string]model.IspClient, ispConfs map[string]*conf.IspConf) {
	for ispId, ispConf := range ispConfs {
		if ISP_NAME_ALIYUN == ispConf.IspType {
			ispClient := model.GetIspClient(ispId)
			if reflect.DeepEqual(ispClient, model.IspClient{}) {
				config := &openapi.Config{RegionId: &ispConf.RegionId, AccessKeyId: &ispConf.AccessKeyId, AccessKeySecret: &ispConf.AccessKeySecret}
				newClient, err := alidns20150109.NewClient(config)
				if err == nil {
					(*IspClientCache)[ispId] = model.NewIspClient(newClient)
					fmt.Printf("[运营商:%v][客户端Id:%v] 初始化[成功].\n", ISP_NAME_ALIYUN, ispId)
				} else {
					fmt.Printf("[运营商:%v][客户端Id:%v] 初始化[失败]...[原因：%v]\n", ISP_NAME_ALIYUN, ispId, err)
				}
			}
		}
	}
}

func (ispActionAliyun IspActionAliyun) PublishUnique(ispClient model.IspClient, domain string, recordMapping model.RecordMapping, ispDomainRecords model.IspDomainRecord) {
	domainRecords, ok := ispDomainRecords.DomainRecords.(*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecords)
	if !ok || len(domainRecords.Record) == 0 {
		ispActionAliyun.AddDomainRecord(ispClient, domain, recordMapping)
		return
	}

	for idx, record := range domainRecords.Record {
		if idx == len(domainRecords.Record)-1 {
			if recordMapping.Record != *record.RR || recordMapping.Ip != *record.Value {
				ispActionAliyun.UpdateDomainRecord(ispClient, domain, record.RecordId, recordMapping)
			}
		} else {
			ispActionAliyun.DeleteDomainRecord(ispClient, domain, record.RecordId)
		}
	}
}

func (ispActionAliyun IspActionAliyun) PublishNotUnique(ispClient model.IspClient, domain string, recordMapping model.RecordMapping, ispDomainRecords model.IspDomainRecord) {
	domainRecords, ok := ispDomainRecords.DomainRecords.(*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecords)
	if !ok || len(domainRecords.Record) == 0 {
		ispActionAliyun.AddDomainRecord(ispClient, domain, recordMapping)
		return
	}

	exist := false
	for _, record := range domainRecords.Record {
		if recordMapping.Record == *record.RR && recordMapping.Ip == *record.Value {
			exist = true
		}
	}
	if !exist {
		ispActionAliyun.AddDomainRecord(ispClient, domain, recordMapping)
	}
}

func (ispActionAliyun IspActionAliyun) DescribeDomainRecords(ispClient model.IspClient, domain string, recordMapping model.RecordMapping) model.IspDomainRecord {
	aliyunClient := ispClient.Client.(*alidns20150109.Client)
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{}
	describeDomainRecordsRequest.Type = tea.String("A")
	describeDomainRecordsRequest.DomainName = tea.String(domain)
	describeDomainRecordsRequest.RRKeyWord = tea.String(recordMapping.Record)
	describeDomainRecordsRequest.PageSize = tea.Int64(50)

	body, err := aliyunClient.DescribeDomainRecords(describeDomainRecordsRequest)
	if err != nil {
		println(err.Error())
		return model.IspDomainRecord{}
	}

	return new(model.IspDomainRecord).IspDomainRecords(body.Body.DomainRecords)
}

func (ispActionAliyun IspActionAliyun) AddDomainRecord(ispClient model.IspClient, domain string, recordMapping model.RecordMapping) {
	aliyunClient := ispClient.Client.(*alidns20150109.Client)
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{}
	addDomainRecordRequest.Type = tea.String("A")
	addDomainRecordRequest.DomainName = &domain
	addDomainRecordRequest.RR = &recordMapping.Record
	addDomainRecordRequest.Value = &recordMapping.Ip
	_, err := aliyunClient.AddDomainRecord(addDomainRecordRequest)
	if err != nil {
		println(err.Error())
	}
}

func (ispActionAliyun IspActionAliyun) UpdateDomainRecord(ispClient model.IspClient, domain string, recordId *string, recordMapping model.RecordMapping) {
	aliyunClient := ispClient.Client.(*alidns20150109.Client)
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{}
	updateDomainRecordRequest.RecordId = recordId
	updateDomainRecordRequest.Type = tea.String("A")
	updateDomainRecordRequest.RR = &recordMapping.Record
	updateDomainRecordRequest.Value = &recordMapping.Ip
	_, err := aliyunClient.UpdateDomainRecord(updateDomainRecordRequest)
	if err != nil {
		println(err.Error())
	}
}

func (ispActionAliyun IspActionAliyun) DeleteDomainRecord(ispClient model.IspClient, domain string, recordId *string) {
	aliyunClient := ispClient.Client.(*alidns20150109.Client)
	deleteDomainRecordRequest := &alidns20150109.DeleteDomainRecordRequest{}
	deleteDomainRecordRequest.RecordId = recordId
	_, err := aliyunClient.DeleteDomainRecord(deleteDomainRecordRequest)
	if err != nil {
		println(err.Error())
	}
}
