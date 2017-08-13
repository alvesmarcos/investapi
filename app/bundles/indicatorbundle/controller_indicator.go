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

func NewIndicatorController(imp IndicatorMapperPSQL) *IndicatorController {
  return &IndicatorController{
    Controller: core.Controller{},
    imp:        imp,
  }
}

func (c *IndicatorController) Index(w http.ResponseWriter, r *http.Request) {
  indicators, err := c.imp.FindAll()

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  c.SendJSON(w, &indicators, http.StatusOK)
}

func (c *IndicatorController) GetById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

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
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.imp.Delete(id); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  c.SendJSON(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func (c *IndicatorController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
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

  indicator.CompareAndSwap(Indicator{ Name: values.Get("name"), Description: values.Get("description"), Metric: values.Get("metric"), Status: values.Get("status")})

  if err = c.imp.Update(&indicator); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  c.SendJSON(w, &indicator, http.StatusOK)
}

func (c *IndicatorController) UpdateSamples(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
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

  valuex, err := strconv.ParseFloat(values.Get("value"), 64)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  sample := Sample{ Date: values.Get("date"), Value: valuex, ReferIndicator: indicator.Id }

  if len(values.Get("index")) == 0 {
    indicator.PushSample(sample)

    if err = c.imp.UpdatePushSample(&sample) ; err != nil {
      c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
    }
  } else {
    index, err := strconv.Atoi(values.Get("index"))

    if err != nil {
      c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
      return
    }
    indicator.UpdateSamples(index, sample)

    if len(sample.Date) == 0 {
      sample.Date = indicator.Samples[index].Date
    }

    if err = c.imp.UpdateSample(&sample) ; err != nil {
      c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
    }
  }
  c.SendJSON(w, &indicator, http.StatusOK)
}
