package handlers

import (
	"go_final/models"
	"go_final/repositories"

	"net/http"
	"strconv"
	"html/template"

	"github.com/gin-gonic/gin"

	"bytes"
    "fmt"
    "io/ioutil"
    "mime/multipart"
    "path/filepath"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	CreateProduct(*gin.Context)
	GetCreateProduct(*gin.Context)
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

func (h *productHandler) GetCreateProduct(ctx *gin.Context) {
	tmpl, err := template.ParseFiles("templates/new_product.html")
    if err != nil {
        ctx.String(http.StatusInternalServerError, "Error loading template")
        return
    }
    tmpl.Execute(ctx.Writer, nil)
}

func (h *productHandler) GetAllProduct(ctx *gin.Context) {
    products, err := h.repo.GetAllproduct()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Group products by category
    categories := make(map[string][]models.Product)
    for _, product := range products {
        categories[product.Category] = append(categories[product.Category], product)
    }

    // Parse templates
    tmpl, err := template.ParseFiles("templates/base.html", "templates/products.html")
    if err != nil {
        fmt.Printf("Error loading templates: %v\n", err)
        http.Error(ctx.Writer, "Error loading templates", http.StatusInternalServerError)
        return
    }

    data := gin.H{
        "Categories": categories,
    }

    // Execute the "base" template
    err = tmpl.ExecuteTemplate(ctx.Writer, "base", data)
    if err != nil {
        fmt.Printf("Error rendering template: %v\n", err)
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

func uploadToS3(file multipart.File, filename string) (string, error) {
    // Initialize AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"), // Update to your region
    })
    if err != nil {
        return "", err
    }

    // Read the file content
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }

    // Create S3 service client
    svc := s3.New(sess)

    // Upload the file to S3
    bucket := "grab-n-go" // Replace with your bucket name
    key := "images/" + filename

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(http.DetectContentType(fileBytes)),
	})

    if err != nil {
        return "", err
    }

    // Construct the file URL
    s3URL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)

    return s3URL, nil
}


func (h *productHandler) CreateProduct(c *gin.Context) {
    // Parse form data
    name := c.PostForm("name")
    quantity := c.PostForm("quantity")
    description := c.PostForm("description")
    price := c.PostForm("price")

    // Validate and convert numeric fields
    qty, err := strconv.Atoi(quantity)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
        return
    }

    prc, err := strconv.Atoi(price)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
        return
    }

    // Get the uploaded file
    file, header, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
        return
    }
    defer file.Close()

    // Generate a unique filename
    filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))

    // Upload the file to S3
	s3URL, err := uploadToS3(file, filename)
    if err != nil {
        // Include the error details in the response
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to upload image",
            "details": err.Error(),
        })
        return
    }

    // Create the product model
    product := models.Product{
        Name:        name,
        Quantity:    qty,
        Description: description,
        Price:       prc,
        Image:       s3URL,
    }

    // Save the product to the database using the repository
    _, err = h.repo.AddProduct(product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
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
