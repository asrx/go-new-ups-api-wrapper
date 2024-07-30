package pojo

type RequestRating struct {
	*RateRequest `json:"RateRequest"`
}
type RateRequest struct {
	*Request      `json:"Request,omitempty"`
	*RateShipment `json:"Shipment" json:"shipment"`
}

type RateShipment struct {
	Shipper  Contact `json:"Shipper"`
	ShipTo   Contact `json:"ShipTo"`
	ShipFrom Contact `json:"ShipFrom"`
	// 付款账户信息
	*PaymentDetails        `json:"PaymentDetails"`
	*FRSPaymentInformation `json:"FRSPaymentInformation,omitempty"`
	// Shipment Rating Options container.
	// 装运评级选项集装箱
	*ShipmentRatingOptions  `json:"ShipmentRatingOptions"`
	*Service                `json:"Service"`
	*ShipmentServiceOptions `json:"ShipmentServiceOptions,omitempty"`
	// Array of objects (Rate_Shipment_Package) <= 200
	Package []*RatePackage `json:"Package"`
	// RatingMethodRequestedIndicator string         `json:"RatingMethodRequestedIndicator,omitempty"`
	// TaxInformationIndicator        string         `json:"TaxInformationIndicator,omitempty"`
	// WWEShipmentIndicator           string         `json:"WWEShipmentIndicator,omitempty"`
}

type FRSPaymentInformation struct {
	Type struct {
		Code string `json:"Code"`
	} `json:"Type"`
	AccountNumber string `json:"AccountNumber"`
}

func NewFRSPaymentInformationGFP(accountNumber string) *FRSPaymentInformation {
	return &FRSPaymentInformation{
		Type: struct {
			Code string `json:"Code"`
		}{"01"},
		AccountNumber: accountNumber,
	}
}

type RatePackage struct {
	*PackagingType         `json:"PackagingType"`
	*Dimensions            `json:"Dimensions"`
	*PackageWeight         `json:"PackageWeight"`
	*Commodity             `json:"Commodity,omitempty"`
	*PackageServiceOptions `json:"PackageServiceOptions,omitempty"`
}
