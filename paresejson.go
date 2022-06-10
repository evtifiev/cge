package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Парсинг данных из JSON
type Patient struct {
	GUID            string `json:"guid"`
	LUID            string `json:"luid"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Age             string `json:"age"`
	Address         string `json:"address"`
	DatePoison      string `json:"date_poison"`
	DateAffFirst    string `json:"date_aff_first"`
	Diagnosis       string `json:"diagnosis"`
	MedicalHelpName string `json:"medical_help_name"`
	PoisoningDesc   string `json:"poisoning_desc"`
}

// Чтение данных из запроса
func readJSONFromUrlPatient(url string) ([]Patient, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var jsonData []Patient
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

type JsonDataAddress struct {
	Address string `json:"address"`
}

// Чтение данных из запроса
func readJSONFromUrlAddress(guid string) (string, error) {
	url := fmt.Sprintf("https://regiz.gorzdrav.spb.ru/N3.BI/getDData?id=1079&args=%s&auth=9f9208b9-f7e1-4e17-8cfc-a6832e03a12f", guid)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var jsonData []JsonDataAddress
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &jsonData); err != nil {
		return "", err
	}

	var address string
	if len(jsonData) > 0 {
		address = jsonData[0].Address
	}

	return address, nil
}

func getDataBI(db *sql.DB) {
	url := "https://regiz.gorzdrav.spb.ru/N3.BI/getDData?id=1078&args=2022-01-01,2022-01-31&auth=9f9208b9-f7e1-4e17-8cfc-a6832e03a12f"
	jsonData, err := readJSONFromUrlPatient(url)
	if err != nil {
		panic(err)
	}
	for _, patient := range jsonData {
		patient.Diagnosis = dbFindDiagnosis(db, patient.Diagnosis)
		patient.Gender, err = dbFindGender(db, patient.Gender)
		if err != nil {
			fmt.Println(err)
		}
		patient.Address, err = readJSONFromUrlAddress(patient.GUID)
		if err != nil {
			fmt.Println(err)
		}
		dbInsertPatient(db, patient)
	}
}
