package main

import (
	"fmt"
	config2 "github.com/skripov-ds-ai/social_network/cmd/config"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/skripov-ds-ai/social_network/generated"
	"github.com/skripov-ds-ai/social_network/internal/app"
	"github.com/skripov-ds-ai/social_network/internal/app/login"
	user3 "github.com/skripov-ds-ai/social_network/internal/app/user"
	"github.com/skripov-ds-ai/social_network/internal/database/user"
	user2 "github.com/skripov-ds-ai/social_network/internal/services/user"
)

func main() {
	config, err := config2.LoadConfig(".")

	db, err := sqlx.Open("pgx", config.URI())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = db.Close()
	}()

	userRepo := user.NewRepository(db)

	userService := user2.NewUserService(userRepo)

	uh := user3.NewUserHandler(userService)
	lh := login.NewLoginHandler(userService)

	h := app.NewHandler(lh, uh)

	srv, err := generated.NewServer(h, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), srv); err != nil {
		log.Fatal(err)
	}
}
