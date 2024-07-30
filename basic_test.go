package go_new_ups_api_wrapper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/asrx/go-new-ups-api-wrapper/pojo"
)

var token = "eyJraWQiOiI2NGM0YjYyMC0yZmFhLTQzNTYtYjA0MS1mM2EwZjM2Y2MxZmEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzM4NCJ9.eyJzdWIiOiJ3b3M5MTc2MUBnbWFpbC5jb20iLCJjbGllbnRpZCI6Im4xNkN1ZUpMOGpuc3VYdDQyNVMxdnNxUk9Uc2lETGNhdjdoUVRFbU9RQWpRbGE4QSIsImlzcyI6Imh0dHBzOi8vYXBpcy51cHMuY29tIiwidXVpZCI6IjI3NTlBMzIwLUFENzMtMTk1Mi04NEQ5LUMyNEEzODIxODQ5NSIsInNpZCI6IjY0YzRiNjIwLTJmYWEtNDM1Ni1iMDQxLWYzYTBmMzZjYzFmYSIsImF1ZCI6IlJhdGluZyIsImF0IjoiV2lZQVE2dVpCYU95M3dJaXRRWGpGZ2VQY1VWTiIsIm5iZiI6MTcyMjMzNjk2OSwiRGlzcGxheU5hbWUiOiJDQS1PTlRBUklPIiwiZXhwIjoxNzIyMzUxMzY5LCJpYXQiOjE3MjIzMzY5NjksImp0aSI6IjhhNGEyNWZmLWRlNzgtNDIxNy04ZDAzLWNiYzZlMGRkZGVjYSJ9.alEiA7cluaJhpa1_Coo47eohpyTPD2L-SZ3z9fspK0D9LvUyHArs_s1pFcvu8Yil1aKmOEQzRiWh5n9nej_UFHIos2mXjiHN4zi_WHnBX0XBEgpukY6u0iObep8WjjX2v89wcYlMfCRwyLQxOGW6-tttHlvRROc-jprf_vTq7mueSpNwXGfJs3Pyk2dFK9Gee9g0IGep4dJ2rzZhEA8PAXnUAnAFm1MTThfhKgqwkG1CAjA0r_vAtbn5HA0Nfbon1mIkvwL8iNkslRfjXsTfn_cQISDAINtbmvfxIph5BzWQks6RdgxbVj807rQIl4hNiWrt5ycRMqwAbzihShuj5Ie5CDjVjFf3NYIKiMJUbrWK9Iz_nDy_Okba5TB-HBhBCQn1iGy7Zzh3pEjRoC16mSsJfEkUQXjhm_KHYQXtlgRa8Z7I55c2yK9Yj5qrlcr-r_rgza-G4eFV40JaF_xgzLXwRwW4MVq9z3Kcw86ndtOrpkkJGLcYRKYf-QEGcK1SJEqc7iW7gzmSoVd06yPv8eSfayar-gUuQIiXq91K7IzZ--5fY55FFWVD0NG-btaDwm38bo5xdP3U1CCexT4voex1RK-ZzVU_3GJN7-4MJpjwvoJd_efBmyyhN96Z_-mYmpx3E5_Pe41sbJ4-KwvckpZWpiyeRL1oOAlBHQUTCiQ"
var shipperNumber = "XXXXXX"
var shipper, shipFrom, shipTo = getAddr(shipperNumber)
var debug = true

func getAddr(shipperNumber string) (pojo.Contact, pojo.Contact, pojo.Contact) {

	// 发件地
	shipFrom := pojo.Contact{
		Name:          "ANL",
		AttentionName: "Donovan",
		Address: pojo.Address{
			AddressLine: []string{
				"2440 S. Milliken Avenue",
			},
			City:              "Ontario",
			StateProvinceCode: "CA",
			PostalCode:        "91761",
			CountryCode:       "US",
		},
		Phone: pojo.Phone{
			Number: "6262258083",
		},
		EMailAddress: "donovan.xu@longyuan-sh.com",
	}

	shipper := shipFrom
	shipper.ShipperNumber = shipperNumber

	// 收件人
	shipTo := pojo.Contact{
		Name:          "Lori",
		AttentionName: "Lori.zhou",
		Address: pojo.Address{

			AddressLine: []string{
				"2250 Roswell Dr #200",
			},
			City:              "Pittsburgh",
			StateProvinceCode: "PA",
			PostalCode:        "15205",
			CountryCode:       "US",
			// ResidentialAddressIndicator: "Y",
		},
		Phone: pojo.Phone{
			Number: "8000000000",
		},
		EMailAddress: "lori.zhou@longyuan-sh.com",
	}
	return shipper, shipFrom, shipTo
}

func TestEmptyAttr(t *testing.T) {

	s := pojo.Contact{
		Name:          "",
		AttentionName: "",
		ShipperNumber: "",
		Address:       pojo.Address{},
		Phone:         pojo.Phone{},
		EMailAddress:  "",
	}
	s = pojo.Contact{
		Name:          "ANL",
		AttentionName: "Donovan",
		Address: pojo.Address{
			AddressLine: []string{
				"2440 S. Milliken Avenue",
			},
			City:              "Ontario",
			StateProvinceCode: "CA",
			PostalCode:        "91761",
			CountryCode:       "US",
		},
		Phone: pojo.Phone{
			Number: "6262258083",
		},
		EMailAddress: "donovan.xu@longyuan-sh.com",
	}

	jsonData, _ := json.Marshal(s)
	fmt.Println(string(jsonData)) // 输出：{"address":"123 Main St"}
}
