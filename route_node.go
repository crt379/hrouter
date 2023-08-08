package hrouter

import "strings"

// 共享数据
type routeShare struct {
	newMethodHandlersObjFunc NewMethodHandlersObjFunc
}

type routeMap map[string]*RouteNode

// 路由节点
type RouteNode struct {
	// /123456 中的 123456，/:id 中的 id
	path string

	// 是否是匹配路由
	isMatch bool

	// 所有 method 都调用的处理链，执行顺序高于 methodHandlers 的处理链
	handlers HandlersChain

	// 不同 method 的处理链管理对象
	methodHandlersObj IMethodHandlersManage

	// 下级路由 map，不含下级匹配路由
	subRoutes routeMap

	// 下级匹配路由
	matchRoute *RouteNode

	// 共享数据
	share *routeShare
}

func (r *RouteNode) addMethodHandlers(methodHandlers ...IMethodHandlers) *RouteNode {
	if r.methodHandlersObj == nil {
		r.methodHandlersObj = r.share.newMethodHandlersObjFunc()
	}
	for _, mh := range methodHandlers {
		r.methodHandlersObj.SetMethodHandlers(mh.GetMethod(), mh.GetHandlers())
	}

	return r
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

	root := r
	paths := strings.Split(relativePath, "/")
	for _, path := range paths {
		isMatch := false
		if path[0] == ':' {
			path = path[1:]
			isMatch = true
		}

		var node *RouteNode
		if isMatch {
			if root.matchRoute == nil {
				root.matchRoute = &RouteNode{
					path:    path,
					isMatch: isMatch,
					share:   root.share,
				}
			}
			node = root.matchRoute
		} else {
			var ok bool
			if root.subRoutes == nil {
				ok = false
				root.subRoutes = make(routeMap)
			} else {
				node, ok = root.subRoutes[path]
			}
			if !ok {
				node = &RouteNode{
					path:    path,
					isMatch: isMatch,
					share:   root.share,
				}
				root.subRoutes[path] = node
			}
		}
		// 如果root 有 handlers 但 node 没有 handlers 时，把 root 的 handlers copy 到 node
		if root.handlers != nil && node.handlers == nil {
			node.handlers = root.handlers
		}
		root = node
	}
	root.addMethodHandlers(methodHandlers...)

	return root
}
