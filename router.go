package httprouter

import (
	"net/http"
	"strings"
)

// Router Create a radix tree based HTTP router
type Router struct {
	tree *RadixTree
}

// NewRouter Initialize the router
func NewRouter() *Router {
	return &Router{
		tree: NewRadixTree(),
	}
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handle)
func (r *Router) GET(path string, handle http.HandlerFunc) {
	r.Handle(http.MethodGet, path, handle)
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (r *Router) POST(path string, handle http.HandlerFunc) {
	r.Handle(http.MethodPost, path, handle)
}

func (r *Router) Handle(method, path string, handle http.HandlerFunc) {
	if method == "" {
		panic("method must not be empty")
	}
	
	if len(path) < 1 || path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}
	
	if handle == nil {
		panic("handle must not be nil")
	}
	
	if r.tree == nil {
		r.tree = NewRadixTree()
	}
	
	r.AddRoute(method, path, handle)
}

// AddRoute Add a route to the router
func (r *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	key := r.makeSearchKey(method, path)
	_, exist := r.tree.Search(key)
	if exist {
		panic("http: multiple registrations for " + path)
	}
	r.tree.Insert(key, handler)
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := r.makeSearchKey(req.Method, req.URL.Path)
	handlerInterface, exist := r.tree.Search(key)
	if handlerInterface != nil && exist {
		handler, ok := handlerInterface.(http.HandlerFunc)
		if ok {
			handler.ServeHTTP(w, req)
		}
	} else {
		w.Write([]byte(key + " not found "))
		http.NotFound(w, req)
	}
}

// makeSearchKey get a common key
func (r *Router) makeSearchKey(method string, path string) string {
	key := strings.ToUpper(method) + ":" + strings.ToUpper(path)
	return key
}
