package db

import (
	"OwnersAnimals/internal/entities/owner"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OwnerStore struct {
	db *sqlx.DB
}

func NewOwnerStore(db *sqlx.DB) *OwnerStore {
	oStore := &OwnerStore{
		db: db,
	}
	return oStore
}

func (o *OwnerStore) Create(owner owner.Owner) {
	//TODO implement me
	panic("implement me")
}

func (o *OwnerStore) Read(id int) (owner.Owner, error) {
	own := owner.Owner{}
	switch id {
	case 1:
		own.ID = id
		own.Age = 25
		own.Name = "Василий"
		own.Animals[0] = 1
		own.Animals[1] = 0
	case 2:
		own.ID = id
		own.Age = 34
		own.Name = "Клава"
		own.Animals[0] = 1
	case 3:
		own.ID = id
		own.Age = 42
		own.Name = "Татьяна"
		own.Animals[0] = 0
	default:
		return owner.Owner{}, fmt.Errorf("owner id=%d notfound", id)
	}
	return own, nil
}

func (o *OwnerStore) ReadAll() {
	//TODO implement me
	panic("implement me")
}

func (o *OwnerStore) Update(id int, owner owner.Owner) {
	//TODO implement me
	panic("implement me")
}

func (o *OwnerStore) Delete(id int) {
	//TODO implement me
	panic("implement me")
}
