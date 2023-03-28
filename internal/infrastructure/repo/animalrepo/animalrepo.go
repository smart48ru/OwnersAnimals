package animalrepo

import "OwnersAnimals/internal/entities/animal"

type AnimalStore interface {
	Create(animal animal.Animal)
	Read(id int) (animal.Animal, error)
	ReadAll() ([]animal.Animal, error)
	Update(id int, animal animal.Animal)
	Delete(id int)
}

type Animal struct {
	store AnimalStore
}

func NewAnimalRepo(store AnimalStore) *Animal {
	a := &Animal{
		store: store,
	}
	return a
}

func (a *Animal) CreateAnimal(an animal.Animal) {
	a.store.Create(an)
}

func (a *Animal) ReadAnimal(id int) (animal.Animal, error) {
	an, err := a.store.Read(id)
	if err != nil {
		return animal.Animal{}, err
	}
	return an, nil
}

func (a *Animal) ReadAllAnimal() {
	a.store.ReadAll()
}

func (a *Animal) UpdateAnimal(id int, an animal.Animal) {
	a.store.Update(id, an)
}

func (a *Animal) DeleteAnimal(id int) {
	a.store.Delete(id)
}
