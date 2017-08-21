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
  ump UserMapperPSQL
}

func NewUserController(ump UserMapperPSQL) *UserController {
  return &UserController { ump: ump }
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
  users, err := c.ump.FindAll()

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  core.SendJSON(w, &users, http.StatusOK)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  username, password := query.Get("username"), query.Get("password")

  if len(username) == 0 || len(password) == 0 {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  user := User { Username: username, Password: password }

  userx, err := c.ump.FindUser(&user)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  core.SendJSON(w, &userx, http.StatusOK)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    core.SendJSON(w,  http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  values, err := url.ParseQuery(string(body))

  user := User { Name: values.Get("name"), Username: values.Get("username"), Password: values.Get("password") }

  if !user.Validate() {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.ump.Insert(&user); err != nil {
    core.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  core.SendJSON(w, user, http.StatusOK)
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  user, err := c.ump.FindUserById(id)
  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }

  core.SendJSON(w, &user, http.StatusOK)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if err = c.ump.Delete(id); err != nil {
    core.SendJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  core.SendJSON(w, http.StatusText(http.StatusOK), http.StatusOK)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  id, err := strconv.Atoi(vars["id"])

  if err != nil {
    core.SendJSON(w,  http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
  user, err := c.ump.FindUserById(id)

  if err != nil {
    core.SendJSON(w,  http.StatusText(http.StatusNotFound), http.StatusNotFound)
    return
  }
  values, err := url.ParseQuery(string(body))

  if err != nil {
    core.SendJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  user.CompareAndSwap(User { Name: values.Get("name"), Username: values.Get("username"), Password: values.Get("password") })

  if err = c.ump.Update(&user) ; err != nil {
    core.SendJSON(w, http.StatusText(http.StatusConflict), http.StatusConflict)
    return
  }
  core.SendJSON(w, user, http.StatusOK)
}
