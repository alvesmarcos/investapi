package reportbundle

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/alvesmarcos/investapi/app/core"
)

type ReportHandler struct {
  routes  []core.Route
}

func NewReportHandler(db *gorm.DB) *ReportHandler {
  rmp := NewReportMapperPSQL(db)
  rc := NewReportController(*rmp)

  r := []core.Route{
    core.Route {
      Method:   http.MethodGet,
      Path:     "/reports",
      Handler:  rc.Index,
    },
    core.Route {
      Method:   http.MethodGet,
      Path:     "/reports/{id}",
      Handler:  rc.GetById,
    },
    core.Route {
      Method:   http.MethodPost,
      Path:     "/reports",
      Handler:  rc.Create,
    },
    core.Route {
      Method:   http.MethodDelete,
      Path:     "/reports/{id}",
      Handler:  rc.Delete,
    },
    core.Route {
      Method:   http.MethodPut,
      Path:     "/reports/{id}",
      Handler:  rc.Update,
    },
    core.Route {
      Method:   http.MethodPut,
      Path:     "/reports/{id}/images",
      Handler:  rc.UpdateImagesById,
    },
  }
  return &ReportHandler{ routes: r }
}

func (r *ReportHandler) GetRoutes() []core.Route {
  return r.routes
}
