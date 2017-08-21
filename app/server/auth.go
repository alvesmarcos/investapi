// package server
//
// import(
//   "io/ioutil"
//   "strings"
//   "net/http"
//   "github.com/auth0/go-jwt-middleware"
//   "github.com/codegangsta/negroni"
//   "github.com/dgrijalva/jwt-go"
// )
//
// type Auth struct {
//   VerifyKey []byte
//   SignKey   []byte
// }
//
// type Credential struct {
//   Username  string  `json:"username"`
//   Password  string  `json:"password"`
// }
//
// type Key struct {
//   Token string  `json:"token"`
// }
//
// func NewAuth(privKeyPath, pubKeyPath string) *Auth {
//   return &Auth{}
// }
//
// func (a *Auth) GetToken(w http.ResponseWriter, r *http.Request) {
//   var credential Credential
//
//   err := json.NewDecoder(r.Body).Decode(&credential)
//
//   if err != nil {
//     return
//   }
//
//   if strings.ToLower(credential.Username) != "" || strings.ToLower(credential.Password) != "" {
//     return
//   }
//
//   token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
//     "role": "admin",
//     "exp": time.Now().Add(time.Minute * 60).Unix(),
//   })
//
//   tokeStr, err := token.SignedString()
//
//   if err != nil {
//     return
//   }
//
// }
//
// func (a *Auth) ValidateMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
//   token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
//     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//     }
//     return VerifyKey, nil
//   })
//
//   if err != nil {
//     return
//   }
//
//   if !token.Valid {
//     return
//   }
//   next(w, r)
// }
