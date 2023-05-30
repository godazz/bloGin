package routing

import (
	"fmt"
	"log"

	"github.com/godazz/bloGin/pkg/config"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in Routing")
	}
}
