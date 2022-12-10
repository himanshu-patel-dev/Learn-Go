package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 172.18.0.2 localhost
var pgconnect = "host=%s user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
var GORMDB *gorm.DB

type JSONB map[string]bool

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte")
	}
	return json.Unmarshal(b, &j)
}

type Request struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address JSONB  `json:"address"`
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	data := Request{}
	json.NewDecoder(r.Body).Decode(&data)
	err := GORMDB.Table("request").Create(&Request{
		Name:    data.Name,
		Age:     data.Age,
		Address: data.Address,
	}).Error
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Failed")
		return
	}
	fmt.Fprint(w, "Inserted")
}

func setupDB(DBhost string) {
	var err error
	pgconnect = fmt.Sprintf(pgconnect, DBhost)
	GORMDB, err = gorm.Open(postgres.Open(pgconnect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func main() {
	DBhost := flag.String("DBhost", "172.18.0.2", "Host Address of DB")
	flag.Parse()

	setupDB(*DBhost)
	http.HandleFunc("/", HandlerHello)
	http.ListenAndServe(":8000", nil)
}
