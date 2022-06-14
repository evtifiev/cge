package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const usageText = `Программа генерации отчетов для ФБУЗ "Центр гигиены и эпидемиологии в городе Санкт-Петербург".
 
Применение:
 
cge [аргументы]
 
Доступные аргументы:
- startDate - дата начала выборки
- endDate - дата окончания выборки
- fileName - имя файла
 
`

func main() {
	flag.Usage = usage
	var startDate, endDate string
	flag.StringVar(&startDate, "startDate", "", "дата начала выборки")
	flag.StringVar(&endDate, "endDate", "", "дата окончания выборки")
	flag.Parse()
	if len(startDate) == 0 || len(endDate) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	db := dbGetConn()
	defer db.Close()
	// Загрузка данных
	getDataBI(db, startDate, endDate)
	patients := dbDisplayPatients(db)

	//dbDisplayGenders(db)

	status := createXML(patients)
	if status {
		log.Println("Создал XML файл ... ")
		dbClearTablePatient(db)
		log.Println("Работа завершена, жду другую команду ... ")
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
