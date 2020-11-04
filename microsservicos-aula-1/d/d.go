package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Coupon struct {
	Code string
}

func checkBeauty(code string) string {
	if strings.Contains(code, "a") {
		return "and beatiful"
	}

	return "but ugly"
}

type Result struct {
	Status string
}

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	beauty := checkBeauty(coupon)

	result := Result{Status: beauty}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}
