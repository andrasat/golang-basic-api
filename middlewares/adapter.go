package middlewares

import (
	. "github.com/juliendschmidt/httprouter"
)

type Adapter func(Handler) Handler

func Adapt(h Handler, adapters ...Adapter) Handler {
	for _, adapter := range adapters {
		h := adapter(h)
	}
	return h
}
