package server

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/scraping"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func StartServer(db *gorm.DB, store *sessions.CookieStore) (err error) {
	var myEnv map[string]string
	myEnv, err = godotenv.Read()
	var mux = http.NewServeMux()
	var appContext = &handlers.AppContext{
		Db:          db,
		C:           scraping.ConfigScraper(),
		CookieStore: store,
	}

	var staticFs = http.FileServer(http.Dir("web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFs))

	mux.HandleFunc(constants.IndexPath, appContext.IndexHandler())
	mux.HandleFunc(constants.RegisterFormPath, appContext.RegisterHandler)
	mux.HandleFunc(constants.LoginFormPath, appContext.LoginHandler)
	mux.HandleFunc(constants.LogoutPath, appContext.LogoutHandler)

	var server = http.Server{
		Addr:         fmt.Sprintf("%s:%s", myEnv["SERVER_HOST"], myEnv["SERVER_PORT"]),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Running server at %s", server.Addr)
	err = server.ListenAndServe()

	return
}
