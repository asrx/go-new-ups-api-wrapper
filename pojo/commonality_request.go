package pojo

const (
	UPS_GROUND = "GROUND"
	UPS_GFP    = "GFP"
)

type Request struct {
	RequestOption        RequestOption `json:"requestOption,omitempty"`
	SubVersion           string        `json:"SubVersion,omitempty"`
	TransactionReference `json:"TransactionReference,omitempty" json:"transactionReference"`
}
type TransactionReference struct {
	CustomerContext string `json:"CustomerContext"`
}

type Contact struct {
	// 托运人or公司名
	Name string `json:"Name"`
	// 托运人
	AttentionName string `json:"AttentionName,omitempty"`
	// 托运人的UPS账号。需要有效的帐号才能收到协商的费率。否则可选。请求UserLevelDiscount时不能出现。
	ShipperNumber string `json:"ShipperNumber,omitempty"` // 仅Shipper含此字段
	Address       `json:"Address"`
	Phone         `json:"Phone,omitempty"`
	EMailAddress  string `json:"EMailAddress,omitempty"`
}
type Phone struct {
	Number    string `json:"Number"`
	Extension string `json:"Extension,omitempty"`
}
type Address struct {
	AddressLine       []string `json:"AddressLine"`
	City              string   `json:"City"`
	StateProvinceCode string   `json:"StateProvinceCode"`
	PostalCode        string   `json:"PostalCode"`
	CountryCode       string   `json:"CountryCode"`
	// 此字段是一个标志，用于指示接收器是否是住宅位置。如果ResidentialAddressIndicator标签存在，则为True。这是一个空标签，其中的任何值都将被忽略。
	ResidentialAddressIndicator string `json:"ResidentialAddressIndicator,omitempty"`
}

type Service struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type Dimensions struct {
	*UnitOfMeasurement `json:"UnitOfMeasurement"`
	// 包装长度。长度必须是容器的最长尺寸。有效值为0到108 IN和0到270 CM。
	Length string `json:"Length"`
	Width  string `json:"Width"`
	Height string `json:"Height"`
}

type UnitOfMeasurement struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}
type PackageWeight struct {
	*UnitOfMeasurement `json:"UnitOfMeasurement"`
	Weight             string `json:"Weight"`
}

/*
Rating => PackagingType
Shipment => Packaging
*/
type PackagingType struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type ShipmentCharge struct {
	Type        string `json:"Type"`
	BillShipper `json:"BillShipper"`
}
type BillShipper struct {
	AccountNumber string `json:"AccountNumber"`
}
type ShipmentRatingOptions struct {
	// NegotiatedRatesIndicator-需要显示两种类型的折扣：1）出价或基于帐户的费率2）网络/促销折扣出价基于帐户的价格：如果该指示器存在，则发货人已获得授权，并且评级API XML请求配置为返回协商价格，则应在响应中返回协商价格。网络/促销折扣：如果该指标存在，则托运人有权获得网络/促销优惠，则应在回复中返回协商价格。
	NegotiatedRatesIndicator string `json:"NegotiatedRatesIndicator"`
	// FRS指标。需要该指标来获取UPS地面运费定价（GFP）的费率。必须为GFP启用帐号。
	FRSShipmentIndicator string `json:"FRSShipmentIndicator,omitempty"`
	// 费率表指示器 - 如果请求中存在，响应将包含RateChart元素。
	RateChartIndicator string `json:"RateChartIndicator,omitempty"`
	// TPFC协商价格指标仅适用于第三方/运费到付货物。
	// TPFCNegotiatedRatesIndicator string `json:"TPFCNegotiatedRatesIndicator,omitempty"`
}

func NewShipmentRatingOptions() *ShipmentRatingOptions {
	return &ShipmentRatingOptions{
		NegotiatedRatesIndicator: "Y",
		// FRSShipmentIndicator:     "Y",
		RateChartIndicator: "Y",
	}
}

// 付款账户信息
type PaymentDetails struct {
	ShipmentCharge `json:"ShipmentCharge"`
}

func NewPaymentDetails(shipperNumber string) *PaymentDetails {
	return &PaymentDetails{
		ShipmentCharge{
			Type: "01",
			BillShipper: BillShipper{
				AccountNumber: shipperNumber,
			},
		},
	}
}

type Commodity struct {
	FreightClass string `json:"FreightClass"`
	NMFC         string `json:"NMFC,omitempty"`
}

func NewCommodity50() *Commodity {
	return &Commodity{
		FreightClass: "60",
	}
}

// 包裹服务
type PackageServiceOptions struct {
	// 申报价值-保险
	*DeclaredValue `json:"DeclaredValue,omitempty"`
	// 签名服务
	*DeliveryConfirmation `json:"DeliveryConfirmation,omitempty"`

	// 标识包含危险货物的包装。如果子版本大于或等于1701，则需要危险品装运。
	// PackageIdentifier string `json:"PackageIdentifier,omitempty"`
	// 临床试验的唯一标识符 Unique identifier for clinical trials
	// ClinicalTrialsID string `json:"ClinicalTrialsID,omitempty"`
	// 当包装中的危险品容器数量大于1时需要。它表示所有危险品材料是保存在一个箱子里还是多个箱子里。当包装中的危险品容器数量大于1时需要。
	// HazMatPackageInformation `json:"HazMatPackageInformation,omitempty"`
}
type DeliveryConfirmation struct {
	DCISType string
}

// 成人签名
func NewDeliveryConfirmationAdult() *DeliveryConfirmation {
	return &DeliveryConfirmation{
		DCISType: "3",
	}
}

type ShipmentServiceOptions struct {
	// 估价时的参数
	*ReturnService        `json:"ReturnService,omitempty"`
	*DeliveryConfirmation `json:"DeliveryConfirmation,omitempty"`

	// 打单Shipment时参数
	*LabelDelivery `json:"LabelDelivery,omitempty"`
}
type LabelDelivery struct {
	EMail struct {
		EMailAddress string `json:"EMailAddress"`
	} `json:"EMail,omitempty"`
	LabelLinksIndicator string `json:"LabelLinksIndicator,omitempty"`
}

func NewLabelDelivery(email string) *LabelDelivery {
	return &LabelDelivery{
		EMail: struct {
			EMailAddress string `json:"EMailAddress"`
		}{
			email,
		},
		LabelLinksIndicator: "Y",
	}
}
