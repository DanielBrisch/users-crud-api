package dto

// RegisterInput representa os dados de entrada para criar usuário
type RegisterInput struct {
	Name     string `json:"name" binding:"required,min=2" example:"Daniel" description:"Nome do usuário"`
	Email    string `json:"email" binding:"required,email" example:"daniel@email.com" description:"Email do usuário"`
	Password string `json:"password" binding:"required,min=6" example:"123456" description:"Senha com mínimo 6 caracteres"`
} // @name RegisterInput

// RegisterRequest representa o corpo esperado na requisição de registro
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2" example:"Daniel"`
	Email    string `json:"email" binding:"required,email" example:"daniel@email.com"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
} // @name RegisterRequest

// LoginInput representa os dados de login interno
type LoginInput struct {
	Email    string `json:"email" binding:"required,email" example:"daniel@email.com"`
	Password string `json:"password" binding:"required" example:"123456"`
} // @name LoginInput

// LoginRequest representa os dados da requisição de login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"daniel@email.com"`
	Password string `json:"password" binding:"required" example:"123456"`
} // @name LoginRequest

// UpdateUserInput representa os dados para atualizar nome/email
type UpdateUserInput struct {
	Name  string `json:"name" example:"Daniel Atualizado"`
	Email string `json:"email" example:"novo@email.com"`
} // @name UpdateUserInput

// UpdateRoleInput representa o payload para troca de role
type UpdateRoleInput struct {
	Role string `json:"role" binding:"required,oneof=admin user" example:"admin"`
} // @name UpdateRoleInput
