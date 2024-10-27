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
    "io"
    "mime/multipart"
    "path/filepath"
    "time"
    "sort"

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
    UpdateBasket(*gin.Context)
    AddToBasket(*gin.Context)
    GetBasket(*gin.Context)
}

type productHandler struct {
    repo repositories.ProductRepository
    basketRepo repositories.BasketRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repositories.NewProductRepository(),
        basketRepo: repositories.NewBasketRepository(),
	}
}

func (h *productHandler) GetBasket(ctx *gin.Context) {
    // Extract userID using the helper function
    userIDInterface, exists := ctx.Get("userID")
    var userID uint
    if exists {
        switch v := userIDInterface.(type) {
        case float64:
            userID = uint(v)
        case uint:
            userID = v
        case int:
            userID = uint(v)
        default:
            fmt.Printf("Unexpected type for userID: %T\n", v)
            userID = 0
        }
    } else {
        // Handle unauthenticated users
        userID = 0
    }

    basket, err := h.basketRepo.GetBasket(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve basket"})
        return
    }

    // Calculate total
    total := 0
    for _, item := range basket.BasketItems {
        total += item.Product.Price * item.Quantity
    }

    // Define template functions if not already defined globally
    funcMap := template.FuncMap{
        "multiply": func(a, b int) int {
            return a * b
        },
        "calculateTotal": func(items []models.BasketItem) int {
            total := 0
            for _, item := range items {
                total += item.Product.Price * item.Quantity
            }
            return total
        },
    }

    // Parse templates with functions
    tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/basket.go.tmpl")
    if err != nil {
        fmt.Printf("Error loading templates: %v\n", err)
        http.Error(ctx.Writer, "Error loading templates", http.StatusInternalServerError)
        return
    }

    sort.Slice(basket.BasketItems, func(i, j int) bool {
        return basket.BasketItems[i].Product.Name < basket.BasketItems[j].Product.Name
    })

    data := gin.H{
        "Basket":   basket.BasketItems,
        "Username": ctx.GetString("username"),
    }

    // Execute the "base" template with "basket.gohtml" as the content
    err = tmpl.ExecuteTemplate(ctx.Writer, "base", data)
    if err != nil {
        fmt.Printf("Error rendering template: %v\n", err)
        http.Error(ctx.Writer, "Error rendering template", http.StatusInternalServerError)
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

    // Safely retrieve userID from context
    userIDInterface, exists := ctx.Get("userID")
    var userID uint
    if exists {
        switch v := userIDInterface.(type) {
        case float64:
            userID = uint(v)
        case uint:
            userID = v
        case int:
            userID = uint(v)
        default:
            fmt.Printf("Unexpected type for userID: %T\n", v)
            userID = 0
        }
    } else {
        // Handle unauthenticated users
        userID = 0
    }

    // Safely retrieve username from context
    usernameInterface, _ := ctx.Get("username")
    var username string
    if usernameInterface != nil {
        username, _ = usernameInterface.(string)
    } else {
        username = "Guest"
    }

    // Group products by category
    categories := make(map[string][]models.Product)
    for _, product := range products {
        categories[product.Category] = append(categories[product.Category], product)
    }

    // Get user's basket
    var basket models.Basket
    if userID != 0 {
        basket, err = h.basketRepo.GetBasket(userID)
        if err != nil {
            fmt.Printf("Error getting basket: %v\n", err)
            basket = models.Basket{}
        }
    } else {
        basket = models.Basket{}
    }

    sort.Slice(basket.BasketItems, func(i, j int) bool {
        return basket.BasketItems[i].Product.Name < basket.BasketItems[j].Product.Name
    })

    // Define template functions
    funcMap := template.FuncMap{
        "multiply": func(a, b int) int {
            return a * b
        },
        "calculateTotal": func(items []models.BasketItem) int {
            total := 0
            for _, item := range items {
                total += item.Product.Price * item.Quantity
            }
            return total
        },
    }

    // Parse templates with functions
    tmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles("templates/base.html", "templates/products.go.tmpl")
    if err != nil {
        fmt.Printf("Error loading templates: %v\n", err)
        http.Error(ctx.Writer, "Error loading templates", http.StatusInternalServerError)
        return
    }

    data := gin.H{
        "Categories": categories,
        "Username":   username,
        "Basket":     basket.BasketItems,
    }

    // Execute the "base" template
    err = tmpl.ExecuteTemplate(ctx.Writer, "base", data)
    if err != nil {
        fmt.Printf("Error rendering template: %v\n", err)
        http.Error(ctx.Writer, "Error rendering template", http.StatusInternalServerError)
    }
}




func (h *productHandler) AddToBasket(c *gin.Context) {
    userIDInterface, exists := c.Get("userID")
    var userID uint
    if exists {
        switch v := userIDInterface.(type) {
        case float64:
            userID = uint(v)
        case uint:
            userID = v
        case int:
            userID = uint(v)
        default:
            fmt.Printf("Unexpected type for userID: %T\n", v)
            userID = 0
        }
    } else {
        // Handle unauthenticated users
        userID = 0
    }


    var requestData struct {
        ProductID uint `json:"product_id"`
    }
    if err := c.ShouldBindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    err := h.basketRepo.UpdateBasket(userID, requestData.ProductID, 1)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to basket"})
        return
    }

    // Get updated basket
    basket, err := h.basketRepo.GetBasket(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve basket"})
        return
    }

    sort.Slice(basket.BasketItems, func(i, j int) bool {
        return basket.BasketItems[i].Product.Name < basket.BasketItems[j].Product.Name
    })

    c.JSON(http.StatusOK, gin.H{"basket": basket.BasketItems})
}


func (h *productHandler) UpdateBasket(c *gin.Context) {
    userIDInterface, exists := c.Get("userID")
    var userID uint
    if exists {
        switch v := userIDInterface.(type) {
        case float64:
            userID = uint(v)
        case uint:
            userID = v
        case int:
            userID = uint(v)
        default:
            fmt.Printf("Unexpected type for userID: %T\n", v)
            userID = 0
        }
    } else {
        // Handle unauthenticated users
        userID = 0
    }


    var requestData struct {
        ProductID uint `json:"product_id"`
        Cnt       int  `json:"cnt"`
    }
    if err := c.ShouldBindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    err := h.basketRepo.UpdateBasket(userID, requestData.ProductID, requestData.Cnt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update basket"})
        return
    }

    // Get updated basket
    basket, err := h.basketRepo.GetBasket(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve basket"})
        return
    }

    sort.Slice(basket.BasketItems, func(i, j int) bool {
        return basket.BasketItems[i].Product.Name < basket.BasketItems[j].Product.Name
    })

    c.JSON(http.StatusOK, gin.H{"basket": basket.BasketItems})
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
    fileBytes, err := io.ReadAll(file)
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
