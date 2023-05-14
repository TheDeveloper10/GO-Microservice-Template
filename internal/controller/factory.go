package controller

import "msvc-template/internal/service"

func NewExampleController(svc service.IExample) *Example {
	return &Example{
		svc: svc,
	}
}
