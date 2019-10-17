package middleware

import (
	"log"
	"net/http"
	"testing"
)

func Handler1(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("this is Handler1")
		handler.ServeHTTP(w, req)
	})
}

func Handler2(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("this is Handler2")
		handler.ServeHTTP(w, req)
	})
}

func Handler3(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("this is Handler3")
		handler.ServeHTTP(w, req)
	})
}

func HandlerFunc1(w http.ResponseWriter, req *http.Request) {
	log.Println("this is HandlerFunc1")
}

func HandlerFunc2(w http.ResponseWriter, req *http.Request) {
	log.Println("this is HandlerFunc2")
}

func HandlerFunc3(w http.ResponseWriter, req *http.Request) {
	log.Println("this is HandlerFunc3")
}

func TestChainMiddleware(t *testing.T) {
	chain := NewChain(Handler1, Handler2, Handler3)
	http.Handle("/1", chain.ThenFunc(HandlerFunc1))
	http.Handle("/2", chain.ThenFunc(HandlerFunc2))
	http.Handle("/3", chain.ThenFunc(HandlerFunc3))
}
