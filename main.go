package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asd123"
	dbname   = "pgadmin"
)

type Employee struct {
	ID         int `gorm:"primary_key`
	FirstName  string
	LastName   string
	Email      string
	City       string
	Department Department `gorm:"foreignKey:HeadDepartment"`
	gorm.Model
}
type Department struct {
	Name           string
	HeadDepartment int
	gorm.Model
}
type Joined struct {
	Name           string
	DepartmentName string
}

func main() {
	dns := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Employee{}, &Department{})

	data := []Employee{
		{
			FirstName:  "John",
			LastName:   "Cena",
			Email:      "asd@gmail.com",
			City:       "Indonesia",
			Department: Department{Name: "Creative"},
		},
		{
			FirstName:  "Jane",
			LastName:   "cena",
			Email:      "asd123@gmail.com",
			City:       "Indonesia",
			Department: Department{Name: "Komputer"},
		},
	}
	db.Create(&data)
	rows := []Joined{}
	db.Model(&Department{}).Select("employees.first_name as name ,departments.name as department_name").Joins("join employees on departments.head_department = employees.id").Find(&rows)
	result, _ := json.Marshal(rows)
	fmt.Println(string(result))

}
