package hrouter

import (
	"strings"
)

type Router struct {
	tree *RouteNode
}

func NewRouter() *Router {
	return &Router{
		tree: newTree(),
	}
}

func newTree() *RouteNode {
	return &RouteNode{
		share: new(routeShare),
	}
}

func (r *Router) SetNewMethodHandlersObjFunc(hm NewMethodHandlersObjFunc) {
	if r.tree == nil {
		r.tree = newTree()
	}
	r.tree.share.newMethodHandlersObjFunc = hm
}

// 添加路由
func (r *Router) AddRoute(absolutePath string, methodHandlers ...IMethodHandlers) *RouteNode {
	if absolutePath == "" || absolutePath[0] != '/' {
		return nil
	}
	if r.tree == nil {
		r.tree = newTree()
	}
	root := r.tree
	if absolutePath == "/" {
		root.addMethodHandlers(methodHandlers...)
		return root
	}

	return root.AddRoute(absolutePath, methodHandlers...)
}

// 获取传入路由路径的对于 method 的 HandlersChain
func (r *Router) GetRouteHandlers(absolutePath string, method string) (HandlersChain, error) {
	node, err := r.GetRouteNode(absolutePath)
	if err != nil {
		return nil, err
	}
	mfs, ok := node.methodHandlersObj.GetMethodHandlers(method)
	if !ok || len(mfs) < 1 {
		return node.handlers, ErrMethodHandlersNotFount
	}
	if len(node.handlers) > 0 {
		return append(node.handlers, mfs...), nil
	}

	return mfs, nil
}

// 获取传入路由路径的 RouteNode
func (r *Router) GetRouteNode(absolutePath string) (*RouteNode, error) {
	if absolutePath == "" || absolutePath[0] != '/' {
		return nil, ErrRouteIncompatible
	}
	if absolutePath == "/" {
		return r.tree, nil
	}

	if absolutePath[len(absolutePath)-1] == '/' {
		absolutePath = absolutePath[:len(absolutePath)-1]
	}

	root := r.tree
	paths := strings.Split(absolutePath[1:], "/")
	for _, path := range paths {
		if path[0] == ':' {
			path = path[1:]
		}

		node, ok := root.subRoutes[path]
		if !ok {
			// 没有匹配下级路由
			if root.matchRoute == nil {
				return nil, ErrRouteNotFount
			}
			node = root.matchRoute
		}
		root = node
	}

	return root, nil
}
