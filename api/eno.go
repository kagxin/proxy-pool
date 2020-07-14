package api

import (
	"github.com/kagxin/golib/web/common/eno"
)

var (
	// OK 成功
	OK = eno.New(0, "成功")
	// RequestError RequestError
	RequestError = eno.New(4000, "Request Param Error")

	// NoFound no found
	NoFound = eno.New(40004, "no found")

	// ServerError servier err
	ServerError = eno.New(50000, "server error")
)
