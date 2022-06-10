package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Сохранение данных в xml
const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n"
)

// Строка объекта данных
type V struct {
	XMLName xml.Name `xml:"v"`
	F       string   `xml:"f,attr"`
	D       string   `xml:",chardata"`
}

// Объект данных
type R struct {
	XMLName xml.Name `xml:"r"`
	V       []V
}

// Структура data - содержимое выгрузки
type Data struct {
	XMLName xml.Name `xml:"data"`
	R       []R
}

// Структура field - описание типа полей
type F struct {
	XMLName xml.Name `xml:"f"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	T       string   `xml:"t,attr"`
	S       string   `xml:"s,attr,omitempty"`
	R       string   `xml:"r,attr,omitempty"`
}

// Структура fields - описание полей с типом
type Fields struct {
	XMLName xml.Name `xml:"fields"`
	F       []F
}

// Структура dataset
type Dataset struct {
	XMLName xml.Name `xml:"dataset"`
	Id      string   `xml:"id,attr"`
	Fields  Fields
	Data    Data
}

// Структура info
type Info struct {
	XMLName     xml.Name `xml:"info"`
	GUID        string   `xml:"GUID"`
	VersionStat string   `xml:"versionStat"`
	Version     string   `xml:"version"`
}

// Структура верхнего уровня
type Package struct {
	XMLName xml.Name `xml:"package"`
	Info    Info
	Dataset Dataset
}

func genFields() *Fields {
	fields := &Fields{}
	fields.F = append(fields.F, F{Id: "1", Name: "id", T: "Integer"})
	fields.F = append(fields.F, F{Id: "2", Name: "NUMNOTICE", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "3", Name: "NAME", T: "String", S: "50", R: "True"})
	fields.F = append(fields.F, F{Id: "4", Name: "C_Gender", T: "Integer"})
	fields.F = append(fields.F, F{Id: "5", Name: "Age", T: "Integer"})
	fields.F = append(fields.F, F{Id: "6", Name: "C_Social", T: "Integer"})
	fields.F = append(fields.F, F{Id: "7", Name: "C_Region", T: "Integer"})
	fields.F = append(fields.F, F{Id: "8", Name: "Note", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "9", Name: "C_PlaceIncident", T: "Integer"})
	fields.F = append(fields.F, F{Id: "10", Name: "Note3", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "11", Name: "DatePoison", T: "Integer"})
	fields.F = append(fields.F, F{Id: "12", Name: "DateFirstRecourse", T: "Integer"})
	fields.F = append(fields.F, F{Id: "13", Name: "DateAffFirst", T: "Integer"})
	fields.F = append(fields.F, F{Id: "14", Name: "C_Diagnosis", T: "Integer"})
	fields.F = append(fields.F, F{Id: "15", Name: "C_BooleanAlc", T: "Integer"})
	fields.F = append(fields.F, F{Id: "16", Name: "C_SetDiagnosis", T: "Integer"})
	fields.F = append(fields.F, F{Id: "17", Name: "C_MedicalHelp", T: "Integer"})
	fields.F = append(fields.F, F{Id: "18", Name: "C_PlaceMortality", T: "Integer"})
	fields.F = append(fields.F, F{Id: "19", Name: "Note5", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "20", Name: "C_TypePoison", T: "Integer"})
	fields.F = append(fields.F, F{Id: "21", Name: "ValPoison", T: "Float"})
	fields.F = append(fields.F, F{Id: "22", Name: "C_AimPoison", T: "Integer"})
	fields.F = append(fields.F, F{Id: "23", Name: "Note7", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "24", Name: "C_PlacePoison", T: "Integer"})
	fields.F = append(fields.F, F{Id: "25", Name: "Note8", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "26", Name: "DateDocument", T: "Integer"})
	fields.F = append(fields.F, F{Id: "27", Name: "NAMEPEOPLEGET", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "28", Name: "CREATEUSER", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "29", Name: "CREATEDATE", T: "DateTime"})
	fields.F = append(fields.F, F{Id: "30", Name: "UPDATEUSER", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "31", Name: "UPDATEDATE", T: "DateTime"})
	fields.F = append(fields.F, F{Id: "32", Name: "FlagColor", T: "Integer"})
	fields.F = append(fields.F, F{Id: "33", Name: "C_GSEN", T: "Integer"})
	fields.F = append(fields.F, F{Id: "34", Name: "S_OBJECTMESS", T: "Integer"})
	fields.F = append(fields.F, F{Id: "35", Name: "S_OBJECTMESSNAME", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "36", Name: "S_STREET", T: "Integer"})
	fields.F = append(fields.F, F{Id: "37", Name: "S_STREETNAME", T: "String", S: "255"})
	fields.F = append(fields.F, F{Id: "38", Name: "HOUSE", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "39", Name: "FLAT", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "40", Name: "DateLock", T: "Integer"})
	fields.F = append(fields.F, F{Id: "41", Name: "S_ObjectMedicalHelp", T: "Integer"})
	fields.F = append(fields.F, F{Id: "42", Name: "S_ObjectMedicalHelpName", T: "String", S: "50"})
	fields.F = append(fields.F, F{Id: "43", Name: "errorfontcolor", T: "Integer"})
	fields.F = append(fields.F, F{Id: "44", Name: "errorfontstyle", T: "Integer"})
	fields.F = append(fields.F, F{Id: "45", Name: "errorcolor", T: "Integer"})
	fields.F = append(fields.F, F{Id: "46", Name: "errorcolfontcolor", T: "Integer"})
	fields.F = append(fields.F, F{Id: "47", Name: "errorcolfontstyle", T: "Integer"})
	fields.F = append(fields.F, F{Id: "48", Name: "errorcolcolor", T: "Integer"})
	fields.F = append(fields.F, F{Id: "48", Name: "errortext", T: "String", S: "254"})
	fields.F = append(fields.F, F{Id: "50", Name: "errorcolumns", T: "String", S: "254"})
	fields.F = append(fields.F, F{Id: "51", Name: "CANREADONLY", T: "SmallInt"})
	fields.F = append(fields.F, F{Id: "52", Name: "CANEDITONLY", T: "SmallInt"})
	fields.F = append(fields.F, F{Id: "53", Name: "CANDELETEONLY", T: "SmallInt"})
	return fields
}

func createXML(patients []Patient) bool {
	status := true
	// Формируем файл xml
	v := &Package{}
	// Добавляем в этот файл информацию
	v.Info = Info{GUID: "{A4F6D1E0-909A-11D5-B08F-000021EF6307}", VersionStat: "404064", Version: "404064"}
	v.Dataset.Id = "CaptionDB"
	v.Dataset.Fields = *genFields()

	// Генерируем данные
	rows := []R{}
	for _, patient := range patients {
		r := R{}
		r.V = append(r.V, V{F: "3", D: patient.Name})   // ФИО
		r.V = append(r.V, V{F: "4", D: patient.Gender}) // Пол
		age, _ := strconv.Atoi(patient.Age)
		age *= 1000
		ageStr := strconv.Itoa(age)
		r.V = append(r.V, V{F: "5", D: ageStr})
		r.V = append(r.V, V{F: "8", D: patient.Address})
		r.V = append(r.V, V{F: "11", D: strings.ReplaceAll(patient.DatePoison, "-", "")})
		r.V = append(r.V, V{F: "13", D: strings.ReplaceAll(patient.DateAffFirst, "-", "")})
		r.V = append(r.V, V{F: "14", D: patient.Diagnosis})
		r.V = append(r.V, V{F: "23", D: patient.PoisoningDesc})
		r.V = append(r.V, V{F: "42", D: patient.MedicalHelpName})
		rows = append(rows, r)
	}
	v.Dataset.Data.R = rows

	filename := "test.xml"
	xmlFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании XML: ", err)
		status = false
		return status
	}
	xmlFile.WriteString(Header)
	xmlWriter := io.Writer(xmlFile)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent(" ", " ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return status
}
