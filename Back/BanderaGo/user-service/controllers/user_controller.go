package controllers

import (
	"application/dtos/input"
	"application/facade"
	"application/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserFacade facade.UserFacade
	constants  utils.Constants
}

func NewUserController(facade facade.UserFacade) *UserController {
	return &UserController{UserFacade: facade, constants: utils.DefaultConstants}
}

// @Summary Create a user
// @Description Create a user with data of request
// @Accept json
// @Produce json
// @Param user body input.CreateUserIn true "Datos del usuario a crear"
// @Success 201 {object} output.CreateUserOut
// @Tags Usuarios
// @Router /api/users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var userIn input.CreateUserIn

	if err := c.ShouldBindJSON(&userIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uc.constants.MessageErrorJson})
		return
	}

	userOut, err := uc.UserFacade.CreateUser(userIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": uc.constants.MessageErrorCreation})
		return
	}

	c.JSON(http.StatusCreated, userOut)
}

// @Summary Get all users
// @Description Get a list of all users
// @Produce json
// @Success 200 {array} output.GetUsersOut
// @Tags Usuarios
// @Router /api/users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	usersOut, err := uc.UserFacade.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": uc.constants.MessageErrorGetUsers})
		return
	}

	c.JSON(http.StatusOK, usersOut)
}

// @Summary Get a single user
// @Description Get details of a single user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} output.GetUserOut
// @Tags Usuarios
// @Router /api/users/{id} [get]
func (uc *UserController) GetSingleUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uc.constants.MessageErrorID})
		return
	}
	userOut, err := uc.UserFacade.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": uc.constants.MessageErrorUserNotFount})
		return
	}

	c.JSON(http.StatusOK, userOut)
}

// @Summary Update a user
// @Description Update an existing user with new data
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body input.UpdateUserIn true "New user data"
// @Success 200 {object} output.UpdateUserOut
// @Tags Usuarios
// @Router /api/users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uc.constants.MessageErrorID})
		return
	}

	var userIn input.UpdateUserIn
	if err := c.ShouldBindJSON(&userIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uc.constants.MessageErrorJson})
		return
	}

	userOut, err := uc.UserFacade.UpdateUser(uint(userID), userIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": uc.constants.MessageErrorUpdateUser})
		return
	}

	c.JSON(http.StatusOK, userOut)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} output.DeleteUserOut
// @Tags Usuarios
// @Router /api/users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uc.constants.MessageErrorID})
		return
	}

	userOut, err := uc.UserFacade.DeleteUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": uc.constants.MessageErrorDeleteUser})
		return
	}

	c.JSON(http.StatusOK, userOut)
}
