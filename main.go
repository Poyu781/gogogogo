package main

import (
	"fmt"
	"net/http"

	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/pkg/setting"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
