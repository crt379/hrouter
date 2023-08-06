package hrouter

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	router := NewHttpRouter()

	router.AddRoute("/", Post(func() {
		fmt.Println("POST /")
	}))
	router.AddRoute("/xxx", Post(func() {
		fmt.Println("POST /xxx")
	}))
	router.AddRoute("/xxx/:id", Post(func() {
		fmt.Println("POST /xxx/:id")
	}))
	router.AddRoute("/xxx/:id/:name", Post(func() {
		fmt.Println("POST /xxx/:id/:name")
	}), Get(func() {
		fmt.Println("GET /xxx/:id/:name")
	}))
	router.AddRoute("/aaa")

	cases := []struct {
		name     string
		path     string
		method   string
		expected error
	}{
		{"test root success", "/", "POST", nil},
		{"test root fail 1", "/", "GET", ErrMethodHandlersNotFount},
		{"test root fail 2", "/", "post", ErrMethodHandlersNotFount},
		{"test method success", "/xxx/xxx", "POST", nil},
		{"test many method success 1", "/xxx/xxx/ddd", "POST", nil},
		{"test many method success 2", "/xxx/xxx/ddd", "GET", nil},
		{"test many get success 1", "/xxx/xxx/ddd", "POST", nil},
		{"test many get success 2", "/xxx/xxx/ddd", "POST", nil},
		{"test not handlers", "/aaa", "GET", ErrMethodHandlersNotFount},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			handlers, err := router.GetRouteHandlers(c.path, c.method)
			if err != c.expected {
				t.Fatalf("%v", err)
			}
			for _, f := range handlers {
				f()
			}
		})
	}
}

func Test2(t *testing.T) {
	router := NewHttpRouter()
	r := router.AddRoute("/").Use(func() {
		fmt.Println("auth")
	})
	{
		r.AddRoute("test", Get(func() {
			fmt.Println("GET /test")
		}))
		r.AddRoute("test/test", Get(func() {
			fmt.Println("GET /test/test")
		}))
	}
	router.AddRoute("/test/test2", Get(func() {
		fmt.Println("GET /test/test2")
	}))

	cases := []struct {
		name     string
		path     string
		method   string
		expected error
	}{
		{"1", "/", "POST", ErrMethodHandlersNotFount},
		{"2", "/test", "GET", nil},
		{"3", "/test/test", "GET", nil},
		{"4", "/test/test2", "GET", nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			handlers, err := router.GetRouteHandlers(c.path, c.method)
			if err != c.expected {
				t.Fatalf("%v", err)
			}
			for _, f := range handlers {
				f()
			}
		})
	}
}

func BenchmarkName(b *testing.B) {
	router := NewHttpRouter()
	routes := []string{
		"/abcd/defg/fghj/lidj/:id/name",
		"/1234/defg/fghj/lidj/:id/name",
		"/1234/7890/fghj/lidj/:id/name",
		"/ccc/ddddd/fff/hhhh/:kkkk/oooo/ppppp/qqqqq/yyyyy/jjjjj/dfklasj/123456",
		"/ccc/ddddd/fff/hhhh/:kkkk/oooo/ppppp/qqqqq/yyyyy/jjjjj/dfklasj/12347",
		"/ccc/ddddd/fff/hhhh/:kkkk/oooo/ppppp/qqqqq/yyyyy/jjjjj/dfklasj/1",
		"/ccc/ddddd/fff/hhhh/:kkkk/oooo/ppppp/qqqqq/yyyyy/jjjjj/dfklasj/12",
		"/ccc/ddddd/fff/hhhh/:kkkk/oooo/ppppp/qqqqq/yyyyy/jjjjj/dfklasj/3",
	}
	c := make(chan string, 10000000)

	for _, r := range routes {
		p := r
		router.AddRoute(p, Post(func() {
			c <- p
		}))
	}

	path := "/ccc/ddddd/fff/hhhh/2222/oooo/ppppp/qqqqq/yyyyy/jjjjj/dfklasj/3"
	for i := 0; i < b.N; i++ {
		handlers, _ := router.GetRouteHandlers(path, "POST")
		for _, f := range handlers {
			f()
		}
	}
}
