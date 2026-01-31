package handlers

import (
	"net/http"
	// "strconv"

	"ecom/internal/database"
	"ecom/internal/dto"
	"ecom/internal/middleware"
	"ecom/internal/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *services.CategoryService
}

// NewCategoryHandler creates a new category handler
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		service: services.NewCategoryService(database.GetDB()),
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new product category
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body dto.CreateCategoryRequest true "Category data"
// @Success 201 {object} middleware.ApiResponse{data=dto.CategoryResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.BadRequest(c, err.Error(), "Validation failed")
		return
	}

	category, err := h.service.CreateCategory(&req)
	if err != nil {
		middleware.InternalError(c, err.Error())
		return
	}

	middleware.Created(c, category, "Category created successfully")
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve all product categories with pagination
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10, max: 100)"
// @Success 200 {object} middleware.ListApiResponse{data=[]dto.CategoryResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	page, limit := middleware.PaginationParams(c)

	categories, total, err := h.service.GetAllCategories(page, limit)
	if err != nil {
		middleware.InternalError(c, err.Error())
		return
	}

	if categories == nil {
		categories = []dto.CategoryResponse{}
	}

	pages := middleware.CalculatePages(total, limit)
	middleware.ListResponse(c, http.StatusOK, categories, page, limit, total, pages, "Categories retrieved successfully")
}

// GetCategory godoc
// @Summary Get category by ID
// @Description Retrieve a specific product category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} middleware.ApiResponse{data=dto.CategoryResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	categoryID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid category ID")
		return
	}

	category, err := h.service.GetCategoryByID(categoryID)
	if err != nil {
		if err.Error() == "category not found" {
			middleware.NotFound(c, "Category not found")
			return
		}
		middleware.InternalError(c, err.Error())
		return
	}

	middleware.OK(c, category, "Category retrieved successfully")
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update an existing product category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param request body dto.UpdateCategoryRequest true "Category data"
// @Success 200 {object} middleware.ApiResponse{data=dto.CategoryResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid category ID")
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.BadRequest(c, err.Error(), "Validation failed")
		return
	}

	category, err := h.service.UpdateCategory(categoryID, &req)
	if err != nil {
		if err.Error() == "category not found" {
			middleware.NotFound(c, "Category not found")
			return
		}
		middleware.InternalError(c, err.Error())
		return
	}

	middleware.OK(c, category, "Category updated successfully")
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete a product category (soft delete)
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} middleware.ApiResponse
// @Failure 400 {object} middleware.ApiResponse
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid category ID")
		return
	}

	err = h.service.DeleteCategory(categoryID)
	if err != nil {
		if err.Error() == "category not found" {
			middleware.NotFound(c, "Category not found")
			return
		}
		middleware.InternalError(c, err.Error())
		return
	}

	middleware.OK(c, nil, "Category deleted successfully")
}

// GetCategoryProducts godoc
// @Summary Get products by category
// @Description Retrieve all products in a specific category with pagination
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10, max: 100)"
// @Success 200 {object} middleware.ListApiResponse{data=[]dto.ProductResponse}
// @Failure 400 {object} middleware.ApiResponse
// @Failure 404 {object} middleware.ApiResponse
// @Failure 500 {object} middleware.ApiResponse
// @Router /api/v1/categories/{id}/products [get]
func (h *CategoryHandler) GetCategoryProducts(c *gin.Context) {
	categoryID, err := middleware.GetIDParam(c, "id")
	if err != nil {
		middleware.BadRequest(c, err.Error(), "Invalid category ID")
		return
	}

	// Verify category exists
	_, err = h.service.GetCategoryByID(categoryID)
	if err != nil {
		if err.Error() == "category not found" {
			middleware.NotFound(c, "Category not found")
			return
		}
		middleware.InternalError(c, err.Error())
		return
	}

	page, limit := middleware.PaginationParams(c)

	products, total, err := h.service.GetProductsByCategory(categoryID, page, limit)
	if err != nil {
		middleware.InternalError(c, err.Error())
		return
	}

	if products == nil {
		products = []dto.ProductResponse{}
	}

	pages := middleware.CalculatePages(total, limit)
	middleware.ListResponse(c, http.StatusOK, products, page, limit, total, pages, "Products retrieved successfully")
}
