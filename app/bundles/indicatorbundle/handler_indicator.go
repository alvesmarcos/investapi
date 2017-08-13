package indicatorbundle

import(
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/alvesmarcos/investapi/app/core"
)

type IndicatorHandler struct {
  routes []core.Route
}

func NewIndicatorHandler(db *gorm.DB) *IndicatorHandler {
  imp := NewIndicatorMapperPSQL(db)
  ic := NewIndicatorController(*imp)

  r := []core.Route{
    core.Route {
      Method:   http.MethodGet,
      Path:     "/indicators",
      Handler:  ic.Index,
    },
    core.Route {
      Method:   http.MethodGet,
      Path:     "/indicators/{id}",
      Handler:  ic.GetById,
    },
    core.Route {
      Method:   http.MethodPost,
      Path:     "/indicators",
      Handler:  ic.Create,
    },
    core.Route {
      Method:   http.MethodDelete,
      Path:     "/indicators/{id}",
      Handler:  ic.Delete,
    },
    core.Route {
      Method:   http.MethodPut,
      Path:     "/indicators/{id}",
      Handler:  ic.Update,
    },
    core.Route {
      Method:   http.MethodPut,
      Path:     "/indicators/{id}/samples",
      Handler:  ic.UpdateSamples,
    },
  }
  return &IndicatorHandler{ routes: r }
}

func (i *IndicatorHandler) GetRoutes() []core.Route {
  return i.routes
}
