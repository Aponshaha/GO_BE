package handlers

import (
	"net/http"
	"strconv"

	"ecom/internal/database"
	"ecom/internal/dto"
	"ecom/internal/middleware"
	"ecom/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler creates a new product handler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		service: services.NewProductService(database.GetDB()),
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product in the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Product data"
// @Success 201 {object} middleware.ApiResponse{data=dto.ProductResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.BadRequest(c, err.Error(), "Validation failed")
		return
	}
	
	product, err := h.service.CreateProduct(&req)
	if err != nil {
		middleware.InternalError(c, "Failed to create product")
		return
	}

	middleware.Created(c, product, "Product created successfully")
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products with pagination
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10, max: 100)"
// @Success 200 {object} middleware.ListApiResponse{data=[]dto.ProductResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	page, limit := middleware.PaginationParams(c)

	products, total, err := h.service.GetAllProducts(page, limit)
	if err != nil {
		middleware.InternalError(c, "Failed to retrieve products")
		return
	}

	if products == nil {
		products = []dto.ProductResponse{}
	}

	pages := middleware.CalculatePages(total, limit)
	middleware.ListResponse(c, http.StatusOK, products, page, limit, total, pages, "Products retrieved successfully")
}

// GetProduct godoc
// @Summary Get product by ID
// @Description Retrieve a specific product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} middleware.ApiResponse{data=dto.ProductResponse}
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	productID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid product ID")
		return
	}

	product, err := h.service.GetProductByID(productID)
	if err != nil {
		if err.Error() == "product not found" {
			middleware.NotFound(c, "Product not found")
			return
		}
		middleware.InternalError(c, "Failed to retrieve product")
		return
	}

	middleware.OK(c, product, "Product retrieved successfully")
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update an existing product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body dto.UpdateProductRequest true "Product data"
// @Success 200 {object} middleware.ApiResponse{data=dto.ProductResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid product ID")
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.BadRequest(c, err.Error(), "Validation failed")
		return
	}

	product, err := h.service.UpdateProduct(productID, &req)
	if err != nil {
		if err.Error() == "product not found" {
			middleware.NotFound(c, "Product not found")
			return
		}
		middleware.InternalError(c, "Failed to update product")
		return
	}

	middleware.OK(c, product, "Product updated successfully")
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product (soft delete)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid product ID")
		return
	}

	err = h.service.DeleteProduct(productID)
	if err != nil {
		if err.Error() == "product not found" {
			middleware.NotFound(c, "Product not found")
			return
		}
		middleware.InternalError(c, "Failed to delete product")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}



// GetProductsById godoc
// @Summary Get products by category ID
// @Description Retrieve products by category ID with pagination
// @Tags Products
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Page size (default: 10)"
// @Success 200 {object} middleware.ApiResponse{data=[]dto.ProductResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/products/category/{category_id} [get]
func (h *ProductHandler) GetProductsByCategoryID(c *gin.Context) {
	categoryIDStr := c.Param("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		middleware.BadRequest(c, "Invalid category ID", "Validation failed")
		return
	}

	page, limit := middleware.PaginationParams(c)

	products, total, err := h.service.GetAllProductsByCategory(categoryID, page, limit)
	if err != nil {
		middleware.InternalError(c, "Failed to retrieve products")
		return
	}

	if products == nil {
		products = []dto.ProductResponse{}
	}

	pages := middleware.CalculatePages(total, limit)
	middleware.ListResponse(c, http.StatusOK, products, page, limit, total, pages, "Products retrieved successfully")
}

