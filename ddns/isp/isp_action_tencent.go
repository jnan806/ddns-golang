package isp

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/jnan806/ddns-golang/ddns/conf"
	"github.com/jnan806/ddns-golang/ddns/model"
	tencentCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencentDnspod20210323 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"reflect"
	"strconv"
)

type IspActionTencent struct {
}

const ISP_NAME_TENCENT = "tencent"

func InitTencentClients(IspClientCache *map[string]model.IspClient, ispConfs map[string]*conf.IspConf) {
	for ispId, ispConf := range ispConfs {
		if ISP_NAME_TENCENT == ispConf.IspType {
			ispClient := model.GetIspClient(ispId)
			if reflect.DeepEqual(ispClient, model.IspClient{}) {
				credential := tencentCommon.NewCredential(ispConf.AccessKeyId, ispConf.AccessKeySecret)
				newClient := tencentCommon.NewCommonClient(credential, ispConf.RegionId, profile.NewClientProfile())
				(*IspClientCache)[ispId] = model.NewIspClient(newClient)
				fmt.Printf("[运营商:%v][客户端Id:%v] 初始化[成功].\n", ISP_NAME_TENCENT, ispId)
			}
		}
	}
}

func (ispActionTencent IspActionTencent) PublishUnique(ispClient model.IspClient, domain string, recordMapping model.RecordMapping, ispDomainRecords model.IspDomainRecord) {
	recordListResponseParams, ok := ispDomainRecords.DomainRecords.(*tencentDnspod20210323.DescribeRecordListResponseParams)
	if !ok || len(recordListResponseParams.RecordList) == 0 {
		ispActionTencent.AddDomainRecord(ispClient, domain, recordMapping)
		return
	}

	for idx, record := range recordListResponseParams.RecordList {
		if idx == len(recordListResponseParams.RecordList)-1 {
			if recordMapping.Record != *record.Name || recordMapping.Ip != *record.Value {
				recordId := strconv.FormatUint(*record.RecordId, 10)
				ispActionTencent.UpdateDomainRecord(ispClient, domain, &recordId, recordMapping)
			}
		} else {
			recordId := strconv.FormatUint(*record.RecordId, 10)
			ispActionTencent.DeleteDomainRecord(ispClient, domain, &recordId)
		}
	}
}

func (ispActionTencent IspActionTencent) PublishNotUnique(ispClient model.IspClient, domain string, recordMapping model.RecordMapping, ispDomainRecords model.IspDomainRecord) {
	recordListResponseParams, ok := ispDomainRecords.DomainRecords.(*tencentDnspod20210323.DescribeRecordListResponseParams)
	if !ok || len(recordListResponseParams.RecordList) == 0 {
		ispActionTencent.AddDomainRecord(ispClient, domain, recordMapping)
		return
	}

	exist := false
	for _, record := range recordListResponseParams.RecordList {
		if recordMapping.Record == *record.Name && recordMapping.Ip == *record.Value {
			exist = true
		}
	}
	if !exist {
		ispActionTencent.AddDomainRecord(ispClient, domain, recordMapping)
	}
}

func (ispActionTencent IspActionTencent) DescribeDomainRecords(ispClient model.IspClient, domain string, recordMapping model.RecordMapping) model.IspDomainRecord {
	tencentClient := ispClient.Client.(*tencentCommon.Client)
	request := tencentDnspod20210323.NewDescribeRecordListRequest()
	request.RecordType = tea.String("A")
	request.Domain = tea.String(domain)
	request.Keyword = tea.String(recordMapping.Record)
	request.Limit = tea.Uint64(50)

	response := tencentDnspod20210323.NewDescribeRecordListResponse()
	err := tencentClient.Send(request, response)
	if err != nil {
		println(err.Error())
		return model.IspDomainRecord{}
	}

	return new(model.IspDomainRecord).IspDomainRecords(response.Response)
}

func (ispActionTencent IspActionTencent) AddDomainRecord(ispClient model.IspClient, domain string, recordMapping model.RecordMapping) {
	tencentClient := ispClient.Client.(*tencentCommon.Client)
	request := tencentDnspod20210323.NewCreateRecordRequest()
	request.Domain = &domain
	request.RecordType = tea.String("A")
	request.SubDomain = &recordMapping.Record
	request.Value = &recordMapping.Ip
	request.RecordLine = tea.String("默认")

	response := tencentDnspod20210323.NewDescribeRecordListResponse()
	err := tencentClient.Send(request, response)
	if err != nil {
		println(err.Error())
	}
}

func (ispActionTencent IspActionTencent) UpdateDomainRecord(ispClient model.IspClient, domain string, recordId *string, recordMapping model.RecordMapping) {
	tencentClient := ispClient.Client.(*tencentCommon.Client)
	request := tencentDnspod20210323.NewModifyRecordRequest()
	recordIdInt, _ := strconv.ParseUint(*recordId, 10, 64)
	request.Domain = &domain
	request.RecordId = &recordIdInt
	request.RecordType = tea.String("A")
	request.SubDomain = &recordMapping.Record
	request.Value = &recordMapping.Ip
	request.RecordLine = tea.String("默认")

	response := tencentDnspod20210323.NewDescribeRecordListResponse()
	err := tencentClient.Send(request, response)
	if err != nil {
		println(err.Error())
	}
}

func (ispActionTencent IspActionTencent) DeleteDomainRecord(ispClient model.IspClient, domain string, recordId *string) {
	tencentClient := ispClient.Client.(*tencentCommon.Client)
	request := tencentDnspod20210323.NewDeleteRecordRequest()
	recordIdInt, _ := strconv.ParseUint(*recordId, 10, 64)
	request.Domain = &domain
	request.RecordId = &recordIdInt

	response := tencentDnspod20210323.NewDeleteRecordResponse()
	err := tencentClient.Send(request, response)
	if err != nil {
		println(err.Error())
	}
}
