package main

type Request struct {
	Host    *string
	Path    *string
	Port    int
	Url     *string
	Method  *string
	Headers *Headers
}

type Headers struct {
}
