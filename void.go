package go_new_ups_api_wrapper

import (
	"encoding/json"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

const voidUrl = "/api/shipments/v2403/void/cancel/"

func Void(token, shipmentidentificationnumber string, trackingNumber []string, debug bool) (*pojo.RespVoid, error) {
	url := getRequestUrl(voidUrl+shipmentidentificationnumber, debug)

	if trackingNumber != nil {
		numbers := strings.Join(trackingNumber, ",")
		url += "?trackingnumber=" + numbers
	}

	reqH := &pojo.RequestHeader{
		Authorization: Bearer + token,
	}
	var headerMap map[string]interface{}
	mapstructure.Decode(reqH, &headerMap)

	jsonStr, err := HttpDelete(url, headerMap, nil)
	if err != nil {
		return nil, err
	}
	result := &pojo.RespVoid{}
	err = json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
