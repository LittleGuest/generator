package main

import (
	"context"
	_ "generator/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	addr         = "127.0.0.1:65535"
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 45
)

func main() {
	log.Println(banner)

	server := &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}

	go func() {
		log.Printf("listen at %v", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("服务启动失败：%v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("服务关闭失败：%v", err)
	}
	os.Exit(0)
}

var banner = `
   _____     _____      __      _    _____   ______       ____     ________     ____     ______    
  / ___ \   / ___/     /  \    / )  / ___/  (   __ \     (    )   (___  ___)   / __ \   (   __ \   
 / /   \_) ( (__      / /\ \  / /  ( (__     ) (__) )    / /\ \       ) )     / /  \ \   ) (__) )  
( (  ____   ) __)     ) ) ) ) ) )   ) __)   (    __/    ( (__) )     ( (     ( ()  () ) (    __/   
( ( (__  ) ( (       ( ( ( ( ( (   ( (       ) \ \  _    )    (       ) )    ( ()  () )  ) \ \  _  
 \ \__/ /   \ \___   / /  \ \/ /    \ \___  ( ( \ \_))  /  /\  \     ( (      \ \__/ /  ( ( \ \_)) 
  \____/     \____\ (_/    \__/      \____\  )_) \__/  /__(  )__\    /__\      \____/    )_) \__/  
`
