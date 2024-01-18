// Code generated by "requestgen -method GET -url /fapi/v2/balance -type FuturesGetAccountBalanceRequest -responseType []FuturesBalance"; DO NOT EDIT.

package binanceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (f *FuturesGetAccountBalanceRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (f *FuturesGetAccountBalanceRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (f *FuturesGetAccountBalanceRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := f.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if f.isVarSlice(_v) {
			f.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (f *FuturesGetAccountBalanceRequest) GetParametersJSON() ([]byte, error) {
	params, err := f.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (f *FuturesGetAccountBalanceRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (f *FuturesGetAccountBalanceRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (f *FuturesGetAccountBalanceRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (f *FuturesGetAccountBalanceRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (f *FuturesGetAccountBalanceRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := f.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (f *FuturesGetAccountBalanceRequest) Do(ctx context.Context) ([]FuturesBalance, error) {

	// no body params
	var params interface{}
	query := url.Values{}

	apiURL := "/fapi/v2/balance"

	req, err := f.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := f.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []FuturesBalance
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
