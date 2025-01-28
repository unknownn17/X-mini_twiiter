package connections

import (
	"database/sql"
	"fmt"
	notification17 "like/internal/client/notification"
	likeretweet17 "like/internal/client/tweet"
	"like/internal/config"
	database "like/internal/database/sql"
	interface17 "like/internal/interface"
	"like/internal/service"
	adjust "like/internal/service/Adjust"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interface17.LRC {
	c := config.Configuration()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	n := notification17.UserClinet()
	return &database.Database{Db: db, N: n}
}

func NewAdjust() *adjust.Adjust {
	db := NewDatabase()
	a := likeretweet17.UserClinet()
	return &adjust.Adjust{D: db, A: a}
}

func NewService() *service.Service {
	a := NewAdjust()
	return &service.Service{A: a}
}
