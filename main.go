package main

import (
	"fmt"
)

func main() {
	fmt.Println("Система учета пациентов")

	db := dbGetConn()
	defer db.Close()
	// Загрузка данных
	getDataBI(db)
	//dbClearTablePatient(db)
	patients := dbDisplayPatients(db)

	//dbDisplayGenders(db)

	createXML(patients)
	// if status {
	// 	dbClearTablePatient(db)
	// }
}
