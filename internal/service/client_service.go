package service

import "net/http"

type Repository interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type ClientService struct {
	repository Repository
}

func NewClientService(repository Repository) *ClientService {
	return &ClientService{repository: repository}
}

func (c *ClientService) Create(w http.ResponseWriter, r *http.Request) {
	c.repository.Create(w, r)
}

func (c *ClientService) Read(w http.ResponseWriter, r *http.Request) {
	c.repository.Read(w, r)
}

func (c *ClientService) Update(w http.ResponseWriter, r *http.Request) {
	c.repository.Update(w, r)
}

func (c *ClientService) Delete(w http.ResponseWriter, r *http.Request) {
	c.repository.Delete(w, r)
}
