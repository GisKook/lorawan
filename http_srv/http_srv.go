package http_srv

import (
	"github.com/giskook/lorawan/conf"
	"github.com/giskook/lorawan/db"
	"github.com/giskook/lorawan/redis_socket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpSrv struct {
	conf   *conf.Conf
	router *mux.Router
	db     *db.DBSocket
	redis  *redis_socket.RedisSocket
}

func NewHttpSrv(conf *conf.Conf) *HttpSrv {
	db, err := db.NewDBSocket(conf.DB)
	if err != nil {
		log.Println(err.Error())

		return nil
	}

	redis, e := redis_socket.NewRedisSocket(conf.Redis)
	if e != nil {
		log.Println(e.Error())
		db.Close()

		return nil
	}
	redis.InitPool()

	return &HttpSrv{
		conf:   conf,
		router: mux.NewRouter(),
		db:     db,
		redis:  redis,
	}
}

func (h *HttpSrv) Start() {
	s := h.router.PathPrefix("/web").Subrouter()
	h.init_web(s)

	h.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(h.conf.Http.Path))))

	if err := http.ListenAndServe(h.conf.Http.Addr, h.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (h *HttpSrv) init_web(r *mux.Router) {
	r.HandleFunc("/user/login", h.handler_web_user_login)
	r.HandleFunc("/user/logout", h.handler_web_user_logout)
	//r.HandleFunc("/user/main", h.handler_web_user_main)
}
