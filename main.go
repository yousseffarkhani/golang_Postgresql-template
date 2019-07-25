package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	host       = "db"
	port       = "5432"
	user       = "postgres"
	password   = "secret"
	dbname     = "basket"
)

type Court struct {
	gorm.Model
	Name           string `json:"nom"`
	Url            string `json:"url"`
	Adress         string `json:"adresse"`
	Arrondissement string `json:"arrondissement"`
	Longitude      string `json:"longitude"`
	Lattitude      string `json:"lattitude"`
	Dimensions     string `json:"dimensions"`
	Revetement     string `json:"revetement"`
	Decouvert      string `json:"decouvert"`
	Eclairage      string `json:"eclairage"`
}

func main() {
	if env := os.Getenv("APP_ENV"); env == "production" {
		fmt.Println("Running in production mode")
	} else {
		fmt.Println("Running in developpment mode")
	}

	db, err := InitialMigration()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Open("assets/courts.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var courts []Court
	err = decoder.Decode(&courts)
	if err != nil {
		log.Fatalln(err)
	}

	for _, court := range courts {
		AddCourt(db, court)
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		courtss := GetAllCourts(db)
		json.NewEncoder(w).Encode(courtss)
	})

	fs := http.FileServer(http.Dir(""))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

func InitialMigration() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	err := Reset(driverName, psqlInfo, dbname)
	if err != nil {
		return nil, err
	}
	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := gorm.Open(driverName, psqlInfo)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Court{})
	return db, nil
}

func Reset(driverName, dataSource, dbname string) error {
	db, err := gorm.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	resetDB(db, dbname)

	return db.Close()
}

func resetDB(db *gorm.DB, name string) {
	db.Exec("DROP DATABASE IF EXISTS " + name)
	createDB(db, name)
}

func createDB(db *gorm.DB, name string) {
	db.Exec("CREATE DATABASE " + name)
}

func GetAllCourts(db *gorm.DB) []Court {
	var courts []Court
	db.Find(&courts)
	return courts
}

func AddCourt(db *gorm.DB, court Court) {
	db.Create(&Court{
		Name:           court.Name,
		Url:            court.Url,
		Adress:         court.Adress,
		Arrondissement: court.Arrondissement,
		Longitude:      court.Longitude,
		Lattitude:      court.Lattitude,
		Dimensions:     court.Dimensions,
		Revetement:     court.Revetement,
		Decouvert:      court.Decouvert,
		Eclairage:      court.Eclairage,
	})
	fmt.Println("New court successfully created")
}
