package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

const shipmentUrl = "/api/shipments/v2403/ship?additionaladdressvalidation="

const labelRecoveryUrl = "/labels/{version}/recovery"

// 打单
func Shipment(token string, shipper, shipFrom, shipTo pojo.Contact,
	packages []*pojo.ShipPackage,
	serviceType string, debug bool) (*pojo.RespShipment, error) {

	reqShip := buildShipRequest(shipper, shipTo, shipFrom, packages, serviceType)

	reqH := &pojo.RequestHeader{
		Authorization: Bearer + token,
	}

	return shipmentDo(reqShip, reqH, debug)
}

// 退件单
func ShipmentReturnService(token string, shipper, shipFrom, shipTo pojo.Contact,
	packages []*pojo.ShipPackage,
	serviceType string, returnServiceEmail string, debug bool) (*pojo.RespShipment, error) {
	reqShip := buildShipRequest(shipper, shipTo, shipFrom, packages, serviceType)

	if len(returnServiceEmail) > 0 {
		reqShip.Shipment.ReturnService = pojo.NewReturnServiceERL()
		reqShip.ShipmentServiceOptions.LabelDelivery = pojo.NewLabelDelivery(returnServiceEmail)
	}

	reqH := &pojo.RequestHeader{
		Authorization: Bearer + token,
	}

	return shipmentDo(reqShip, reqH, debug)
}

func buildShipRequest(shipper pojo.Contact, shipTo pojo.Contact, shipFrom pojo.Contact, packages []*pojo.ShipPackage, serviceType string) *pojo.RequestShipping {
	reqShip := &pojo.RequestShipping{
		pojo.ShipmentRequest{
			Request: &pojo.Request{
				RequestOption: pojo.RequestOptionShoptimeintransit,
				SubVersion:    SubVersion,
			},
			Shipment: &pojo.Shipment{
				Shipper:                        shipper,
				ShipTo:                         shipTo,
				ShipFrom:                       shipFrom,
				PaymentInformation:             pojo.NewPaymentDetails(shipper.ShipperNumber),
				Service:                        pojo.Service03_Ground(),
				ShipmentRatingOptions:          pojo.NewShipmentRatingOptions(),
				RatingMethodRequestedIndicator: "Y",
				Package:                        packages,
				ShipmentServiceOptions:         &pojo.ShipmentServiceOptions{},
			},
			LabelSpecification: pojo.NewLabelSpecificationGif(),
			// LabelSpecification: pojo.NewLabelSpecificationZPL(),
		},
	}

	// GFP
	if strings.ToUpper(serviceType) == pojo.UPS_GFP {
		reqShip.Shipment.ShipmentRatingOptions.FRSShipmentIndicator = "Y"
		reqShip.Shipment.FRSPaymentInformation = pojo.NewFRSPaymentInformationGFP(shipper.ShipperNumber)
		reqShip.Shipment.PaymentInformation = nil
	}
	return reqShip
}

func shipmentDo(reqShip *pojo.RequestShipping, header *pojo.RequestHeader, debug bool) (*pojo.RespShipment, error) {
	var headerMap map[string]interface{}
	mapstructure.Decode(header, &headerMap)

	url := getRequestUrl(shipmentUrl+reqShip.ShipTo.City, debug)

	marshal, _ := json.Marshal(reqShip)
	fmt.Println(string(marshal))

	jsonStr, err := HttpPost(url, headerMap, reqShip)
	if err != nil {
		return nil, err
	}
	ship := &pojo.RespShipment{}
	err = json.Unmarshal([]byte(jsonStr), ship)
	if err != nil {
		return nil, err
	}
	return ship, nil
}
