package pojo

func PackageType00_UNKNOWN() *PackagingType {
	return &PackagingType{
		Code:        "00",
		Description: "UNKNOWN",
	}
}
func PackageType01_UpsLetter() *PackagingType {
	return &PackagingType{
		Code:        "01",
		Description: "UpsLetter",
	}
}
func PackageType02_Package() *PackagingType {
	return &PackagingType{
		Code:        "02",
		Description: "Package",
	}
}
func PackageType03_Tube() *PackagingType {
	return &PackagingType{
		Code:        "03",
		Description: "Tube",
	}
}
func PackageType04_Pak() *PackagingType {
	return &PackagingType{
		Code:        "04",
		Description: "Pak",
	}
}
func PackageType21_ExpressBox() *PackagingType {
	return &PackagingType{
		Code:        "21",
		Description: "ExpressBox",
	}
}
func PackageType24_25kgBox() *PackagingType {
	return &PackagingType{
		Code:        "24",
		Description: "25kgBox",
	}
}
func PackageType25_10kgBox() *PackagingType {
	return &PackagingType{
		Code:        "25",
		Description: "10kgBox",
	}
}
func PackageType30_Pallet() *PackagingType {
	return &PackagingType{
		Code:        "30",
		Description: "Pallet",
	}
}
func PackageType2a_SmallExpressBox() *PackagingType {
	return &PackagingType{
		Code:        "2a",
		Description: "SmallExpressBox",
	}
}
func PackageType2b_MediumExpressBox() *PackagingType {
	return &PackagingType{
		Code:        "2b",
		Description: "MediumExpressBox",
	}
}
func PackageType2c_LargeExpressBox() *PackagingType {
	return &PackagingType{
		Code:        "2c",
		Description: "LargeExpressBox",
	}
}

func NewUnitDimesions() *UnitOfMeasurement {
	return &UnitOfMeasurement{
		Code:        "IN",
		Description: "Inches",
	}
}
func NewUnitWeight() *UnitOfMeasurement {
	return &UnitOfMeasurement{
		Code:        "LBS",
		Description: "Pounds",
	}
}

// 尺寸容器。此容器不适用于GFP评级请求。重型货物服务所需。简单费率将忽略包装尺寸
func BuildDimensions(length, width, height string, unit *UnitOfMeasurement) *Dimensions {
	return &Dimensions{
		UnitOfMeasurement: unit,
		Length:            length,
		Width:             width,
		Height:            height,
	}
}

func BuildWeight(weight string, unitWeight *UnitOfMeasurement) *PackageWeight {
	return &PackageWeight{
		UnitOfMeasurement: unitWeight,
		Weight:            weight,
	}
}

func BuildRatePackage(weight, length, width, height string,
	packagingType *PackagingType,
	unitDimension *UnitOfMeasurement,
	unitWeight *UnitOfMeasurement, declaredValue string, signatureServiceFlag bool) *RatePackage {
	p := &RatePackage{
		PackagingType: packagingType,
		Dimensions:    BuildDimensions(length, width, height, unitDimension),
		PackageWeight: BuildWeight(weight, unitWeight),
	}
	if len(declaredValue) > 0 {
		p.PackageServiceOptions = &PackageServiceOptions{
			DeclaredValue: NewDeclaredValue(declaredValue),
			// PackageIdentifier: "",
			// ClinicalTrialsID:  "",
		}
	}
	if signatureServiceFlag {
		if p.PackageServiceOptions == nil {
			p.PackageServiceOptions = &PackageServiceOptions{
				DeliveryConfirmation: NewDeliveryConfirmationAdult(),
			}
		} else {
			p.PackageServiceOptions.DeliveryConfirmation = NewDeliveryConfirmationAdult()
		}
	}
	return p
}

func BuildShipPackage(weight, length, width, height string,
	packagingType *PackagingType, unitDimension, unitWeight *UnitOfMeasurement,
	declaredValue string,
	referenceNumber []*ReferenceNumber, signatureServiceFlag bool) *ShipPackage {
	p := &ShipPackage{
		Description:   "desc",
		Packaging:     packagingType,
		Dimensions:    BuildDimensions(length, width, height, unitDimension),
		PackageWeight: BuildWeight(weight, unitWeight),
	}
	if referenceNumber != nil {
		p.ReferenceNumber = referenceNumber
	}
	if len(declaredValue) > 0 {
		p.PackageServiceOptions = &PackageServiceOptions{
			DeclaredValue: NewDeclaredValue(declaredValue),
			// PackageIdentifier: "",
			// ClinicalTrialsID:  "",
		}
	}
	if signatureServiceFlag {
		if p.PackageServiceOptions == nil {
			p.PackageServiceOptions = &PackageServiceOptions{
				DeliveryConfirmation: NewDeliveryConfirmationAdult(),
			}
		} else {
			p.PackageServiceOptions.DeliveryConfirmation = NewDeliveryConfirmationAdult()
		}
	}
	return p
}
