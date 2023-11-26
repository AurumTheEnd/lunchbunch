package server

import (
	"encoding/gob"
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
	gob.Register(handlers.AppContext{})

	var staticFs = http.FileServer(http.Dir("web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFs))

	mux.HandleFunc(
		constants.IndexPath,
		Chain(
			ReusableHandler(appContext, true, appContext.GetIndex, appContext.PostIndex),
			DisallowSubtreeWrapper(constants.IndexPath)))
	mux.HandleFunc(constants.RegisterFormPath, ReusableHandler(appContext, true, appContext.GetRegisterForm, appContext.PostRegisterForm))
	mux.HandleFunc(constants.LoginFormPath, ReusableHandler(appContext, true, appContext.GetLoginForm, appContext.PostLoginForm))
	mux.HandleFunc(constants.LogoutPath, ReusableHandler(appContext, false, appContext.GetLogout, nil))

	var server = http.Server{
		Addr:         fmt.Sprintf("%s:%s", myEnv["SERVER_HOST"], myEnv["SERVER_PORT"]),
		Handler:      mux,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}

	log.Printf("Running server at %s", server.Addr)
	err = server.ListenAndServe()

	return
}
