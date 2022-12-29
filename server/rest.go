package server

import (
	"fmt"
	"log"

	"github.com/foonsoo/ggur-homepage-code/infra"
)

func GetUsers() {
	db := infra.GetConnector()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	var name string

	rows, err := db.Query("select username from users", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(name)
	}
}
