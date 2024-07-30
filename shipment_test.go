package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

func TestShipment(t *testing.T) {

	packages := getShipGroundPkg()

	shipResp, err := Shipment(token, shipper, shipFrom, shipTo, packages,
		pojo.UPS_GROUND, debug)

	if err != nil {
		fmt.Printf("Shipment is error: %s\n", err.Error())
		return
	}
	data, _ := json.Marshal(shipResp)
	fmt.Println(string(data))
}

func TestReturnShipment(t *testing.T) {
	packages := getShipGroundPkg()

	_, err := ShipmentReturnService(token, shipper, shipFrom, shipTo, packages,
		pojo.UPS_GROUND, "donovan.xu@longyuan-sh.com", debug)

	if err != nil {
		fmt.Printf("Return Shipment is error: %s\n", err.Error())
		return
	}
}

func getShipGroundPkg() []*pojo.ShipPackage {
	packageCount := 3

	packages := make([]*pojo.ShipPackage, 0)
	weight := "30"
	length := "16"
	width := "13"
	height := "10"
	pkgType := pojo.PackageType02_Package()
	unitDis := pojo.NewUnitDimesions()
	unitWeight := pojo.NewUnitWeight()

	// 申报价值
	declaredValue := ""
	signatureServiceFlag := false
	// 参考号
	referenceNumbers := pojo.NewReferenceNumbers("111ADC", "222TOP")

	for i := 0; i < packageCount; i++ {
		packages = append(packages,
			pojo.BuildShipPackage(weight, length, width, height, pkgType, unitDis, unitWeight, declaredValue, referenceNumbers, signatureServiceFlag),
		)
	}
	return packages
}
