package hrouter

type HandlerFunc func()

type HandlersChain []HandlerFunc

type MethodHandlersChain map[string][]HandlerFunc

type IMethodHandlers interface {
	GetMethod() string
	GetHandlers() HandlersChain
}

type IMethodHandlersManage interface {
	SetMethodHandlers(method string, handlers HandlersChain) error
	GetMethodHandlers(method string) (HandlersChain, bool)
}

type NewMethodHandlersObjFunc func() IMethodHandlersManage
