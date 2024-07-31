package go_new_ups_api_wrapper

import (
	"encoding/json"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

// const ratingUrl = "/api/rating/{version}/{requestoption}"
const ratingUrl = "/api/rating/v2403/"

func Rating(token string, shipper, shipFrom, shipTo pojo.Contact, packages []*pojo.RatePackage, serviceType string, debug bool) (*pojo.RespRate, error) {
	reqRate := &pojo.RequestRating{
		&pojo.RateRequest{
			Request: &pojo.Request{
				// RequestOption: pojo.RequestOptionRate,
				// RequestOption: pojo.RequestOptionShop,
				// RequestOption: pojo.RequestOptionRatetimeintransit,
				RequestOption: pojo.RequestOptionShoptimeintransit,
				SubVersion:    SubVersion,
			},
			RateShipment: &pojo.RateShipment{
				Shipper:                        shipper,
				ShipTo:                         shipTo,
				ShipFrom:                       shipFrom,
				PaymentDetails:                 pojo.NewPaymentDetails(shipper.ShipperNumber),
				ShipmentRatingOptions:          pojo.NewShipmentRatingOptions(),
				Service:                        pojo.Service03_Ground(),
				Package:                        packages,
				RatingMethodRequestedIndicator: "Y",
				// TaxInformationIndicator:        "Y",
				// PromotionalDiscountInformation: "Y",
				// MasterCartonIndicator: "Y",
				// WWEShipmentIndicator: "Y",
			},
		},
	}
	if strings.ToUpper(serviceType) == pojo.UPS_GFP {
		reqRate.RateRequest.ShipmentRatingOptions.FRSShipmentIndicator = "Y"
		reqRate.RateRequest.FRSPaymentInformation = pojo.NewFRSPaymentInformationGFP(shipper.ShipperNumber)
	}

	reqH := &pojo.RequestHeader{
		Authorization: Bearer + token,
	}
	jsonStr, err := ratingDo(reqRate, reqH, debug)
	if err != nil {
		return nil, err
	}

	rate := &pojo.RespRate{}
	err = json.Unmarshal([]byte(jsonStr), rate)
	if err != nil {
		return nil, err
	}
	return rate, nil
}

func ratingDo(reqRating *pojo.RequestRating, header *pojo.RequestHeader, debug bool) (string, error) {
	var headerMap map[string]interface{}
	mapstructure.Decode(header, &headerMap)

	url := getRequestUrl(ratingUrl+string(pojo.RequestOptionRate), debug)

	return HttpPost(url, headerMap, reqRating)
}
