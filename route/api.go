package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_final/handlers"
	"go_final/middleware"
)

func RunAPI(address string) error {
	// Initialize handlers
	userHandler := handlers.NewUserHandler()
	productHandler := handlers.NewProductHandler()
	orderHandler := handlers.NewOrderHandler()

	// Initialize Gin router with default middleware (logger and recovery)
	r := gin.Default()


	// Static files (CSS, JS, images)
	r.Static("/assets", "./assets")

	// ====================
	// Web Routes
	// ====================

	// Home Page Route
	r.GET("/", middleware.AuthorizeJWT(), userHandler.Home)

	// User Routes for Web Pages
	webUserRoutes := r.Group("/user")
	{
		// Show Registration Page
		webUserRoutes.GET("/register", middleware.Already(), userHandler.ShowRegisterPage)

		// Handle Registration Form Submission
		webUserRoutes.POST("/register", userHandler.CreateUser)

		// Show Login Page
		webUserRoutes.GET("/login", middleware.Already(), userHandler.ShowLoginPage)

		// Handle Login Form Submission
		webUserRoutes.POST("/login", userHandler.SignInUser)

		// Handle Logout
		webUserRoutes.GET("/logout", func(ctx *gin.Context) {
			// Clear the JWT cookie
			ctx.SetCookie("jwt", "", -1, "/", "", false, true)
			// Redirect to home page
			ctx.Redirect(http.StatusFound, "/")
		})
	}

	// ====================
	// API Routes
	// ====================

	apiRoutes := r.Group("/api")
	{
		// Public User Routes
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.POST("/register", userHandler.CreateUser)
			userRoutes.POST("/signin", userHandler.SignInUser)
		}

		// Secured User Routes
		userSecuredRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
		{
			userSecuredRoutes.GET("/", userHandler.GetAllUsers)
			userSecuredRoutes.GET("/:user_id", userHandler.GetUser)

			// Admin-only User Routes
			adminRoutes := userSecuredRoutes.Group("/", middleware.CheckAdmin())
			{
				adminRoutes.PUT("/:user_id", userHandler.UpdateUser)
				adminRoutes.DELETE("/:user_id", userHandler.DeleteUser)
			}
		}

		// Product Routes
		productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
		{
			productRoutes.GET("/", productHandler.GetAllProduct)
			productRoutes.GET("/:product_id", productHandler.GetProduct)

			// Admin-only Product Routes
			adminRoutes := productRoutes.Group("/", middleware.CheckAdmin())
			{
				adminRoutes.POST("/", productHandler.CreateProduct)
				adminRoutes.PUT("/:product_id", productHandler.UpdateProduct)
				adminRoutes.DELETE("/:product_id", productHandler.DeleteProduct)
			}
		}

		// Order Routes
		orderRoutes := apiRoutes.Group("/order", middleware.AuthorizeJWT())
		{
			orderRoutes.POST("/", orderHandler.OrderProducts)
			orderRoutes.GET("/", orderHandler.GetOrders)

			// Admin-only Order Routes
			adminRoutes := orderRoutes.Group("/", middleware.CheckAdmin())
			{
				adminRoutes.GET("/:order_id", orderHandler.GetOrderByID)
				adminRoutes.GET("/:order_id/order_items", orderHandler.GetOrderItems)
				adminRoutes.PUT("/", orderHandler.UpdateOrder)
				adminRoutes.PUT("/order_status/:order_id/:status", orderHandler.UpdateOrderStatus)
				adminRoutes.DELETE("/", orderHandler.DeleteOrder)
				adminRoutes.DELETE("/order_items/:order_item_id", orderHandler.DeleteOrderItem)
			}
		}
	}

	return r.Run(address)
}
