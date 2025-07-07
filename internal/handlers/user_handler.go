package handlers

import (
	"net/http"
	"strconv"
	"users-crud/internal/dto"
	"users-crud/internal/middleware"
	usecases "users-crud/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(u usecases.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// @Summary Cria um novo usuário
// @Description Registra um novo usuário comum
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "Dados de registro"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.RegisterUser(dto.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// @Summary Login do usuário
// @Description Realiza autenticação e retorna o token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Credenciais de login"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.Login(dto.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// @Summary Lista todos os usuários
// @Description Retorna todos os usuários cadastrados
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /get-all [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar usuários"})
		return
	}

	var response []gin.H
	for _, u := range users {
		response = append(response, gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Busca um usuário por ID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := h.usecase.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// @Summary Atualiza dados de um usuário
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param input body dto.UpdateUserInput true "Dados para atualizar"
// @Success 200 {object} map[string]interface{}
// @Failure 400,403,500 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	targetID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if uint(targetID) != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "sem permissão para editar este usuário"})
		return
	}

	var req dto.UpdateUserInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.usecase.UpdateUser(uint(targetID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    updatedUser.ID,
		"name":  updatedUser.Name,
		"email": updatedUser.Email,
	})
}

// @Summary Remove um usuário
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} map[string]string
// @Failure 400,403,500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	targetID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if uint(targetID) != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "sem permissão para deletar este usuário"})
		return
	}

	err = h.usecase.DeleteUser(uint(targetID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao deletar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "usuário deletado com sucesso"})
}

// @Summary Atualiza o papel (role) de um usuário
// @Description Somente administradores podem alterar o role
// @Tags Admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param input body dto.UpdateRoleInput true "Novo role do usuário"
// @Success 200 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]string
// @Router /admin/users/{id}/role [put]
func (h *UserHandler) UpdateRole(c *gin.Context) {
	idParam := c.Param("id")
	targetID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.UpdateRoleInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.usecase.UpdateUserRole(uint(targetID), req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role atualizada com sucesso",
		"user": gin.H{
			"id":    updatedUser.ID,
			"name":  updatedUser.Name,
			"email": updatedUser.Email,
			"role":  updatedUser.Role,
		},
	})
}
