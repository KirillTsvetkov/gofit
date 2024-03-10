package handler

import (
	"log"
	"net/http"

	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHander struct {
	service *services.Service
}

func NewAuthHander(service *services.Service) *AuthHander {
	return &AuthHander{service: service}
}

type signRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHander) GetToken(c *gin.Context) {
	request := new(signRequest)
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.service.CreateToken(c, request.Email, request.Password)
	if err != nil {
		log.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHander) singUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign up page"})
}

func (h *AuthHander) singIn(c *gin.Context) {

}
