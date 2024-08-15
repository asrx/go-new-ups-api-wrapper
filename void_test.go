package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

func TestVoid(t *testing.T) {
	shipmentidentificationnumber := "1ZXY3554A893282353"
	nums := []string{
		"1ZXY3554A893282353",
	}
	void, err := Void(token, shipmentidentificationnumber, nums, debug)
	if err != nil {
		fmt.Println("Void is error: ", err)
		return
	}
	data, _ := json.Marshal(void)
	fmt.Println(string(data))
}

func TestRepy(t *testing.T) {
	str := `{"VoidShipmentResponse":{"Response":{"ResponseStatus":{"Code":"1", "Description":"Success"}}, "SummaryResult":{"Status":{"Code":"1", "Description":"Partially Voided"}}, "PackageLevelResult":{"TrackingNumber":"1ZXY3554A893282353", "Status":{"Code":"1", "Description":"Already Voided"}}}}`
	result := &pojo.RespVoid{}
	err := json.Unmarshal([]byte(str), result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
}
