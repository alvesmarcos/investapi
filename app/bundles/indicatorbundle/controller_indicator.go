package indicatorbundle

import(
  "io/ioutil"
  "net/url"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
  "github.com/alvesmarcos/investapi/app/core"
)

type IndicatorController struct {
  core.Controller
  imp IndicatorMapperPSQL
}

func NewIndicatorController(imp IndicatorController) *IndicatorController {
  return &IndicatorController{
    Controller: core.Controller{},
    imp:        imp,
  }
}

func (c *IndicatorController) Index(w http.ResponeWriter, r *http.Request) {
  indicators, err := imp.FindAll()

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  c.SendJSON(w, &indicators, http.StatusOK)
}

func (c *IndicatorController) GetById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars()

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  indicator, err := imp.FindIndicatorById(id)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }

  c.SendJSON(w, &indicator, http.StatusOK)
}

func (c *IndicatorController) Create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  values, err := url.ParseQuery(string(body))

  indicator := Indicator{ Name: values.Get("name"), Description: values.Get("description"), Metric: values.Get("metric"), Status: "equal"}

  if err = c.imp.Insert(&indicator); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  c.SendJSON(w, &indicator, http.StatusOK)
}

func (c *IndicatorController) Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars()

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = imp.Delete(id); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  c.SendJSON(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func (c *IndicatorController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.vars()
  body, err := ioutil.ReadAll(r.Body)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  indicator, err := c.imp.FindIndicatorById(id)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }

  values, err := url.ParseQuery(string(body))

  indicator := Indicator{ Name: values.Get("name"), Description: values.Get("description"), Metric: values.Get("metric"), Status: "equal"}

  if err = c.imp.Update(&indicator); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  c.SendJSON(w, &indicator, http.StatusOK)
}

func (c *IndicatorController) UpdateSamples(w http.ResponseWriter, r *http.Request) {
  vars := mux.vars()
  body, err := ioutil.ReadAll(r.Body)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  indicator, err := c.imp.FindIndicatorById(id)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }

  values, err := url.ParseQuery(string(body))

  indicator := Indicator{ Name: values.Get("name"), Description: values.Get("description"), Metric: values.Get("metric"), Status: "equal"}

  if err = c.imp.Update(&indicator); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  c.SendJSON(w, &indicator, http.StatusOK)
}
