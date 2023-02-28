package delivery

import (
	"github.com/ssshekhu53/car-model-api/service"
	"net/http"
)

type handler struct {
	service service.Car
}

func New(svc service.Car) handler {
	return handler{svc}
}

func (h handler) Get(w http.ResponseWriter, req *http.Request) {

}

func (h handler) Post(w http.ResponseWriter, req *http.Request) {

}

func (h handler) Put(w http.ResponseWriter, req *http.Request) {

}

func (h handler) Delete(w http.ResponseWriter, req *http.Request) {

}