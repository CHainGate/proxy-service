package utils

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Opts *OptsType
)

type OptsType struct {
	ServerPort                  int
	SendGridApiKey              string
	EmailFrom                   string
	CoinMarketCapBaseUrl        string
	CoinMarketCapBaseUrlSandbox string
	CoinMarketCapApiKey         string
}

func NewOpts() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not find env file [%v], using defaults", err)
	}

	o := &OptsType{}
	flag.IntVar(&o.ServerPort, "SERVER_PORT", lookupEnvInt("SERVER_PORT", 8001), "Server PORT")
	flag.StringVar(&o.SendGridApiKey, "SENDGRID_API_KEY", lookupEnv("SENDGRID_API_KEY"), "SendGrid API Key")
	flag.StringVar(&o.EmailFrom, "EMAIL_FROM", lookupEnv("EMAIL_FROM"), "Email From")
	flag.StringVar(&o.CoinMarketCapBaseUrl, "COIN_MARKET_CAP_BASE_URL", lookupEnv("COIN_MARKET_CAP_BASE_URL"), "COIN MARKET CAP BASE URL")
	flag.StringVar(&o.CoinMarketCapBaseUrlSandbox, "COIN_MARKET_CAP_BASE_URL_SANDBOX", lookupEnv("COIN_MARKET_CAP_BASE_URL_SANDBOX"), "COIN MARKET CAP BASE URL SANDBOX")
	flag.StringVar(&o.CoinMarketCapApiKey, "COIN_MARKET_CAP_API_KEY", lookupEnv("COIN_MARKET_CAP_API_KEY"), "COIN MARKET CAP Api Key")

	Opts = o
}

func lookupEnv(key string, defaultValues ...string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	for _, v := range defaultValues {
		if v != "" {
			return v
		}
	}
	return ""
}

func lookupEnvInt(key string, defaultValues ...int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("LookupEnvInt[%s]: %v", key, err)
			return 0
		}
		return v
	}
	for _, v := range defaultValues {
		if v != 0 {
			return v
		}
	}
	return 0
}
