package pojo

func Service01_NextDayAir() *Service {
	return &Service{
		Code:        "01",
		Description: "Next Day Air",
	}
}
func Service02_2ndDayAir() *Service {
	return &Service{
		Code:        "02",
		Description: "2nd Day Air",
	}
}
func Service03_Ground() *Service {
	return &Service{
		Code:        "03",
		Description: "Ground",
	}
}
func Service12_3DaySelect() *Service {
	return &Service{
		Code:        "12",
		Description: "3 Day Select",
	}
}
func Service13_NextDayAirSaver() *Service {
	return &Service{
		Code:        "13",
		Description: "Next Day Air Saver",
	}
}
func Service14_UpsNextDayAirEarly() *Service {
	return &Service{
		Code:        "14",
		Description: "UPS Next Day Air Early",
	}
}
func Service59_2ndDayAirAM() *Service {
	return &Service{
		Code:        "59",
		Description: "2nd Day Air A.M.",
	}
}
func Service75_UpsHeavyGoods() *Service {
	return &Service{
		Code:        "75",
		Description: "UPS Heavy Goods",
	}
}
