package api

import (
	"fmt"
	"net/http"
	"zlei/pkg/sdk"
)

type BindRequest struct {
	sdk.Next
}

type Send struct {
	sdk.Next
}

type Print struct {
	sdk.Next
}

// Do 校验参数的逻辑
func (h *BindRequest) Do(c *sdk.Context) (err error) {
	c.Request.URL = "https://httpbin.org/get"
	c.Request.Method = http.MethodGet
	return
}

// Do 校验参数的逻辑
func (h *Send) Do(c *sdk.Context) (err error) {
	resp, err := c.Send()
	if err != nil {
		return err
	}
	c.Response = resp
	return
}

func (h *Print) Do(c *sdk.Context) (err error) {
	fmt.Println(c.Response.StatusCode())
	fmt.Println(c.Response.IsSuccess())
	fmt.Println(c.Response.IsError())
	fmt.Println(string(c.Response.Body()))
	return
}
