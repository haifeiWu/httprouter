package httprouter

import (
	"net/http"
	"testing"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func TestRouterHandle(t *testing.T) {
	router := NewRouter()
	
	routed := false
	router.Handle(http.MethodGet, "/user/gopher", func(writer http.ResponseWriter, request *http.Request) {
		routed = true
		query := request.URL.Query()
		name := query.Get("name")
		if name != "haifeisi" {
			t.Fatalf("wrong wildcard values: want %v, got %v", "haifeisi", name)
		}
	})
	
	w := new(mockResponseWriter)
	
	req, _ := http.NewRequest(http.MethodGet, "/user/gopher?name=haifeisi", nil)
	router.ServeHTTP(w, req)
	
	if !routed {
		t.Fatal("routing failed")
	}
}

func TestRouter_GET(t *testing.T) {
	router := NewRouter()
	
	routed := false
	router.GET("/user/gopher", func(writer http.ResponseWriter, request *http.Request) {
		routed = true
		query := request.URL.Query()
		name := query.Get("name")
		if name != "haifeisi" {
			t.Fatalf("wrong wildcard values: want %v, got %v", "haifeisi", name)
		}
	})
	
	w := new(mockResponseWriter)
	
	req, _ := http.NewRequest(http.MethodGet, "/user/gopher?name=haifeisi", nil)
	router.ServeHTTP(w, req)
	
	if !routed {
		t.Fatal("routing failed")
	}
}

func TestRouter_POST(t *testing.T) {
	router := NewRouter()
	
	routed := false
	router.POST("/user/gopher", func(writer http.ResponseWriter, request *http.Request) {
		routed = true
		query := request.URL.Query()
		name := query.Get("name")
		if name != "haifeisi" {
			t.Fatalf("wrong wildcard values: want %v, got %v", "haifeisi", name)
		}
	})
	
	w := new(mockResponseWriter)
	
	req, _ := http.NewRequest(http.MethodPost, "/user/gopher?name=haifeisi", nil)
	router.ServeHTTP(w, req)
	
	if !routed {
		t.Fatal("routing failed")
	}
}

func TestHttpServer(t *testing.T) {
	router := NewRouter()
	router.POST("/user/gopher", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		name := query.Get("name")
		writer.Write([]byte(name))
		return
	})
	
	// http.HandleFunc("/a/b/c/", func(writer http.ResponseWriter, request *http.Request) {
	// 	query := request.URL.Query()
	// 	name := query.Get("name")
	// 	if name == "" {
	// 		writer.Write([]byte("query not found"))
	// 		return
	// 	}
	// 	writer.Write([]byte(name))
	// 	return
	// })
	http.ListenAndServe(":8090", router)
}
