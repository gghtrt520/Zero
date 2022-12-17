package test

import (
	"fmt"
	"testing"
	"zlei/pkg/sdk"
	"zlei/pkg/sdk/api"
)

func TestPing(t *testing.T) {
	client := sdk.NewSdk()
	client.SetNext(&api.BindRequest{}).SetNext(&api.Send{}).SetNext(&api.Print{})
	if err := client.Run(client.Context); err != nil {
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
}
