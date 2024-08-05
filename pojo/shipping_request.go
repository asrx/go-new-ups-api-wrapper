package pojo

type RequestShipping struct {
	ShipmentRequest `json:"ShipmentRequest"`
}

type ShipmentRequest struct {
	*Request            `json:"Request,omitempty"`
	*Shipment           `json:"Shipment"`
	*LabelSpecification `json:"LabelSpecification,omitempty"`
}

type Shipment struct {
	Description                    string          `json:"Description,omitempty"`
	Shipper                        Contact         `json:"Shipper"`
	ShipTo                         Contact         `json:"ShipTo"`
	ShipFrom                       Contact         `json:"ShipFrom,omitempty"`
	PaymentInformation             *PaymentDetails `json:"PaymentInformation"`
	*FRSPaymentInformation         `json:"FRSPaymentInformation,omitempty"`
	*Service                       `json:"Service"`
	*ShipmentRatingOptions         `json:"ShipmentRatingOptions"`
	RatingMethodRequestedIndicator string         `json:"RatingMethodRequestedIndicator,omitempty"`
	Package                        []*ShipPackage `json:"Package"`
	*ReturnService                 `json:"ReturnService,omitempty"`
	*ShipmentServiceOptions        `json:"ShipmentServiceOptions,omitempty"`
	//
}

// 退件服务
type ReturnService struct {
	Code        string `json:"Code"`
	Description string `json:"Description,omitempty"`
}

func NewReturnServiceERL() *ReturnService {
	return &ReturnService{
		Code:        "8",
		Description: "ERL",
	}
}
func NewReturnServicePRL() *ReturnService {
	return &ReturnService{
		Code:        "9",
		Description: "PRL",
	}
}

// 尺寸信息容器。注：目前尺寸不适用于地面运费定价。
// 长度+2*（宽度+高度）必须小于或等于165英寸或330厘米。重型货物服务需要。简单费率将忽略包装尺寸
type ShipPackage struct {
	Description            string         `json:"Description"`
	Packaging              *PackagingType `json:"Packaging"`
	*Dimensions            `json:"Dimensions,omitempty"`
	*PackageWeight         `json:"PackageWeight"`
	ReferenceNumber        []*ReferenceNumber `json:"ReferenceNumber,omitempty"`
	*PackageServiceOptions `json:"PackageServiceOptions,omitempty"`
	*Commodity             `json:"Commodity,omitempty"`
}

// 参考号
type ReferenceNumber struct {
	BarCodeIndicator string `json:"BarCodeIndicator,omitempty"`
	Code             string `json:"Code,omitempty"`
	Value            string `json:"Value"`
}

func NewReferenceNumbers(value ...string) []*ReferenceNumber {
	referenceNumbers := make([]*ReferenceNumber, 0)
	code := ""
	for i, v := range value {
		if v == "" {
			continue
		}
		if i == 0 {
			code = "PC"
		} else {
			code = "DP"
		}
		referenceNumbers = append(referenceNumbers, &ReferenceNumber{
			Code:  code,
			Value: v,
		})
	}
	if len(referenceNumbers) == 0 {
		return nil
	}
	return referenceNumbers
}

// 申报价值
type DeclaredValue struct {
	CurrencyCode  string `json:"CurrencyCode"`
	MonetaryValue string `json:"MonetaryValue"`
}

func NewDeclaredValue(value string) *DeclaredValue {
	return &DeclaredValue{
		CurrencyCode:  "USD",
		MonetaryValue: value,
	}
}

// 标签规格
type LabelSpecification struct {
	*LabelImageFormat `json:"LabelImageFormat"`
	// *HTTPUserAgent     string `json:"HTTPUserAgent,omitempty"`
	*LabelStockSize `json:"LabelStockSize"`
}

func NewLabelSpecificationGif() *LabelSpecification {
	return &LabelSpecification{
		LabelImageFormat: &LabelImageFormat{
			Code:        "GIF",
			Description: "GIF",
		},
		LabelStockSize: &LabelStockSize{
			Width:  "4",
			Height: "6",
		},
	}
}
func NewLabelSpecificationZPL() *LabelSpecification {
	return &LabelSpecification{
		LabelImageFormat: &LabelImageFormat{
			Code:        "ZPL",
			Description: "ZPL",
		},
		LabelStockSize: &LabelStockSize{
			Width:  "4",
			Height: "6",
		},
	}
}

// 标签类型
type LabelImageFormat struct {
	// 标签打印方法代码决定了生成标签的格式。对于EPL2格式的标签，使用EPL，对于SPL格式的标签使用SPL，对于ZPL格式的标签用ZPL，对于图像格式使用GIF。对于没有退货服务的货物，有效值为GIF、ZPL、EPL和SPL。对于提供PRL退货服务的货物，有效值为EPL、ZPL、SPL和GIF。对于UPS Premier Silver发货，仅支持ZPL。
	Code        string `json:"Code"`
	Description string `json:"Description,omitempty"`
}

// 标签尺寸大小
type LabelStockSize struct {
	// 标签图像的宽度。对于IN，使用整英寸。适用于EPL2、ZPL和SPL标签。有效值为4。注意：即使请求4 X 8，标签图像也只能缩放到4 X 6。
	Width string `json:"Width"`
	// 标签图像的高度。对于IN，使用整英寸。适用于EPL2、ZPL和SPL标签。只有6或8是有效值。注意：即使请求4 X 8，标签图像也只能缩放到4 X 6。
	Height string `json:"Height"`
}
