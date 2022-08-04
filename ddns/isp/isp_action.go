package isp

import (
	"github.com/jnan806/ddns-golang/ddns/conf"
	"github.com/jnan806/ddns-golang/ddns/model"
)

type IspActionInterface interface {
	PublishUnique(ispClient model.IspClient, domain string, recordMapping model.RecordMapping, ispDomainRecords model.IspDomainRecord)
	PublishNotUnique(ispClient model.IspClient, domain string, recordMapping model.RecordMapping, ispDomainRecords model.IspDomainRecord)
	AddDomainRecord(ispClient model.IspClient, domain string, recordMapping model.RecordMapping)
	UpdateDomainRecord(ispClient model.IspClient, domain string, recordId *string, recordMapping model.RecordMapping)
	DeleteDomainRecord(ispClient model.IspClient, domain string, recordId *string)
	DescribeDomainRecords(ispClient model.IspClient, domain string, recordMapping model.RecordMapping) model.IspDomainRecord
}

func Publish(ispClient model.IspClient, domainRecordConfItem conf.DomainRecordConfItem, recordMappings []model.RecordMapping) {
	ispActionInterface := DeterminedIspClientInterface(ispClient)
	if ispActionInterface != nil {
		doPublish(ispActionInterface, ispClient, domainRecordConfItem, recordMappings)
	}
}

func doPublish(ispActionInterface IspActionInterface, ispClient model.IspClient, domainRecordConfItem conf.DomainRecordConfItem, recordMappings []model.RecordMapping) {
	isUnique := domainRecordConfItem.IsUnique
	for _, recordMapping := range recordMappings {
		domainRecords := ispActionInterface.DescribeDomainRecords(ispClient, domainRecordConfItem.Domain, recordMapping)
		if isUnique {
			ispActionInterface.PublishUnique(ispClient, domainRecordConfItem.Domain, recordMapping, domainRecords)
		} else {
			ispActionInterface.PublishNotUnique(ispClient, domainRecordConfItem.Domain, recordMapping, domainRecords)
		}
	}
}
