package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
	"lunchbunch/internal/database"
	"lunchbunch/internal/server"
	"lunchbunch/internal/session"
)

func main() {
	var db *gorm.DB = nil
	var store *sessions.CookieStore = nil

	var err = database.CreateDbIfNotExists()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err = database.Setup()
	if err != nil {
		fmt.Println(err)
		return
	}

	store, err = session.CreateSessionStore()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = server.StartServer(db, store); err != nil {
		fmt.Println(err)
		return
	}
}
