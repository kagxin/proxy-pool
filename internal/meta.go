package internal

import "time"

var (
	UA             = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
	HttpsBin       = "https://httpbin.org/get"
	HttpBin        = "http://httpbin.org/get"
	HttpBinTimeOut = 10 * time.Second
)

var (
	Running int32 = 1
	Stop    int32 = 0
)

var (
	Concurrency = 10
	Interval    = 10 * time.Minute
)
