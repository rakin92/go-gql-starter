package resolvers

import (
	"testing"

	"github.com/rakin92/go-gql-starter/db"
	"github.com/rakin92/go-gql-starter/model"
	"github.com/rakin92/go-gql-starter/utils"
)

func TestSignIn(t *testing.T) {
	utils.ConfigureEnv()
	db, err := db.ConnectDB()

	defer db.DB.Close()

	if err != nil {
		t.Errorf(err.Error())
	}
	user := model.User{}
	db.DB.Where("email = ?", "notexisting@test.com").First(&user)

	t.Log(user.ID)
}
