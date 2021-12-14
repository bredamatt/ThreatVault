package main

import (
	"fmt"

	"github.com/AlienVault-OTX/OTX-Go-SDK/src/otxapi"
)

type OTXClient struct {
	client *otxapi.Client
}

func newClient() OTXClient {
	return OTXClient{client: otxapi.NewClient(nil)}
}

func main() {
	client := newClient()
	fmt.Println(client.client)
}
