/*
 * OpenAPI proxy service
 *
 * This is the OpenAPI definition of the proxy service.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package services

import (
	"chaingate/proxy-service/proxyApi"
	"chaingate/proxy-service/utils"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ConversionApiService is a service that implements the logic for the ConversionApiService
// This service should implement the business logic for every endpoint for the ConversionApiService API.
// Include any external packages or services that will be required by this service.
type ConversionApiService struct {
}

// NewConversionApiService creates a default api service
func NewConversionApiService() proxyApi.ConversionApiServicer {
	return &ConversionApiService{}
}

type CoinMarketCapConversion struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data []struct {
		Id          int                               `json:"id"`
		Symbol      string                            `json:"symbol"`
		Name        string                            `json:"name"`
		Amount      int                               `json:"amount"`
		LastUpdated time.Time                         `json:"last_updated"`
		Quote       map[string]map[string]interface{} `json:"quote"`
	} `json:"data"`
}

// GetPriceConversion - get price conversion
func (s *ConversionApiService) GetPriceConversion(_ context.Context, amount string, srcCurrency string, dstCurrency string) (proxyApi.ImplResponse, error) {
	conversion, err := getPriceConversion(amount, srcCurrency, dstCurrency)
	if err != nil {
		return proxyApi.Response(http.StatusInternalServerError, nil), err
	}

	price, err := getFloat(conversion.Data[0].Quote[strings.ToUpper(dstCurrency)]["price"])
	if err != nil {
		return proxyApi.Response(http.StatusInternalServerError, nil), err
	}

	priceConversionResponseDto := proxyApi.PriceConversionResponseDto{
		SrcCurrency: srcCurrency,
		DstCurrency: dstCurrency,
		Price:       price,
	}

	return proxyApi.Response(http.StatusOK, priceConversionResponseDto), nil
}

func getFloat(unknown interface{}) (float64, error) {
	switch t := unknown.(type) {
	case float64:
		return t, nil
	default:
		return 0, errors.New("Could not parse to float ")
	}
}

func getPriceConversion(amount string, srcCurrency string, dstCurrency string) (*CoinMarketCapConversion, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", utils.Opts.CoinMarketCapBaseUrl+"v2/tools/price-conversion", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("amount", amount)
	q.Add("symbol", srcCurrency)
	q.Add("convert", dstCurrency)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", utils.Opts.CoinMarketCapApiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var target CoinMarketCapConversion

	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
