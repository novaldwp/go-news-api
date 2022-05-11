package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/app/responses"
)

type categoryHandler struct {
	service CategoryServiceInterface
}

func NewCategoryHandler(categoryService CategoryServiceInterface) *categoryHandler {
	return &categoryHandler{categoryService}
}

// get all categories
func (h *categoryHandler) GetCategories(c *gin.Context) {
	// set value query parameter
	status := c.Query("status")

	// get data from service using param status
	categories, err := h.service.GetCategories(status)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToCategoryResponses(categories)
	response := helper.SuccessResponseWithData("categories", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single category by param id
func (h *categoryHandler) GetCategory(c *gin.Context) {
	// convert param id to int
	categoryId, _ := strconv.Atoi(c.Param("id"))

	category, err := h.service.GetCategoryById(categoryId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing category
	if category.ID == 0 {
		response := helper.NotFoundResponse("category")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToCategoryResponse(category)
	response := helper.SuccessResponseWithData("category", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	// set validate
	var input requests.CreateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		// set response error
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// insert with error
	if err := h.service.CreateCategory(input); err != nil {
		// set response error
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("category", "create")

	// return
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	// convert param id to int
	categoryId, _ := strconv.Atoi(c.Param("id"))

	// validate if category id doesn't exist
	if category, _ := h.service.GetCategoryById(categoryId); category.ID == 0 {
		response := helper.NotFoundResponse("category") // return not found
		c.JSON(http.StatusNotFound, response)

		return
	}

	// validate input
	var input requests.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// update category
	err := h.service.UpdateCategory(categoryId, input)
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("category", "update")

	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdateCategoryStatus(c *gin.Context) {

	categoryId, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.UpdateCategoryStatus(categoryId); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("category", "update")

	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	// convert param id to int
	categoryId, _ := strconv.Atoi(c.Param("id"))

	// validate if category id doesn't exist
	if category, _ := h.service.GetCategoryById(categoryId); category.ID == 0 {
		response := helper.NotFoundResponse("category") // return not found
		c.JSON(http.StatusNotFound, response)

		return
	}

	// delete category
	h.service.DeleteCategory(categoryId)

	// set response
	response := helper.SuccessResponse("category", "delete")

	c.JSON(http.StatusOK, response)
}
