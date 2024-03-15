package handler

import (
	"github.com/gin-gonic/gin"
	GO_Project "github.com/zhetibayev0410/Go-1-project"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input GO_Project.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input GO_Project.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//id, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
//if err != nil {
//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
//	return
//}
