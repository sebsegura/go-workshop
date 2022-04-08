package main

import "fmt"

type APIClient interface {
	SayHello() string
}

type apiClient struct {}

func (a *apiClient) SayHello() string {
	return "Hello"
}

type BridgeService interface {
	Hello()
}

type bridgeService struct {
	Client APIClient
}

func NewBridgeService() BridgeService {
	a := apiClient{}
	return &bridgeService{
		Client: &a,
	}
}

func (b *bridgeService) Hello() {
	fmt.Println(b.Client.SayHello())
}

func main() {
	bridge := NewBridgeService()
	bridge.Hello()
}


