package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/inegmetov/bookings/internal/config"
	"github.com/inegmetov/bookings/internal/handlers"
	"github.com/inegmetov/bookings/internal/helpers"
	"github.com/inegmetov/bookings/internal/models"
	"github.com/inegmetov/bookings/internal/render"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	return nil
}
