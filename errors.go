package hrouter

import "errors"

var (
	ErrMethodHandlersNotFount = errors.New("MethodHandlersNotFountErr")
	ErrRouteNotFount          = errors.New("ErrRouteNotFount")
	ErrRouteIncompatible      = errors.New("ErrRouteIncompatible")
)
