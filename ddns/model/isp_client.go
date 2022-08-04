package model

type IspClient struct {
	Client interface{}
}

var IspClientCache map[string]IspClient

func NewIspClient(client interface{}) IspClient {
	ispClient := new(IspClient)
	ispClient.Client = client
	return *ispClient
}

func GetIspClient(ispId string) IspClient {
	return IspClientCache[ispId]
}
