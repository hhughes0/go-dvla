package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"bytes"
	"io/ioutil"
	"os"
	"encoding/json"
)

type details struct {
	TypeApproval string
	YearOfManufacture int
	TaxDueDate string
	MonthOfFirstRegistration string
	Make string
	TaxStatus string
	RegistrationNumber string
	Co2Emissions  int
	MotStatus string
	RevenueWeight int
	DateOfLastV5CIssued string
	EngineCapacity int
	EuroStatus string
	MarkedForExport bool
	FuelType string
	Colour string
	Wheelplan string
}

func main() {
	plate := os.Args[1]
	dvla_info := Get_dvla(plate)
	fmt.Println(fmt.Sprintf("Car Make is: %v", dvla_info.Make))
	fmt.Println(fmt.Sprintf("Car Colour is: %v", dvla_info.Colour))
	fmt.Println(fmt.Sprintf("Car Tax due date is: %v", dvla_info.TaxDueDate))
}

func Get_dvla(plate string) details {
	var car details
	apiUrl := "https://driver-vehicle-licensing.api.gov.uk"
	resource := "/vehicle-enquiry/v1/vehicles"
	plate_string := fmt.Sprintf(`{"registrationNumber":"%s"}`, plate)
	jsonStr := []byte(plate_string)

	u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := u.String()

	client := &http.Client{}
    r, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(jsonStr))
	r.Header.Add("x-api-key", "123456789") //Replace with user-specific, provided API key
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Accept-Encoding", "gzip, deflate, br")
	r.Header.Add("Accept", "*/*")

	resp, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}
	
	if resp.Body != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal([]byte(body), &car)
		if err != nil {
			log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
		}
	}
	return car
}
