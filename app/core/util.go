package core

import (
  "fmt"
  "strings"
  "strconv"
  "database/sql/driver"
  "encoding/json"
  "io"
  "log"
  "net/http"
)

func SendJSON(w http.ResponseWriter, v interface{}, code int) {
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

// http://blog.tamizhvendan.in/blog/2017/07/03/leveraging-interfaces-in-golang-part-2/

type StringSlice []string

func (stringSlice StringSlice) Value() (driver.Value, error) {
  var quotedStrings []string

  for _, str := range stringSlice {
    quotedStrings = append(quotedStrings, strconv.Quote(str))
  }
  value := fmt.Sprintf("{ %s }", strings.Join(quotedStrings, ","))

  return value, nil
}

func (stringSlice *StringSlice) Scan(src interface{}) error {
  val, ok := src.([]byte)

  if !ok {
    return fmt.Errorf("unable to scan")
  }
  value := strings.TrimPrefix(string(val), "{")
  value = strings.TrimSuffix(value, "}")

  *stringSlice = strings.Split(value, ",")

  return nil
}
