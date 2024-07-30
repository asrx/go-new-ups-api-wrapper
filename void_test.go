package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVoid(t *testing.T) {
	shipmentidentificationnumber := "1ZXY35540399177842"
	nums := []string{
		"1ZXY35540399177842",
		"1ZXY35540396191859",
		"1ZXY35540395727666",
	}
	void, err := Void(token, shipmentidentificationnumber, nums, debug)
	if err != nil {
		fmt.Println("Void is error: ", err)
		return
	}
	data, _ := json.Marshal(void)
	fmt.Println(string(data))
}
