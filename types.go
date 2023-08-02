package hrouter

type HandlerFunc func()

type HandlersChain []HandlerFunc

type MethodHandlersChain map[string][]HandlerFunc

type IMethodHandlers interface {
	GetMethod() string
	GetHandlers() HandlersChain
}
