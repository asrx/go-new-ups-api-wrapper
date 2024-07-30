package pojo

type RequestOption string

const (
	// Rate = The server rates (The default Request option is Rate if a Request Option is not provided).
	// Rate is the only valid request option for UPS Ground Freight Pricing requests.
	RequestOptionRate RequestOption = "Rate"

	// Shop = The server validates the shipment, and returns rates for all UPS products from the ShipFrom to the ShipTo addresses.
	RequestOptionShop RequestOption = "Shop"

	// Ratetimeintransit = The server rates with transit time information
	RequestOptionRatetimeintransit RequestOption = "Ratetimeintransit"

	// Shoptimeintransit = The server validates the shipment, and returns rates and transit times for all UPS products from the ShipFrom to the ShipTo addresses.
	RequestOptionShoptimeintransit RequestOption = "Shoptimeintransit"
)
