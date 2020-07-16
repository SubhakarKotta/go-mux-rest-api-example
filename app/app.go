package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-mux-rest-api-example/app/handler"
	"go-mux-rest-api-example/app/model"
	"go-mux-rest-api-example/config"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		panic("Could not connect database")
	}
	fmt.Println("Connected to database successfully")
	a.Router = mux.NewRouter()
	a.DB = model.DBMigrate(db)
	a.setRouters()
}
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
func (a *App) setRouters() {
	a.Get("/projects", a.handleRequest(handler.GetAllProjects))
	a.Get("/projects/{title}", a.handleRequest(handler.GetProject))
	a.Post("/projects", a.handleRequest(handler.CreateProject))
	a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
	a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
}
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
