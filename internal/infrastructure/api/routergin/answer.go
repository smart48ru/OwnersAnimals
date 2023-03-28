package routergin

import (
	"OwnersAnimals/internal/entities/animal"
	"OwnersAnimals/internal/entities/owner"
)

type GinAnswerError struct {
	Code    int    `json:"code,omitempty" example:"400"`
	Message string `json:"message,omitempty" example:"BadRequest"`
}

type GinAnswerOK struct {
	Code    int    `json:"code,omitempty" example:"200"`
	Message string `json:"message,omitempty" example:"BadRequest"`
}

type GinAnswerOwner struct {
	Name    string            `json:"name,omitempty" example:"Сергей"`
	Age     int               `json:"age,omitempty" example:"22"`
	Animals []GinAnswerAnimal `json:"animals,omitempty" `
}

func (o *GinAnswerOwner) Serialize(owner owner.Owner, animals []GinAnswerAnimal) {
	o.Age = owner.Age
	o.Name = owner.Name
	o.Age = owner.Age
	o.Animals = animals
}

type GinAnswerAnimal struct {
	NickName string `json:"nick_name,omitempty"`
	Type     string `json:"type,omitempty"`
	Age      int    `json:"age,omitempty"`
	Color    string `json:"color,omitempty"`
}

type GinAnswerAnimals struct {
	Animal []GinAnswerAnimal
}

func (a *GinAnswerAnimals) Serialize(animals []animal.Animal) {
	ginAnimal := GinAnswerAnimal{}
	var GinAnimals []GinAnswerAnimal
	for _, animal := range animals {
		ginAnimal.Color = animal.Color
		ginAnimal.NickName = animal.NickName
		ginAnimal.Type = animal.Type
		GinAnimals = append(GinAnimals, ginAnimal)
	}
}
