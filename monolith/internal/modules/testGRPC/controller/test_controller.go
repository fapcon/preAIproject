package controller

import (
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/testGRPC/service"
)

type GrpcTester interface {
	Hello(w http.ResponseWriter, r *http.Request)
	Bye(w http.ResponseWriter, r *http.Request)
}

type TestGrpcController struct {
	service *service.TestServiceGRPC
}

func NewTestGrpcController(grpc *service.TestServiceGRPC) *TestGrpcController {
	return &TestGrpcController{grpc}
}

func (t *TestGrpcController) Hello(w http.ResponseWriter, r *http.Request) {
	message, err := t.service.Hello()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (t *TestGrpcController) Bye(w http.ResponseWriter, r *http.Request) {
	message, err := t.service.Bye()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
