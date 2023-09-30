package controllers

import "net/http"

type Context interface {
	Bind(any) error
	JSON(int, any) error
	Param(string) string
	Request() *http.Request
}
