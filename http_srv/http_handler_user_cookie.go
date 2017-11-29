package http_srv

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type JwtAuth struct {
	Auth   string `json:"au"`
	Userid string `json:"id"`
	jwt.StandardClaims
}

func (h *HttpSrv) gen_jwt_token(auth string, id string) (string, error) {
	claims := JwtAuth{
		auth,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(h.conf.Http.Expire)).Unix(),
			Issuer:    "zk",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(h.conf.Http.Secret))
}

func (h *HttpSrv) set_cookie(res http.ResponseWriter, req *http.Request, auth string, userid string) error {
	token, err := h.gen_jwt_token(auth, userid)
	if err != nil {
		return err
	}
	exp_cookie_time := time.Now().Add(time.Second * time.Duration(h.conf.Http.Expire))
	cookie := http.Cookie{Name: "Auth", Value: token, Expires: exp_cookie_time, HttpOnly: true}
	http.SetCookie(res, &cookie)

	return nil
}
