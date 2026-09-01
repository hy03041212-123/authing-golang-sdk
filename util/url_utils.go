package util

import (
	"github.com/hy03041212-123/authing-golang-sdk/v3/constant"
	"net/url"
)

func GetQueryString(queryMap map[string]string) string {
	if queryMap == nil || len(queryMap) == 0 {
		return constant.StringEmpty
	}
	queryValue := url.Values{}
	for key, value := range queryMap {
		if value == "" {
			continue
		}
		queryValue.Add(key, value)
	}
	return queryValue.Encode()

}
