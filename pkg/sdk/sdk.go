package sdk

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type Context struct {
	*resty.Request
	*resty.Response
}

// Handler 处理
type Handler interface {
	// 自身的业务
	Do(c *Context) error
	// 设置下一个对象
	SetNext(h Handler) Handler
	// 执行
	Run(c *Context) error
}

// Next 抽象出来的 可被合成复用的结构体
type Next struct {
	// 下一个对象
	nextHandler Handler
}

// SetNext 实现好的 可被复用的SetNext方法
// 返回值是下一个对象 方便写成链式代码优雅
// 例如 Client.SetNext(argumentsHandler).SetNext(signHandler).SetNext(frequentHandler)
func (n *Next) SetNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

// Run 执行
func (n *Next) Run(c *Context) (err error) {
	// 由于go无继承的概念 这里无法执行当前handler的Do
	// n.Do(c)
	if n.nextHandler != nil {
		// 合成复用下的变种
		// 执行下一个handler的Do
		if err = (n.nextHandler).Do(c); err != nil {
			return
		}
		// 执行下一个handler的Run
		return (n.nextHandler).Run(c)
	}
	return
}

// Client 空Handler
// 由于go无继承的概念 作为链式调用的第一个载体 设置实际的下一个对象
type Client struct {
	// 合成复用Next的`nextHandler`成员属性、`SetNext`成员方法、`Run`成员方法
	Next
	*Context
}

// Do 空Handler的Do
func (h *Client) Do(c *Context) (err error) {
	return
}

func NewSdk() *Client {
	return &Client{
		Context: &Context{
			Request: resty.New().OnError(func(r *resty.Request, err error) {
				if v, ok := err.(*resty.ResponseError); ok {
					v.Err = errors.New("网络连接错误")
				}
			}).R(),
		},
	}
}
