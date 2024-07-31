package pojo

type RespShipment struct {
	ShipmentResponse struct {
		Response struct {
			ResponseStatus struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"ResponseStatus"`
			TransactionReference struct {
				CustomerContext string `json:"CustomerContext"`
			} `json:"TransactionReference"`
		} `json:"Response"`
		ShipmentResults struct {
			ShipmentCharges struct {
				RateChart             string `json:"RateChart"`
				TransportationCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"TransportationCharges"`
				ItemizedCharges []struct {
					Code          string `json:"Code"`
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"ItemizedCharges"`
				ServiceOptionsCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"ServiceOptionsCharges"`
				TotalCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"TotalCharges"`
			} `json:"ShipmentCharges"`
			NegotiatedRateCharges struct {
				TotalCharge struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"TotalCharge"`
			} `json:"NegotiatedRateCharges"`
			BillingWeight struct {
				UnitOfMeasurement struct {
					Code        string `json:"Code"`
					Description string `json:"Description"`
				} `json:"UnitOfMeasurement"`
				Weight string `json:"Weight"`
			} `json:"BillingWeight"`
			ShipmentIdentificationNumber string `json:"ShipmentIdentificationNumber"`
			PackageResults               []struct {
				TrackingNumber    string `json:"TrackingNumber"`
				BaseServiceCharge struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"BaseServiceCharge"`
				ServiceOptionsCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"ServiceOptionsCharges"`
				ShippingLabel struct {
					ImageFormat struct {
						Code        string `json:"Code"`
						Description string `json:"Description"`
					} `json:"ImageFormat"`
					GraphicImage string `json:"GraphicImage"`
					HTMLImage    string `json:"HTMLImage"`
				} `json:"ShippingLabel"`
				ItemizedCharges []struct {
					Code          string `json:"Code"`
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
					SubType       string `json:"SubType,omitempty"`
				} `json:"ItemizedCharges"`
				NegotiatedCharges struct {
					BaseServiceCharge struct {
						CurrencyCode  string `json:"CurrencyCode"`
						MonetaryValue string `json:"MonetaryValue"`
					} `json:"BaseServiceCharge"`
					TransportationCharges struct {
						CurrencyCode  string `json:"CurrencyCode"`
						MonetaryValue string `json:"MonetaryValue"`
					} `json:"TransportationCharges"`
					ServiceOptionsCharges struct {
						CurrencyCode  string `json:"CurrencyCode"`
						MonetaryValue string `json:"MonetaryValue"`
					} `json:"ServiceOptionsCharges"`
					TotalCharge struct {
						CurrencyCode  string `json:"CurrencyCode"`
						MonetaryValue string `json:"MonetaryValue"`
					} `json:"TotalCharge"`
					ItemizedCharges []struct {
						Code          string `json:"Code"`
						CurrencyCode  string `json:"CurrencyCode"`
						MonetaryValue string `json:"MonetaryValue"`
					} `json:"ItemizedCharges"`
				} `json:"NegotiatedCharges"`
				ShippingReceipt struct {
					ImageFormat struct {
						Code        string `json:"Code"`
						Description string `json:"Description"`
					} `json:"ImageFormat"`
					GraphicImage string `json:"GraphicImage"`
				} `json:"ShippingReceipt,omitempty"`
			} `json:"PackageResults"`
			LabelURL   string `json:"LabelURL,omitempty"`
			ReceiptURL string `json:"ReceiptURL,omitempty "`
		} `json:"ShipmentResults"`
	} `json:"ShipmentResponse"`
}
