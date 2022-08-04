package model

import (
	"github.com/jnan806/ddns-golang/ddns/conf"
	"strings"
)

type RecordMapping struct {
	Record string
	Ip     string
}

func ParseRecordMapping(publicIp string, domainRecordConfItem conf.DomainRecordConfItem) []RecordMapping {
	recordmappings := make([]RecordMapping, len(domainRecordConfItem.RecordMapping))
	for idx, mappingItem := range domainRecordConfItem.RecordMapping {
		if strings.Index(mappingItem, ":") < 0 {
			mappingItem += ":"
		}
		split := strings.Split(mappingItem, ":")
		recordmapping := RecordMapping{split[0], split[1]}
		if recordmapping.Record == "" {
			continue
		}
		if recordmapping.Ip == "" {
			recordmapping.Ip = publicIp
		}
		recordmappings[idx] = recordmapping
	}
	return recordmappings
}
