package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

func TestRatingGround(t *testing.T) {

	packages := getRateGroundPkgs()

	rating, err := Rating(token, shipper, shipFrom, shipTo, packages, pojo.UPS_GROUND, debug)
	if err != nil {
		fmt.Printf("Rating is error: %s\n", err.Error())
		return
	}
	data, _ := json.Marshal(rating)
	fmt.Println(string(data))
}

func TestRatingGFP(t *testing.T) {
	packages := getRateGfpPkgs()

	_, err := Rating(token, shipper, shipFrom, shipTo, packages, pojo.UPS_GFP, debug)
	if err != nil {
		fmt.Printf("Rating is error: %s\n", err.Error())
		return
	}
}

func TestRatingHundredWeight(t *testing.T) {

}

func getRateGroundPkgs() []*pojo.RatePackage {
	packageCount := 1

	packages := make([]*pojo.RatePackage, 0)
	weight := "30"
	length := "16"
	width := "13"
	height := "10"
	pkgType := pojo.PackageType02_Package()
	unitDis := pojo.NewUnitDimesions()
	unitWeight := pojo.NewUnitWeight()

	declaredValue := ""
	signatureServiceFlag := true
	for i := 0; i < packageCount; i++ {
		packages = append(packages,
			pojo.BuildRatePackage(weight, length, width, height, pkgType, unitDis, unitWeight, declaredValue, signatureServiceFlag),
		)
	}
	return packages
}
func getRateGfpPkgs() []*pojo.RatePackage {
	packageCount := 20

	packages := make([]*pojo.RatePackage, 0)
	weight := "30"
	length := "16"
	width := "13"
	height := "10"
	pkgType := pojo.PackageType02_Package()
	unitDis := pojo.NewUnitDimesions()
	unitWeight := pojo.NewUnitWeight()

	declaredValue := ""
	signatureServiceFlag := true
	for i := 0; i < packageCount; i++ {
		p := pojo.BuildRatePackage(weight, length, width, height, pkgType, unitDis, unitWeight, declaredValue, signatureServiceFlag)
		p.Commodity = pojo.NewCommodity50()
		packages = append(packages, p)
	}
	return packages
}
