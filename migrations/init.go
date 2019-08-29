package main

import (
	"github.com/rakin92/go-gql-starter/db"
	"github.com/rakin92/go-gql-starter/model"
	"github.com/rakin92/go-gql-starter/utils"
)

func main() {
	utils.ConfigureEnv()

	d, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer d.Close()

	d.DropTableIfExists(&model.User{})
	d.CreateTable(&model.User{})
}
