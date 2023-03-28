package service

import (
	"OwnersAnimals/internal/entities/animal"
	"OwnersAnimals/internal/entities/owner"
	"OwnersAnimals/internal/infrastructure/repo/animalrepo"
	"OwnersAnimals/internal/infrastructure/repo/ownerrepo"
)

type Service struct {
	oRepo *ownerrepo.Owner
	aRepo *animalrepo.Animal
}

func NewService(aRepo *animalrepo.Animal, oRepo *ownerrepo.Owner) *Service {
	a := &Service{
		aRepo: aRepo,
		oRepo: oRepo,
	}
	return a
}

func (s *Service) ReadOwnerByID(id int) (owner.Owner, error) {
	own, err := s.oRepo.ReadOwner(id)
	if err != nil {
		return owner.Owner{}, err
	}
	return own, nil
}

func (s *Service) ReadAnimals(ids []int) (animals []animal.Animal, err error) {
	for _, id := range ids {
		animal, err := s.aRepo.ReadAnimal(id)
		if err != nil {
			return animals, err
		}
		animals = append(animals, animal)
	}

	return animals, nil
}
