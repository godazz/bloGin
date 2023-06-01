package seeder

import (
	"fmt"
	"log"

	articleModel "github.com/godazz/bloGin/internal/modules/article/models"
	userModel "github.com/godazz/bloGin/internal/modules/user/models"
	"github.com/godazz/bloGin/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := db.Connecion()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("godazz"), 12)
	if err != nil {
		log.Fatal("unable to has the password")
		return
	}
	user := userModel.User{Name: "godazz", Email: "godazz@gmail.com", Password: string(hashedPassword)}
	db.Create(&user)

	log.Printf("Created User Sucessfuly %s\n", user.Email)

	for i := 0; i < 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Random title %d", i+1), Content: fmt.Sprintf("Random content %d", i+1), UserID: 1}
		db.Create(&article)
		log.Printf("Created Article Sucessfuly with title: %s\n", article.Title)
	}

	log.Println("Seeder done...")
}
