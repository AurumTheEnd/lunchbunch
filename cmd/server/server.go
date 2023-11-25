package main

import (
	"encoding/gob"
	"fmt"
	"github.com/gorilla/sessions"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gorm.io/gorm"
	"log"
)

func main() {
	var db *gorm.DB
	var store *sessions.CookieStore

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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

	gob.Register(&session.Data{})
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
