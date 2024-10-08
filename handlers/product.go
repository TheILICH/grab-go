package handlers

import (
	"go_final/models"
	"go_final/repositories"

	"net/http"
	"strconv"
	"html/template"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	CreateProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct {
	repo repositories.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repositories.NewProductRepository(),
	}
}

func (h *productHandler) GetAllProduct(ctx *gin.Context) {
	product, err := h.repo.GetAllproduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl, err := template.ParseFiles("templates/base.html", "templates/products.html")
    if err != nil {
        http.Error(ctx.Writer, "Error loading templates", http.StatusInternalServerError)
        return
    }

	data := gin.H{
		"Products": product,
	}

	err = tmpl.ExecuteTemplate(ctx.Writer, "base", data) // "base" is the root template
    if err != nil {
        http.Error(ctx.Writer, "Error rendering template", http.StatusInternalServerError)
    }

}

func (h *productHandler) GetProduct(ctx *gin.Context) {
	prodStr := ctx.Param("product_id")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := h.repo.Getproduct(prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}

func (h *productHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.repo.AddProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}
func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prodStr := ctx.Param("product_id")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(prodID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
func (h *productHandler) DeleteProduct(ctx *gin.Context) {

	var product models.Product
	prodStr := ctx.Param("product_id")
	prodID, _ := strconv.Atoi(prodStr)
	product.ID = uint(prodID)
	product, err := h.repo.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}
