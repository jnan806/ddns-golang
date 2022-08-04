package model

type IspDomainRecord struct {
	DomainRecords interface{}
}

func (ispDomainRecord IspDomainRecord) IspDomainRecords(records interface{}) IspDomainRecord {
	ispDomainRecord.DomainRecords = records
	return ispDomainRecord
}
