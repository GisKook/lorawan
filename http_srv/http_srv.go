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
	h.init_web_user(s)
	h.init_web_device(s)

	h.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(h.conf.Http.Path))))

	if err := http.ListenAndServe(h.conf.Http.Addr, h.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (h *HttpSrv) init_web_user(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/login", h.handler_web_user_login)
	s.HandleFunc("/logout", h.handler_web_user_logout)
	s.HandleFunc("/main", h.validate(h.handler_web_user_main))
	s.HandleFunc("/add", h.validate(h.handler_web_user_add))
	s.HandleFunc("/search", h.validate(h.handler_web_user_search))
	s.HandleFunc("/del", h.validate(h.handler_web_user_del))
}

func (h *HttpSrv) init_web_device(r *mux.Router) {
	s := r.PathPrefix("/device").Subrouter()
	s.HandleFunc("/add", h.validate(h.handler_web_device_add))
}