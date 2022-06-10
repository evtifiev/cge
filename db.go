package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Инициализация БД
func dbGetConn() *sql.DB {
	db, err := sql.Open("sqlite3", "./cge.db") // Open the created SQLite File
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	} // Defer Closing the database
	return db
}

// Поиск кода МКБ
func dbFindDiagnosis(db *sql.DB, code string) string {
	var findCode int
	row, err := db.Query("SELECT code FROM diagnosis ORDER BY diagnosis_mkb_code4 = ? LIMIT 1", code)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&findCode)
	}
	return strconv.Itoa(findCode)
}

// Поиск кода пола
func dbFindGender(db *sql.DB, gender string) (string, error) {
	var findGender int
	row, err := db.Query("SELECT code FROM genders WHERE name = ?", gender)
	if err != nil {
		return "", err
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&findGender)
	}
	return strconv.Itoa(findGender), nil
}

// Вывод списка
func dbDisplayPatients(db *sql.DB) []Patient {
	patients := []Patient{}
	row, err := db.Query("SELECT name, gender, age, address, date_poison, date_aff_first, diagnosis, poisoning_desc, medical_help_name FROM patients")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var patient Patient
		row.Scan(&patient.Name, &patient.Gender, &patient.Age, &patient.Address, &patient.DatePoison, &patient.DateAffFirst,
			&patient.Diagnosis, &patient.PoisoningDesc, &patient.MedicalHelpName)
		patients = append(patients, patient)
	}
	return patients
}

func dbInsertPatient(db *sql.DB, patient Patient) {
	log.Println("Inserting patient record ... ", patient.GUID, "    ", patient.Gender)
	insertPatientSQL := `INSERT INTO patients(
		guid, luid, name, gender, age, address, date_poison, date_aff_first,
		diagnosis, poisoning_desc, medical_help_name) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertPatientSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	var age, gender int
	if patient.Age == "" {
		age = 0
	} else {
		age, _ = strconv.Atoi(patient.Age)
	}
	if patient.Gender == "" {
		gender = 0
	} else {
		gender, _ = strconv.Atoi(patient.Gender)
	}
	diagnosis, _ := strconv.Atoi(patient.Diagnosis)
	_, err = statement.Exec(
		&patient.GUID,
		&patient.LUID,
		&patient.Name,
		gender,
		age,
		&patient.Address,
		&patient.DatePoison,
		&patient.DateAffFirst,
		diagnosis,
		&patient.PoisoningDesc,
		&patient.MedicalHelpName,
	)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func dbClearTablePatient(db *sql.DB) {
	log.Println("Delete pacient record ...")
	deletePatientsSQL := `DELETE FROM patients; VACUUM;`
	statement, err := db.Prepare(deletePatientsSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	statement.Exec()
}
