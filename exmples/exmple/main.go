package main

import (
	"proxy-pool/check"
	"proxy-pool/fetch"
	"proxy-pool/http"
	"proxy-pool/stroage"
	"time"
)

func main() {
	mem := stroage.NewMemoryStroage()
	manager := fetch.New(mem, 60, 10*time.Second)
	manager.Register([]fetch.Fetcher{fetch.GetQuanWang})
	go manager.Run()
	defer manager.Stop()

	checker := check.New(mem, 10*time.Second, 1)
	go checker.Run()
	defer checker.Stop()
	srv := http.InitHttp(mem)
	srv.Run()
}
