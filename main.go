package main

import (
	"context"
	"cts/server"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//context.WithCancel()
	ginRouter := server.Router()
	address := ":9999"
	srv := &http.Server{
		Addr:    address,
		Handler: ginRouter,
	}

	srv.RegisterOnShutdown(func() {
		fmt.Println("clean resource on shutdown...")
		time.Sleep(time.Second * 2)
		fmt.Println("clean resource success.")
		wg.Done()
	})
	wg.Add(2)
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit)
		<-quit // block

		timeCtx, cf := context.WithTimeout(context.Background(), time.Second*5)
		defer cf()
		var done = make(chan struct{}, 1)

		go func() {
			if err := srv.Shutdown(timeCtx); err != nil {
				fmt.Println("web server shutdown error: ", err)
			} else {
				fmt.Println("web server shutdown ok.")
			}
			done <- struct{}{}
			wg.Done()
		}()

		select {
		case <-timeCtx.Done():
			fmt.Println("web server shutdown timeout.")
		case <-done:
		}
	}()
	err := srv.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("web server start failed: %v\n", err)
			return
		}
	}
	wg.Wait() // 等待wg计数减少为0
	fmt.Println("program exit ok.")
}
