package ownerrepo

import "OwnersAnimals/internal/entities/owner"

type OwnerStore interface {
	Create(owner owner.Owner)
	Read(id int) (owner.Owner, error)
	ReadAll()
	Update(id int, owner owner.Owner)
	Delete(id int)
}

type Owner struct {
	store OwnerStore
}

func NewOwnerRepo(store OwnerStore) *Owner {
	a := &Owner{
		store: store,
	}
	return a
}

func (o *Owner) CreateOwner(owner owner.Owner) {
	o.store.Create(owner)
}

func (o *Owner) ReadOwner(id int) (owner.Owner, error) {
	ow, err := o.store.Read(id)
	if err != nil {
		return owner.Owner{}, err
	}
	return ow, nil
}

func (o *Owner) ReadAllOwner() {
	o.store.ReadAll()
}

func (o *Owner) UpdateOwner(id int, owner owner.Owner) {
	o.store.Update(id, owner)
}

func (o *Owner) DeleteOwner(id int) {
	o.store.Delete(id)
}
