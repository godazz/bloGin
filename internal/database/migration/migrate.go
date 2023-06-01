package migration

import (
	"fmt"
	"log"

	articleModels "github.com/godazz/bloGin/internal/modules/article/models"
	userModels "github.com/godazz/bloGin/internal/modules/user/models"
	"github.com/godazz/bloGin/pkg/db"
)

func Migrate() {
	db := db.Connecion()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate db")
		return
	}

	fmt.Println("Migration is completed sucessfully")

}
