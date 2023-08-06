package hrouter

const (
	GET     = 1
	HEAD    = 2
	POST    = 3
	PUT     = 4
	DELETE  = 5
	CONNECT = 6
	OPTIONS = 7
	TRACE   = 8
	PATCH   = 9
)

func MethodStringToNum(method string) int {
	switch method {
	case "GET":
		return GET
	case "HEAD":
		return HEAD
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	case "CONNECT":
		return CONNECT
	case "OPTIONS":
		return OPTIONS
	case "TRACE":
		return TRACE
	case "PATCH":
		return PATCH
	default:
		return 0
	}
}

func MethodStringToFlag(method string) int {
	return 1 << MethodStringToNum(method)
}

type HttpMethod struct {
	flag     int
	method   string
	handlers HandlersChain
}

func (h *HttpMethod) GetMethod() string {
	return h.method
}

func (h *HttpMethod) GetMethodFlag() int {
	return h.flag
}

func (h *HttpMethod) GetHandlers() HandlersChain {
	return h.handlers
}

func newHttpMethod(method string, f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     MethodStringToNum(method),
		method:   method,
		handlers: f,
	}
}

func Get(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("GET", f...)
}

func Head(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("HEAD", f...)
}

func Post(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("POST", f...)
}

func Put(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("PUT", f...)
}

func Delete(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("DELETE", f...)
}

func Connect(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("CONNECT", f...)
}

func Options(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("OPTIONS", f...)
}

func Trace(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("TRACE", f...)
}

func Patch(f ...HandlerFunc) *HttpMethod {
	return newHttpMethod("PATCH", f...)
}

type HttpMethodHandlersM struct {
	methodHandlers [9]HandlersChain
}

func NewHttpMethodHandlersObj() IMethodHandlersManage {
	return new(HttpMethodHandlersM)
}

func (hm *HttpMethodHandlersM) SetMethodHandlers(method string, handlers HandlersChain) error {
	index := MethodStringToNum(method) - 1
	if index < 0 {
		return ErrMethodNonConformity
	}
	hm.methodHandlers[index] = handlers
	return nil
}

func (hm *HttpMethodHandlersM) GetMethodHandlers(method string) (HandlersChain, bool) {
	index := MethodStringToNum(method) - 1
	if index < 0 {
		return nil, false
	}
	return hm.methodHandlers[index], true
}

type HttpRouter struct {
	Router
}

func NewHttpRouter() *HttpRouter {
	router := &HttpRouter{}
	router.SetNewMethodHandlersObjFunc(NewHttpMethodHandlersObj)
	return router
}
