package go_new_ups_api_wrapper

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

const trackingUrl = "/api/track/v1/details/"

func Tracking(token string, TransId string, debug bool) (*pojo.RespTracking, error) {
	reqT := &pojo.RequestTracking{
		Locale:           "en_US",
		ReturnSignature:  true,
		ReturnMilestones: false,
		ReturnPOD:        true,
	}
	_src := "client"
	if debug {
		_src = "testing"
	}
	reqH := &pojo.RequestHeader{
		Authorization:  Bearer + token,
		TransId:        TransId,
		TransactionSrc: _src, // testing, client
	}
	jsonStr, err := trackingDo(reqT, reqH, debug)
	if err != nil {
		return nil, err
	}
	track := &pojo.RespTracking{}
	err = json.Unmarshal([]byte(jsonStr), track)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func trackingDo(reqTracking *pojo.RequestTracking, header *pojo.RequestHeader, debug bool) (string, error) {
	var trackMap map[string]interface{}
	var headerMap map[string]interface{}

	mapstructure.Decode(reqTracking, &trackMap)
	mapstructure.Decode(header, &headerMap)

	url := getRequestUrl(trackingUrl+header.TransId, debug)

	return HttpGet(url, headerMap, trackMap)
}
