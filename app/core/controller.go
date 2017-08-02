package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Controller struct {}

func (c *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {
  w.Header().Add("Content-Type", "application/json")

  // Marshal returns the JSON encoding of v
  b, err := json.Marshal(v)

  if err != nil {
    log.Print(fmt.Sprint("Erro while encoding JSON: %v", err))
    w.WriteHeader(http.StatusInternalServerError)
    io.WriteString(w, `{"erro": "Internal server error"}`)
  } else {
    w.WriteHeader(code);
    io.WriteString(w, string(b))
  }
}

func (c *Controller) HandleError(err error, w http.ResponseWriter) bool {
  if err == nil {
    return false
  }
  msg := map[string]string{
    "message": "An error occurred"
  }
  c.SendJSON(w, &msg, http.StatusInternalServerError)

  return true
}
