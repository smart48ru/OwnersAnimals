package routergin

import (
	"OwnersAnimals/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RouterGin struct {
	*gin.Engine
	h *handlers.Handlers
}

func (g *RouterGin) GetAllOwner(c *gin.Context) {

}

func (g *RouterGin) GetOwnerByID(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	c.Header("access-control-allow-origin", "*")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		answer := GinAnswerError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, answer)
		return
	}
	ansOwner := GinAnswerOwner{}
	var ansAnimals GinAnswerAnimals
	own, err := g.h.ReadOwner(intID)
	if err != nil {
		answer := GinAnswerError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, answer)
		return
	}

	an1, err := g.h.ReadAnimalByID(own.Animals)
	if err != nil {
		answer := GinAnswerError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, answer)
		return
	}

	ansAnimals.Serialize(an1)

	c.JSON(http.StatusOK, ansOwner)

}

func (g *RouterGin) CreateOwner(c *gin.Context) {

}

func (g *RouterGin) DeleteOwner(c *gin.Context) {

}

func (g *RouterGin) OpdateOwner(c *gin.Context) {

}

func (g *RouterGin) GetAllAnimal(context *gin.Context) {

}

func NewRouterGin(h *handlers.Handlers) *RouterGin {
	gin.SetMode(gin.ReleaseMode)
	ret := gin.New()
	ret.Use(gin.Recovery())

	ro := &RouterGin{
		h: h,
	}

	owner := ret.RouterGroup
	owner.GET("/owner", ro.GetAllOwner)
	owner.GET("/owner/:id", ro.GetOwnerByID)
	owner.POST("/owner", ro.CreateOwner)
	owner.DELETE("/owner/:id", ro.DeleteOwner)
	owner.PUT("/owner", ro.OpdateOwner)

	animal := ret.RouterGroup
	animal.GET("/animalrepo", ro.GetAllAnimal)
	animal.GET("/animalrepo/:id", ro.GetAllAnimal)
	ro.Engine = ret

	return ro
}
