package pojo

type RespVoid struct {
	VoidShipmentResponse *struct {
		Response *struct {
			ResponseStatus *struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"ResponseStatus"`
			Alert *struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"Alert"`
			TransactionReference *struct {
				CustomerContext string `json:"CustomerContext"`
			} `json:"TransactionReference"`
		} `json:"Response"`
		SummaryResult *struct {
			Status *struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"Status"`
		} `json:"SummaryResult"`
		PackageLevelResult []*struct {
			TrackingNumber string `json:"TrackingNumber"`
			Status         *struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"Status"`
		} `json:"PackageLevelResult"`
	} `json:"VoidShipmentResponse"`
}
