package thor

import "net/http"

type Route struct {

}

type Router struct {

}

func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

}

func NewRouter() *Router {
	return &Router{}
}
