package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/muhammad-zaka/bookings/pkg/config"
	"github.com/muhammad-zaka/bookings/pkg/handlers"
	"github.com/muhammad-zaka/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

// main is the main function
func main() {

	// change this to tro trye when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//fmt.Println("hello, World")
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
