package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Carros struct {
	l *log.Logger
}

func NewCarros(lo *log.Logger) *Carros {
	return &Carros{l: lo}
}

func (c *Carros) HandleSearch(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("Handling search")
	n := r.URL.Query().Get("nome")
	rw.Write([]byte(fmt.Sprintf("Voce gostaria de carros da %s", n)))
}

func (c *Carros) HandleList(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("Handling list")
	rw.Write([]byte("You got a list of cars"))
}
