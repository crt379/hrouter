package hrouter

import "errors"

var (
	ErrMethodHandlersNotFount = errors.New("MethodHandlersNotFountErr")
	ErrRouteNotFount          = errors.New("RouteNotFountErr")
	ErrRouteIncompatible      = errors.New("RouteIncompatibleErr")
	ErrMethodNonConformity    = errors.New("MethodNonConformityErr")
)
