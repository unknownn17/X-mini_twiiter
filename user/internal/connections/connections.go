package connections

import (
	"context"
	"database/sql"
	"fmt"

	"log"

	"cmd/main.go/internal/auth"
	"cmd/main.go/internal/config"
	redis17 "cmd/main.go/internal/database/redis"
	database "cmd/main.go/internal/database/sql"
	interface17 "cmd/main.go/internal/interface"
	"cmd/main.go/internal/service"
	adjust "cmd/main.go/internal/service/Adjust"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func NewDatabase() interface17.User {
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

func Redis() *redis17.Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &redis17.Redis{R: client, Ctx: ctx}
}

func NewAuthorization() interface17.Authorazition {
	db := NewDatabase()
	ctx := context.Background()
	r := Redis()

	return &auth.Auth{Db: db, Ctx: ctx, R: r}
}

func NewAdjust() *adjust.Adjust {
	db := NewDatabase()
	a := NewAuthorization()
	return &adjust.Adjust{Auth: a, D: db}
}

func Newservice() *service.Grpc {
	a := NewAdjust()
	return &service.Grpc{A: a}
}
