package main

import (
	"blogger/pkg/setting"
	"blogger/routers"
	"fmt"
	"net/http"
)

func main() {
	r :=routers.InitRouter()

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HttpPort),
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
		MaxHeaderBytes:    1<<20,
	}
	s.ListenAndServe()
}
