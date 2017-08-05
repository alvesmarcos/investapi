package server

import(
  "log"
  "net/http"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/alvesmarcos/investapi/app/bundles/userbundle"
	"github.com/alvesmarcos/investapi/app/core"
)

type Service struct {
  cfg core.Config
}

func NewService() *Service {
  return &Service{ core.Config{} }
}

func (s *Service) Start() error {
  s.cfg.Fetch()

	db, err := gorm.Open(s.cfg.DBTYPE, s.cfg.DBConnection)
	defer db.Close()

	if err != nil {
		return err
	}

  migrateModels(db)

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/").Subrouter()

  for _, h := range initHandlers(db) {
    for _, route := range h.GetRoutes() {
      api.HandleFunc(route.Path, route.Handler).Methods(route.Method)
    }
  }

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func initHandlers(db *gorm.DB) []core.Handler {
  return []core.Handler{ userbundle.NewUserHandler(db) }
}

func migrateModels(db *gorm.DB) {
  db.AutoMigrate(&userbundle.User{})
  // db.Create(userbundle.NewUser("Marcos", "admin123"))
  // db.Create(userbundle.NewUser("SDA_API", "09212"))
}