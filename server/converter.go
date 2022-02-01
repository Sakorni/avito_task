package server

import (
	"avito_task/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (h *Handler) getConvertData(currency string) (map[string]float64, error) {
	path := fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest"+
		"?access_key=%s"+
		"&symbols=%s,%s",
		h.apiKey, service.DefaultCurrency, currency)
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("invalid response from convertation service, status:[%s], body:[%s]", resp.Status, string(body))
	}
	body := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&body)
	conv, ok := body["rates"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response body format, body:[%s]", body)
	}
	rates := ratesConvert(conv)
	return rates, nil
}

func ratesConvert(source map[string]interface{}) map[string]float64 {
	res := make(map[string]float64)
	for k, v := range source {
		res[k] = v.(float64)
	}
	return res
}
