package main

import(
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/alvesmarcos/investapi/app/bundles/userbundle"
	"github.com/alvesmarcos/investapi/app/core"
)

func main() {
	initDB()
	startServer()
}

func initDB() error {
	cfg := &core.Config{}
	cfg.Fetch()

	db, err := gorm.Open(cfg.DBTYPE, cfg.DBConnection)
	defer db.Close()

	if err != nil {
		return err
	}
	db.AutoMigrate(&userbundle.User{})
	db.Create(userbundle.NewUser("Marcos", "admin123"))
	db.Create(userbundle.NewUser("SDA_API", "09212"))

	return nil
}

func startServer() error {
	cfg := &core.Config{}
	cfg.Fetch()

	db, err := gorm.Open(cfg.DBTYPE, cfg.DBConnection)
	defer db.Close()

	if err != nil {
		return err
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	userHandler := userbundle.NewUserHandler(db)
	routes := userHandler.GetRoutes()

	s.HandleFunc(routes[0].Path, routes[0].Handler).Methods(routes[0].Method)

	// Routes handling
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}
