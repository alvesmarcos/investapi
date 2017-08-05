package userbundle

import(
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/alvesmarcos/investapi/app/core"
)

type UserHandler struct {
  routes  []core.Route
}

func NewUserHandler(db *gorm.DB) *UserHandler {
  ump := NewUserMapperPSQL(db)
  uc := NewUserController(*ump)

  r := []core.Route {
    core.Route {
      Method:   http.MethodGet,
      Path:     "/users",
      Handler:  uc.Index,
    },
    core.Route {
      Method:   http.MethodGet,
      Path:     "/users/{username}",
      Handler:  uc.Get,
    },
    core.Route {
      Method:   http.MethodPost,
      Path:     "/users",
      Handler:  uc.Create,
    },
    core.Route {
      Method:   http.MethodDelete,
      Path:     "/users/{id}",
      Handler:  uc.Delete,
    },
    core.Route {
      Method:   http.MethodPut,
      Path:     "/users/{id}",
      Handler:  uc.Update,
    },
  }

  return &UserHandler { routes: r }
}

func (u *UserHandler) GetRoutes() []core.Route {
  return u.routes
}
