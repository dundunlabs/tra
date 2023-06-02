package tra

import "net/http"

type Group struct {
	router *Router
	path   string
}

func (g *Group) Group(path string) *Group {
	return &Group{
		path: path,
	}
}

func (g *Group) WithGroup(path string, fn func(g *Group)) {
	fn(g.Group(path))
}

func (g *Group) handle(method string, path string, handler Handler) {
	route := route{
		method:  method,
		path:    path,
		handler: handler,
	}
	g.router.root.addRoute(route)
}

func (g *Group) GET(path string, handler Handler) {
	g.handle(http.MethodGet, path, handler)
}

func (g *Group) POST(path string, handler Handler) {
	g.handle(http.MethodPost, path, handler)
}

func (g *Group) PUT(path string, handler Handler) {
	g.handle(http.MethodPut, path, handler)
}

func (g *Group) PATCH(path string, handler Handler) {
	g.handle(http.MethodPatch, path, handler)
}

func (g *Group) DELETE(path string, handler Handler) {
	g.handle(http.MethodDelete, path, handler)
}