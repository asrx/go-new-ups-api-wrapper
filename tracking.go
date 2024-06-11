package go_new_ups_api_wrapper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/idoubi/goz"
	"github.com/mitchellh/mapstructure"
)

const trackingUrl = "/api/track/v1/details/"

// 请求表单
type RequestTracking struct {
	Locale           string `json:"Locale,omitempty"`
	ReturnSignature  bool   `json:"ReturnSignature,omitempty"`
	ReturnMilestones bool   `json:"ReturnMilestones,omitempty"`
	ReturnPOD        bool   `json:"ReturnPOD,omitempty"`
}

// 请求 header
type RequestHeader struct {
	Authorization  string `json:"Authorization,require" :"Authorization" :"Authorization"`
	TransId        string `json:"TransId,require" :"TransId"`
	TransactionSrc string `json:"TransactionSrc" require`
}

func Tracking(token string, TransId string, debug bool) (map[string]interface{}, error) {
	reqT := &RequestTracking{
		Locale:           "en_US",
		ReturnSignature:  true,
		ReturnMilestones: false,
		ReturnPOD:        true,
	}
	reqH := &RequestHeader{
		Authorization:  Bearer + token,
		TransId:        TransId,
		TransactionSrc: "testing", // client
	}
	return TrackingDo(reqT, reqH, debug)
}

func TrackingDo(reqTracking *RequestTracking, header *RequestHeader, debug bool) (map[string]interface{}, error) {
	var trackMap map[string]interface{}
	var headersMap map[string]interface{}

	mapstructure.Decode(reqTracking, trackMap)
	mapstructure.Decode(header, headersMap)
	client := goz.NewClient()
	url := getRequestUrl(trackingUrl+header.TransId, debug)
	response, err := client.Get(url, goz.Options{
		Timeout: 3000,
		// Query:        nil,
		Headers: headersMap,
		// Cookies:      nil,
		FormParams: trackMap,
	})
	// 请求验证
	if err != nil {
		return nil, err
	}
	// 获取响应体
	body, err := response.GetBody()
	if err != nil {
		return nil, err
	}

	// 判断结果
	if response.GetStatusCode() != http.StatusOK {
		respErr := &ResponseErr{}
		err = json.Unmarshal([]byte(body.String()), respErr)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(respErr.Response.Errors[0].Message)
	}
	// 返回响应结构体
	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(body.String()), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 跟踪响应体
type ResponseTracking struct {
	TrackResponse struct {
		Shipment []struct {
			InquiryNumber string `json:"inquiryNumber"`
			Package       []struct {
				TrackingNumber string        `json:"trackingNumber"`
				DeliveryDate   []interface{} `json:"deliveryDate"`
				Activity       []struct {
					Location struct {
						Address struct {
							City          string `json:"city,omitempty"`
							StateProvince string `json:"stateProvince,omitempty"`
							CountryCode   string `json:"countryCode,omitempty"`
							Country       string `json:"country,omitempty"`
						} `json:"address"`
						Slic string `json:"slic,omitempty"`
					} `json:"location"`
					Status struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Code        string `json:"code"`
						StatusCode  string `json:"statusCode"`
					} `json:"status"`
					Date      string `json:"date"`
					Time      string `json:"time"`
					GmtDate   string `json:"gmtDate"`
					GmtOffset string `json:"gmtOffset"`
					GmtTime   string `json:"gmtTime"`
				} `json:"activity"`
				CurrentStatus struct {
					Description               string `json:"description"`
					Code                      string `json:"code"`
					SimplifiedTextDescription string `json:"simplifiedTextDescription"`
				} `json:"currentStatus"`
				PackageAddress []struct {
					Type    string `json:"type"`
					Address struct {
						City          string `json:"city"`
						StateProvince string `json:"stateProvince"`
						CountryCode   string `json:"countryCode"`
						Country       string `json:"country"`
					} `json:"address"`
				} `json:"packageAddress"`
				Weight struct {
					UnitOfMeasurement string `json:"unitOfMeasurement"`
					Weight            string `json:"weight"`
				} `json:"weight"`
				Service struct {
					Code        string `json:"code"`
					LevelCode   string `json:"levelCode"`
					Description string `json:"description"`
				} `json:"service"`
				ReferenceNumber []struct {
					Type   string `json:"type"`
					Number string `json:"number"`
				} `json:"referenceNumber"`
				DeliveryInformation struct {
					ReceivedBy    string `json:"receivedBy"`
					Location      string `json:"location"`
					DeliveryPhoto struct {
						PhotoCaptureInd        string `json:"photoCaptureInd"`
						IsNonPostalCodeCountry bool   `json:"isNonPostalCodeCountry"`
					} `json:"deliveryPhoto"`
				} `json:"deliveryInformation"`
				Dimension struct {
					Height          string `json:"height"`
					Length          string `json:"length"`
					Width           string `json:"width"`
					UnitOfDimension string `json:"unitOfDimension"`
				} `json:"dimension"`
				PackageCount int `json:"packageCount"`
			} `json:"package"`
		} `json:"shipment"`
	} `json:"trackResponse"`
}
