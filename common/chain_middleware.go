// 借鉴
// 将给个handler链接起来
package common

import "net/http"

type Constructor func(handler http.Handler) http.Handler

type Chain struct {
	constructor [] Constructor
}

func NewChain(constructor ...Constructor) Chain {
	return Chain{append([]Constructor(nil), constructor...)}
}

func (c Chain) Then(handler http.Handler) http.Handler {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	for k := range c.constructor {
		handler = c.constructor[len(c.constructor)-1-k](handler)
	}
	return handler
}

func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	return c.Then(fn)
}

func (c Chain) Append(constructor ...Constructor) Chain {
	newCons := make([]Constructor, 0, len(c.constructor)+len(constructor))
	newCons = append(newCons, c.constructor...)
	newCons = append(newCons, constructor...)
	return Chain{newCons}
}

func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.constructor...)
}
