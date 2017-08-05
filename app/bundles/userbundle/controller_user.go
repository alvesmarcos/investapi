package userbundle

import(
  "errors"
  "log"
  "io/ioutil"
  "net/url"
  "net/http"
	"strconv"
	"github.com/gorilla/mux"
  "github.com/alvesmarcos/investapi/app/core"
)

var (
  ErrorNotFound         = errors.New("Not Found")
  ErrorInvalidUser      = errors.New("Invalid username or password")
  ErrorInvalidId        = errors.New("Invalid id")
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
  users, err := c.ump.FindAll()

  if c.HandleError(err, w) {
    c.SendJSON(w, ErrorNotFound, http.StatusNotFound)
    return
  }
  c.SendJSON(w, &users, http.StatusOK)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w, nil, http.StatusInternalServerError)
    return
  }

  values, err := url.ParseQuery(string(body))

  user := User { Username: values.Get("username"), Password: values.Get("password") }

  if !user.Validate() {
    c.SendJSON(w, ErrorInvalidUser, http.StatusBadRequest)
    return
  }

  if err = c.ump.Insert(&user); err != nil {
    c.SendJSON(w, nil, http.StatusInternalServerError)
    return
  }
  c.SendJSON(w, user, http.StatusOK)
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  query := r.URL.Query()

  username, password := vars["username"], query.Get("password")

  user := User { Username: username, Password: password }

  if !user.Validate() {
    c.SendJSON(w, ErrorInvalidUser, http.StatusBadRequest)
    return
  }

  users, err := c.ump.FindUser(&user)
  if err != nil {
    c.SendJSON(w, ErrorInvalidUser, http.StatusInternalServerError)
  }

  c.SendJSON(w, &users, http.StatusOK)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  id, err := strconv.Atoi(vars["id"])

  if c.HandleError(err, w) {
    c.SendJSON(w, ErrorInvalidId, http.StatusBadRequest)
    return
  }

  if c.HandleError(c.ump.Delete(id), w) {
    c.SendJSON(w, ErrorNotFound, http.StatusNotFound)
    return
  }
  c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    c.SendJSON(w, nil, http.StatusInternalServerError)
    return
  }

  values, err := url.ParseQuery(string(body))

  id, err := strconv.Atoi(vars["id"])

  if c.HandleError(err, w) {
    c.SendJSON(w, ErrorInvalidId, http.StatusBadRequest)
    return
  }
  user := User { Id: id, Username: values.Get("username"), Password: values.Get("password") }
    log.Println(user)
  if c.HandleError(c.ump.Update(&user), w) {
    c.SendJSON(w, ErrorNotFound, http.StatusNotFound)
    return
  }
  c.SendJSON(w, nil, http.StatusOK)
}
