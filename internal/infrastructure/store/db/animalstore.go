package db

import (
	"OwnersAnimals/internal/entities/animal"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AnimalStore struct {
	db *sqlx.DB
}

func NewAnimalStore(db *sqlx.DB) *AnimalStore {
	aStore := &AnimalStore{
		db: db,
	}
	return aStore
}

func (o *AnimalStore) Create(animal animal.Animal) {
	//TODO implement me
	panic("implement me")
}

func (o *AnimalStore) Read(id int) (animal.Animal, error) {
	an := animal.Animal{}
	switch id {
	case 1:
		an.ID = id
		an.NickName = "Жучка"
		an.Type = "Собака"
		an.Age = 2
		an.Color = "Черный"
	case 2:
		an.ID = id
		an.NickName = "Мурка"
		an.Type = "Кошка"
		an.Age = 5
		an.Color = "Рыжая"
	default:
		return an, fmt.Errorf("animal id=%d not found", id)
	}
	return an, nil
}

func (o *AnimalStore) ReadAll() ([]animal.Animal, error) {
	an := animal.Animal{}
	var animals []animal.Animal
	an.ID = 1
	an.NickName = "Мурка"
	an.Type = "Кошка"
	an.Age = 5
	an.Color = "Рыжая"
	animals[0] = an
	an.ID = 2
	an.NickName = "Мурка"
	an.Type = "Кошка"
	an.Age = 5
	an.Color = "Рыжая"
	animals[1] = an

	return animals, nil
}

func (o *AnimalStore) Update(id int, animal animal.Animal) {
	//TODO implement me
	panic("implement me")
}

func (o *AnimalStore) Delete(id int) {
	//TODO implement me
	panic("implement me")
}
