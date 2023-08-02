package hrouter

import "strings"

type subRoutes map[string]*RouteNode

// 路由节点
type RouteNode struct {
	// /123456 中的 123456，/:id 中的 id
	path string

	// 是否是匹配路由
	isMatch bool

	// 所有 method 都调用的处理链，执行顺序高于 methodHandlers 的处理链
	handlers HandlersChain

	// 不同 method 对于的处理链
	methodHandlers MethodHandlersChain

	// 下级路由 map，不含下级匹配路由
	subRoutes subRoutes

	// 下级匹配路由
	matchRoute *RouteNode
}

// RouteNode 添加 handlers
func (r *RouteNode) Use(handlers ...HandlerFunc) *RouteNode {
	if handlers == nil {
		r.handlers = make(HandlersChain, 0, len(handlers))
	}
	r.handlers = append(r.handlers, handlers...)

	return r
}

// RouteNode 去除 handlers
func (r *RouteNode) RemoUse(handlers ...HandlerFunc) *RouteNode {
	r.handlers = nil

	return r
}

func (r *RouteNode) addMethodHandlers(methodHandlers ...IMethodHandlers) *RouteNode {
	if r.methodHandlers == nil {
		r.methodHandlers = make(MethodHandlersChain)
	}
	for _, mh := range methodHandlers {
		r.methodHandlers[mh.GetMethod()] = mh.GetHandlers()
	}

	return r
}

// 在 RouteNode 下添加 route
func (r *RouteNode) AddRoute(relativePath string, methodHandlers ...IMethodHandlers) *RouteNode {
	if relativePath == "" || relativePath == "/" {
		return nil
	}

	if relativePath[0] == '/' {
		relativePath = relativePath[1:]
	}

	if relativePath[len(relativePath)-1] == '/' {
		relativePath = relativePath[:len(relativePath)-1]
	}

	paths := strings.Split(relativePath, "/")
	root := r
	for _, path := range paths {
		isMatch := false
		if path[0] == ':' {
			path = path[1:]
			isMatch = true
		}

		var node *RouteNode
		if isMatch {
			if root.matchRoute == nil {
				node = &RouteNode{
					path:    path,
					isMatch: isMatch,
				}
				root.matchRoute = node
			} else {
				node = root.matchRoute
			}
		} else {
			var ok bool
			if root.subRoutes == nil {
				root.subRoutes = make(subRoutes)
				ok = false
			} else {
				node, ok = root.subRoutes[path]
			}
			if !ok {
				node = &RouteNode{
					path:    path,
					isMatch: isMatch,
				}
				root.subRoutes[path] = node
			}
		}
		// 如果root 有 handlers 且 node 没有 handlers 时 把 root 的 handlers copy 到 node
		if root.handlers != nil && node.handlers == nil {
			node.handlers = root.handlers
		}
		root = node
	}
	root.addMethodHandlers(methodHandlers...)

	return root
}
