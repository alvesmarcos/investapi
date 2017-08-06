package userbundle

import(
  "io/ioutil"
  "net/url"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
  "github.com/alvesmarcos/investapi/app/core"
)

type UserController struct {
  core.Controller
  ump UserMapperPSQL
}

func NewUserController(ump UserMapperPSQL) *UserController {
  return &UserController {
    Controller: core.Controller{},
    ump:        ump,
  }
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  username, password := query.Get("username"), query.Get("password")

  user := User { Username: username, Password: password }

  if len(username) == 0 || len(password) == 0 {
    users, err := c.ump.FindAll()

    if err != nil {
      c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      return
    }
    c.SendJSON(w, &users, http.StatusOK)
  } else {
    userx, err := c.ump.FindUser(&user)
    if err != nil {
      c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      return
    }
    c.SendJSON(w, &userx, http.StatusOK)
  }
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w,  http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  values, err := url.ParseQuery(string(body))

  user := User { Username: values.Get("username"), Password: values.Get("password") }

  if !user.Validate() {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.ump.Insert(&user); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  c.SendJSON(w, user, http.StatusOK)
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  user, err := c.ump.FindUserById(id)
  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }

  c.SendJSON(w, &user, http.StatusOK)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.ump.Delete(id); err != nil {
    c.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  c.SendJSON(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    c.SendJSON(w,  http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  user, err := c.ump.FindUserById(id)

  if err != nil {
    c.SendJSON(w,  http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  values, err := url.ParseQuery(string(body))

  if err != nil {
    c.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  user.CompareAndSwap(User { Username: values.Get("username"), Password: values.Get("password") })

  if err = c.ump.Update(&user) ; err != nil {
    c.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  c.SendJSON(w, user, http.StatusOK)
}
