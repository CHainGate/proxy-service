package main

import (
	"chaingate/proxy-service/proxyApi"
	"chaingate/proxy-service/services"
	"chaingate/proxy-service/utils"
	"log"
	"net/http"
	"strconv"
)

func main() {
	utils.NewOpts() // create utils.Opts (env variables)

	EmailApiService := services.NewEmailApiService()
	EmailApiController := proxyApi.NewEmailApiController(EmailApiService)

	ExchangeRateApiService := services.NewConversionApiService()
	ExchangeRateApiController := proxyApi.NewConversionApiController(ExchangeRateApiService)

	WebhookApiService := services.NewWebhookApiService()
	WebhookApiController := proxyApi.NewWebhookApiController(WebhookApiService)

	router := proxyApi.NewRouter(EmailApiController, ExchangeRateApiController, WebhookApiController)

	// https://ribice.medium.com/serve-swaggerui-within-your-golang-application-5486748a5ed4
	sh := http.StripPrefix("/api/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/api/swaggerui/").Handler(sh)

	log.Println("Starting proxy-service on port " + strconv.Itoa(utils.Opts.ServerPort))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(utils.Opts.ServerPort), router))
}

func firstTestFunc() string {
	return "firstTestFunc"
}
