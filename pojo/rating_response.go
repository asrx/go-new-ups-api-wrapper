package pojo

type RespRate struct {
	RateResponse struct {
		Response struct {
			ResponseStatus struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"ResponseStatus"`
			Alert []struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"Alert"`
			TransactionReference struct {
				CustomerContext string `json:"CustomerContext"`
			} `json:"TransactionReference"`
		} `json:"Response"`
		RatedShipment struct {
			Service struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"Service"`
			RateChart          string `json:"RateChart"`
			RatedShipmentAlert []struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"RatedShipmentAlert"`
			BillingWeight struct {
				UnitOfMeasurement struct {
					Code        string `json:"Code"`
					Description string `json:"Description"`
				} `json:"UnitOfMeasurement"`
				Weight string `json:"Weight"`
			} `json:"BillingWeight"`
			TransportationCharges struct {
				CurrencyCode  string `json:"CurrencyCode"`
				MonetaryValue string `json:"MonetaryValue"`
			} `json:"TransportationCharges"`
			BaseServiceCharge struct {
				CurrencyCode  string `json:"CurrencyCode"`
				MonetaryValue string `json:"MonetaryValue"`
			} `json:"BaseServiceCharge"`
			ServiceOptionsCharges struct {
				CurrencyCode  string `json:"CurrencyCode"`
				MonetaryValue string `json:"MonetaryValue"`
			} `json:"ServiceOptionsCharges"`
			TotalCharges struct {
				CurrencyCode  string `json:"CurrencyCode"`
				MonetaryValue string `json:"MonetaryValue"`
			} `json:"TotalCharges"`
			NegotiatedRateCharges struct {
				BaseServiceCharge struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"BaseServiceCharge"`
				TotalCharge struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"TotalCharge"`
			} `json:"NegotiatedRateCharges"`
			RatedPackage []struct {
				TransportationCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"TransportationCharges"`
				BaseServiceCharge struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"BaseServiceCharge"`
				ServiceOptionsCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"ServiceOptionsCharges"`
				ItemizedCharges []struct {
					Code          string `json:"Code"`
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
					SubType       string `json:"SubType,omitempty"`
				} `json:"ItemizedCharges"`
				TotalCharges struct {
					CurrencyCode  string `json:"CurrencyCode"`
					MonetaryValue string `json:"MonetaryValue"`
				} `json:"TotalCharges"`
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
				Weight        string `json:"Weight"`
				BillingWeight struct {
					UnitOfMeasurement struct {
						Code        string `json:"Code"`
						Description string `json:"Description"`
					} `json:"UnitOfMeasurement"`
					Weight string `json:"Weight"`
				} `json:"BillingWeight"`
			} `json:"RatedPackage"`
		} `json:"RatedShipment"`
	} `json:"RateResponse"`
}
