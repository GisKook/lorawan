package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func handler(name string)func (http.ResponseWriter, *http.Request) {
    dummyHandler := func (w http.ResponseWriter, r *http.Request) {
        route := mux.CurrentRoute(r)
        url := r.URL.String()
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%s [%s] %s\n", route.GetName(), name, url)
        vars := mux.Vars(r)
        for k, v := range vars {
            fmt.Fprintf(w, "  [%s]: \"%s\"\n", k, v)
        }
    }
    return dummyHandler
}

func subsetup1(router *mux.Router, prefix string) {
    sub := router.PathPrefix(prefix).Subrouter()
    // there is no way to make 1.1 work - mux requires all patterns to start with '/'
    // sub.HandleFunc("", handler("router.subrouter")).Name("1.1")

    sub.HandleFunc("/", handler("router.subrouter/")).Name("1.2")
    sub.HandleFunc("/foo", handler("router.subrouter/foo")).Name("1.3")
    sub.HandleFunc("/foo/{id}", handler("router.subrouter/foo/id")).Name("1.4")

    // 1.1 workaround
    router.HandleFunc(prefix, handler("router.handlefunc")).Name("1.1")
}

func subsetup2(router *mux.Router, prefix string) {
    sub := mux.NewRouter()
    sub.HandleFunc(prefix, handler("router.handle")).Name("2.1")
    sub.HandleFunc(prefix + "/", handler("router.handle/")).Name("2.2")
    sub.HandleFunc(prefix + "/foo", handler("router.handle/foo")).Name("2.3")
    sub.HandleFunc(prefix + "/foo/{id}", handler("router.handle/foo/id")).Name("2.4")
    // required for 2.1
    router.Handle(prefix, sub)
    // required for 2.2 and 2.3
    router.Handle(prefix + "/{path:.*}", sub)
}

func subsetup3(router *mux.Router, prefix string) {
    sub := mux.NewRouter()
    sub.HandleFunc(prefix, handler("http.handle")).Name("3.1")
    sub.HandleFunc(prefix + "/", handler("http.handle/")).Name("3.2")
    sub.HandleFunc(prefix + "/foo", handler("http.handle/foo")).Name("3.3")
    sub.HandleFunc(prefix + "/foo/{id}", handler("http.handle/foo/id")).Name("3.4")
    // required for 3.1 (otherwise, 3.1 does not get called and net/http redirects you to 3.2)
    http.Handle(prefix, sub)
    // required for 3.2 and 3.3
    http.Handle(prefix + "/", sub)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", handler("root/")).Name("0.1")

    subsetup1(router, "/s1")
    subsetup2(router, "/s2")
    subsetup3(router, "/s3")

    http.Handle("/", router)

    err := http.ListenAndServe(":1337", nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
