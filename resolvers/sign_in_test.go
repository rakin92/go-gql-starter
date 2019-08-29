package resolvers

import (
	"testing"

	"github.com/rakin92/go-gql-starter/db"
	"github.com/rakin92/go-gql-starter/model"
)

func TestSignIn(t *testing.T) {
	db, err := db.GetMockDB()

	defer db.DB.Close()

	if err != nil {
		t.Errorf(err.Error())
	}
	user := model.User{}
	db.DB.Where("email = ?", "notexisting@test.com").First(&user)

	t.Log(user.ID)
}
