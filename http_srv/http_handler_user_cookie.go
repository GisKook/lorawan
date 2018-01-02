package http_srv

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	"fmt"
	"context"
)

type ctx_auth_key int

const( 
	COOKIE_AUTH string = "Auth"
	auth_key ctx_auth_key = 0
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
	cookie := http.Cookie{Name: COOKIE_AUTH, Value: token, Path:"/", Expires: exp_cookie_time, HttpOnly: true}
	http.SetCookie(res, &cookie)

	return nil
}

func (h *HttpSrv) validate(page http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			RecordReq(req)
			cookie, err := req.Cookie(COOKIE_AUTH)
			if err != nil {
				http.NotFound(res, req)
				return
			}
	
			token, err := jwt.ParseWithClaims(cookie.Value, &JwtAuth{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method")
				}
				return []byte(h.conf.Http.Secret), nil
			})
			if err != nil {
				http.NotFound(res, req)
				return
			}
			if claims, ok := token.Claims.(*JwtAuth); ok && token.Valid {
				ctx := context.WithValue(req.Context(), auth_key, *claims)
				page(res, req.WithContext(ctx))
			} else {
				http.NotFound(res, req)
				return
			}
		})
	}