package handlers

import (
	"OwnersAnimals/internal/entities/animal"
	"OwnersAnimals/internal/entities/owner"
	"OwnersAnimals/internal/infrastructure/service"
)

type ServiceHandler interface {
	ReadOwnerByID(id int) (owner.Owner, error)
	ReadAnimals(ids []int) (animals []animal.Animal, err error)
}

type Handlers struct {
	service ServiceHandler
}

func NewHandler(r *service.Service) *Handlers {
	h := &Handlers{
		service: r,
	}
	return h
}

func (h *Handlers) ReadAnimalByID(ids []int) ([]animal.Animal, error) {
	ow, err := h.service.ReadAnimals(ids)
	if err != nil {
		return ow, err
	}
	return ow, nil
}

func (h *Handlers) ReadOwner(id int) (owner.Owner, error) {
	ow, err := h.service.ReadOwnerByID(id)
	if err != nil {
		return ow, err
	}
	return ow, nil
}
