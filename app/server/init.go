package server

import(
  "log"
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/alvesmarcos/investapi/app/bundles/userbundle"
  "github.com/alvesmarcos/investapi/app/bundles/reportbundle"
  "github.com/alvesmarcos/investapi/app/bundles/indicatorbundle"
  "github.com/alvesmarcos/investapi/app/core"
)

type Server struct {
  cfg core.Config
}

func NewServer() *Server {
  return &Server{ core.Config{} }
}

func (s *Server) Start() error {
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
  log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(api)))

  return nil
}

func initHandlers(db *gorm.DB) []core.Handler {
  return []core.Handler{ userbundle.NewUserHandler(db), reportbundle.NewReportHandler(db), indicatorbundle.NewIndicatorHandler(db) }
}

func migrateModels(db *gorm.DB) {
  db.AutoMigrate(&userbundle.User{})
  db.AutoMigrate(&reportbundle.Report{})
  db.AutoMigrate(&indicatorbundle.Indicator{}, &indicatorbundle.Sample{})
}
