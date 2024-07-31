package pojo

// 跟踪响应体
type RespTracking struct {
	TrackResponse *struct {
		Shipment []*struct {
			InquiryNumber string `json:"inquiryNumber"`
			Package       []*struct {
				TrackingNumber string        `json:"trackingNumber"`
				DeliveryDate   []interface{} `json:"deliveryDate"`
				Activity       []*struct {
					Location *struct {
						Address *struct {
							City          string `json:"city,omitempty"`
							StateProvince string `json:"stateProvince,omitempty"`
							CountryCode   string `json:"countryCode,omitempty"`
							Country       string `json:"country,omitempty"`
						} `json:"address"`
						Slic string `json:"slic,omitempty"`
					} `json:"location"`
					Status *struct {
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
				CurrentStatus *struct {
					Description               string `json:"description"`
					Code                      string `json:"code"`
					SimplifiedTextDescription string `json:"simplifiedTextDescription"`
				} `json:"currentStatus"`
				PackageAddress []*struct {
					Type    string `json:"type"`
					Address *struct {
						City          string `json:"city"`
						StateProvince string `json:"stateProvince"`
						CountryCode   string `json:"countryCode"`
						Country       string `json:"country"`
					} `json:"address"`
				} `json:"packageAddress"`
				Weight *struct {
					UnitOfMeasurement string `json:"unitOfMeasurement"`
					Weight            string `json:"weight"`
				} `json:"weight"`
				Service *struct {
					Code        string `json:"code"`
					LevelCode   string `json:"levelCode"`
					Description string `json:"description"`
				} `json:"service"`
				ReferenceNumber []*struct {
					Type   string `json:"type"`
					Number string `json:"number"`
				} `json:"referenceNumber"`
				DeliveryInformation *struct {
					ReceivedBy    string `json:"receivedBy"`
					Location      string `json:"location"`
					DeliveryPhoto *struct {
						PhotoCaptureInd        string `json:"photoCaptureInd"`
						IsNonPostalCodeCountry bool   `json:"isNonPostalCodeCountry"`
					} `json:"deliveryPhoto"`
				} `json:"deliveryInformation"`
				Dimension *struct {
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
