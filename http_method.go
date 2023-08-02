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

func MethodStringToFlag(method string) int {
	switch method {
	case "GET":
		return 1 << GET
	case "HEAD":
		return 1 << HEAD
	case "POST":
		return 1 << POST
	case "PUT":
		return 1 << PUT
	case "DELETE":
		return 1 << DELETE
	case "CONNECT":
		return 1 << CONNECT
	case "OPTIONS":
		return 1 << OPTIONS
	case "TRACE":
		return 1 << TRACE
	case "PATCH":
		return 1 << PATCH
	default:
		return 0
	}
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

func Get(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << GET,
		method:   "GET",
		handlers: f,
	}
}

func Head(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << HEAD,
		method:   "HEAD",
		handlers: f,
	}
}

func Post(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << POST,
		method:   "POST",
		handlers: f,
	}
}

func Put(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << PUT,
		method:   "PUT",
		handlers: f,
	}
}

func Delete(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << DELETE,
		method:   "DELETE",
		handlers: f,
	}
}

func Connect(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << CONNECT,
		method:   "CONNECT",
		handlers: f,
	}
}

func Options(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << OPTIONS,
		method:   "OPTIONS",
		handlers: f,
	}
}

func Trace(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << TRACE,
		method:   "TRACE",
		handlers: f,
	}
}

func Patch(f ...HandlerFunc) *HttpMethod {
	return &HttpMethod{
		flag:     1 << PATCH,
		method:   "PATCH",
		handlers: f,
	}
}
