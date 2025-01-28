package connections

import (
	"database/sql"
	"fmt"
	userservice "tweets/internal/client/user"
	"tweets/internal/config"
	database "tweets/internal/database/sql"
	interface17 "tweets/internal/interface"
	"tweets/internal/service"
	adjust "tweets/internal/service/Adjust"

	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interface17.TweetsService {
	c := config.Configuration()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return &database.Database{Db: db}
}

func NewAdjust() *adjust.Adjust {
	d := NewDatabase()
	u := userservice.UserClinet()
	return &adjust.Adjust{D: d, User: u}
}

func NewService() *service.Service {
	a := NewAdjust()
	return &service.Service{A: a}
}
