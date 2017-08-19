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
  rmp ReportMapperPSQL
}

func NewReportController(rmp ReportMapperPSQL) *ReportController {
  return &ReportController { rmp: rmp }
}

func (c *ReportController) Index(w http.ResponseWriter, r *http.Request) {
  reports, err := c.rmp.FindAll()

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  core.SendJSON(w, &reports, http.StatusOK)
}

func (c *ReportController) GetById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  report, err := c.rmp.FindReportById(id)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  core.SendJSON(w, &report, http.StatusOK)
}

func (c *ReportController) Create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    core.SendJSON(w,  http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  values, err := url.ParseQuery(string(body))

  report := Report{ Title: values.Get("title"), Body: values.Get("body"), Images: values["images"] }

  if !report.Validate() {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.rmp.Insert(&report); err != nil {
    core.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  core.SendJSON(w, &report, http.StatusOK)
}

func (c *ReportController) Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.rmp.Delete(id); err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  core.SendJSON(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func (c *ReportController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  report, err := c.rmp.FindReportById(id)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  values, err := url.ParseQuery(string(body))

  report.CompareAndSwap(Report {Title: values.Get("title"), Body: values.Get("Body") })

  if err = c.rmp.Update(&report) ; err != nil {
    core.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  core.SendJSON(w, &report, http.StatusOK)
}

func (c *ReportController) UpdateImagesById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  id, err := strconv.Atoi(vars["id"])

  report, err := c.rmp.FindReportById(id)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  values, err := url.ParseQuery(string(body))

  if len(values.Get("path")) == 0 {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if len(values.Get("index")) == 0 {
    report.PushImage(values.Get("path"))
  } else {
    index, err := strconv.Atoi(values.Get("index"))

    if err != nil {
      core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
      return
    }
    report.UpdateImages(index, values.Get("path"))
  }

  if err = c.rmp.Update(&report) ; err != nil {
    core.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  core.SendJSON(w, &report, http.StatusOK)
}
