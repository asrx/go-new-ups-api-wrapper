package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTracking(t *testing.T) {
	tracksId := "1ZC1E8120304128073"
	tracking, err := Tracking(token, tracksId, debug)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tracking.TrackResponse.Shipment[0].InquiryNumber)

	data, err := json.Marshal(tracking)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
