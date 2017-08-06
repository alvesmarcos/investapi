package reportbundle

import(
  "io/ioutil"
  "net/url"
  "net/http"
	"strconv"
	"github.com/gorilla/mux"
  "github.com/alvesmarcos/investapi/app/core"
)

type ReportController struct {
  core.Controller
  rmp ReportMapperPSQL
}

func NewReportMapperPSQL(rmp ReportMapperPSQL) *ReportMapperPSQL {
  return &ReportController {
    Controller: core.Controller{},
    rmp:        rmp,
  }
}

func (c *ReportController) Index(w http.ResponseWriter, r *http.Request) {
  reports, err = c.rmp.FindAll()

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  c.SendJSON(w, &reports, http.StatusOK)
}

func (c *ReportController) GetById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.BadRequest), http.BadRequest)
    return
  }
  report, err := c.rmp.FindReportById(id)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.NotFound), http.NotFound)
    return
  }
  c.SendJSON(w, &report, http.StatusOK)
}

func (c *ReportController) Create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w,  http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  values, err := url.ParseQuery(string(body))

  report = Report{ Title: values.Get("title"), Body: values.Get("body"), Images: values["images"]) }

  if !report.Validate() {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.rmp.Insert(&report); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  c.SendJSON(w, &report, http.StatusOK)
}

func (c *ReportController) Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.BadRequest), http.BadRequest)
    return
  }

  if err = c.rmp.Delete(id); err != nil {
    c.SendJSON(w, http.StatusText(http.NotFound), http.NotFound)
    return
  }
  c.SendJSON(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func (c *ReportController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.BadRequest), http.BadRequest)
    return
  }
  report, err := c.rmp.FindReportById(id)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.NotFound), http.NotFound)
    return
  }
  values, err := url.ParseQuery(string(body))

  report.CompareAndSwap(Report {Title: values.Get("title"), Body: values.Get("Body"), Images: values["images"]})

  if err = c.rmp.Update(&report) ; err != nil {
    c.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  c.SendJSON(w, &report, http.StatusOK)
}
