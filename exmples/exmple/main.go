package main

import (
	"proxy-pool/check"
	"proxy-pool/fetch"
	"proxy-pool/fetch/fetcher"
	"proxy-pool/http"
	"proxy-pool/stroage"
	"time"
)

func main() {
	mem := stroage.NewMemoryStroage()
	manager := fetch.New(mem, fetch.ConcurrencyOption(10), fetch.IntervalOption(10*time.Minute))
	manager.Register([]fetch.Fetcher{fetcher.GetXiChi, fetcher.GetIPKuByAPI, fetcher.GetIPYunDaiLi})
	defer manager.Stop()
	go manager.Run()

	checker := check.New(mem, check.ConcurrencyOption(10), check.IntervalOption(10*time.Minute))
	defer checker.Stop()
	go checker.Run()

	srv := http.InitHttp(mem)
	srv.Run()
}
