package handlers

import (
	"go_final/models"
	"go_final/repositories"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"net/http"
	"strconv"
	"html/template"
)

type UserHandler interface {
	SignInUser(*gin.Context)
	CreateUser(*gin.Context)
	GetUser(*gin.Context)
	GetAllUsers(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	ShowLoginPage(*gin.Context)
	ShowRegisterPage(*gin.Context)
	Home(*gin.Context)
	Base(*gin.Context)
}

type userHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repositories.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) Base(ctx *gin.Context) {

	tmpl, err := template.ParseFiles("templates/base.html")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error loading template")
		return
	}

	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("role")
	username, _ := ctx.Get("username")

	data := gin.H{
		"UserID": userID,
		"Role":   role,
		"UserName": username,
	}

	tmpl.Execute(ctx.Writer, data)	

}

func (h *userHandler) Home(ctx *gin.Context) {

	// Get user ID and role from context if available
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error loading template")
		return
	}

	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("role")
	username, _ := ctx.Get("username")

	data := gin.H{
		"UserID": userID,
		"Role":   role,
		"UserName": username,
	}

	tmpl.Execute(ctx.Writer, data)	

}

func (h *userHandler) ShowRegisterPage(ctx *gin.Context) {
    tmpl, err := template.ParseFiles("templates/register.html")
    if err != nil {
        ctx.String(http.StatusInternalServerError, "Error loading template")
        return
    }
    tmpl.Execute(ctx.Writer, nil)
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
    // Parse form data
    var input models.UserRegister
    if err := ctx.Request.ParseForm(); err != nil {
        ctx.String(http.StatusBadRequest, "Invalid form data")
        return
    }
    input.Name = ctx.Request.FormValue("name")
    input.Email = ctx.Request.FormValue("email")
    input.Password = ctx.Request.FormValue("password")

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
        Role:     models.CUSTOMER_ROLE,
    }

    hashPassword(&user.Password)
    user, err := h.repo.CreateUser(user)
    if err != nil {
        ctx.String(http.StatusInternalServerError, "Error creating user")
        return
    }

    // Redirect to login page after successful registration
    ctx.Redirect(http.StatusFound, "/user/login")
}

func (h *userHandler) ShowLoginPage(ctx *gin.Context) {
    tmpl, err := template.ParseFiles("templates/login.html")
    if err != nil {
        ctx.String(http.StatusInternalServerError, "Error loading template")
        return
    }
    tmpl.Execute(ctx.Writer, nil)
}

func (h *userHandler) SignInUser(ctx *gin.Context) {
    // Parse form data
    var user models.UserLogin
    if err := ctx.Request.ParseForm(); err != nil {
        ctx.String(http.StatusBadRequest, "Invalid form data")
        return
    }
    user.Email = ctx.Request.FormValue("email")
    user.Password = ctx.Request.FormValue("password")

    dbUser, err := h.repo.GetByEmail(user.Email)
    if err != nil {
        ctx.String(http.StatusUnauthorized, "No such user found")
        return
    }

    if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
        token := GenerateToken(dbUser.ID, dbUser.Role, dbUser.Name)
        // Set the token as an HttpOnly cookie
        ctx.SetCookie("jwt", token, 3600*24, "/", "", false, true)
        // Redirect to home page
        ctx.Redirect(http.StatusFound, "/")
        return
    }

    ctx.String(http.StatusUnauthorized, "Password didn't match")
}

func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	user, err := h.repo.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var input models.UserUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("user_id")
	intID, _ := strconv.Atoi(id)

	_, err := h.repo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No such user in database!"})
		return
	}

	var existingUser models.User
	existingUser.ID = uint(intID)
	if input.Name != "" {
		existingUser.Name = input.Name
	}
	if input.Email != "" {
		existingUser.Email = input.Email
	}
	passwordChanged := false
	if input.Password != "" {
		existingUser.Password = input.Password
		hashPassword(&existingUser.Password)
		passwordChanged = true
	}

	updatedUser, err := h.repo.UpdateUser(existingUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.APIUser{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Role:  updatedUser.Role,
	}

	if passwordChanged {
		ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully", "changedFields": user})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "No password change", "changedFields": user})
	}
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("user_id")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)

	user, err := h.repo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

