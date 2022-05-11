package subcategory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/modules/category"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/app/responses"
)

type subcategoryHandler struct {
	sService SubcategoryServiceInterface
	cService category.CategoryServiceInterface
}

func NewSubcategoryHandler(subcategoryService SubcategoryServiceInterface, cService category.CategoryServiceInterface) *subcategoryHandler {
	return &subcategoryHandler{subcategoryService, cService}
}

// get all categories
func (h *subcategoryHandler) GetSubcategories(c *gin.Context) {
	// set value query parameter
	status := c.Query("status")

	// get data from service using param status
	subcategories, err := h.sService.GetSubcategories(status)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToSubcategoryResponses(subcategories)
	response := helper.SuccessResponseWithData("subcategories", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single category by param id
func (h *subcategoryHandler) GetSubcategory(c *gin.Context) {
	// convert param id to int
	subcategoryId, _ := strconv.Atoi(c.Param("id"))

	subcategory, err := h.sService.GetSubcategoryById(subcategoryId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing category
	if subcategory.ID == 0 {
		response := helper.NotFoundResponse("subcategory")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToSubcategoryResponse(subcategory)
	response := helper.SuccessResponseWithData("subcategory", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}

func (h *subcategoryHandler) CreateSubcategory(c *gin.Context) {
	// set validate
	var input requests.CreateSubcategoryRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		// set response error
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// validate categoryId if not exist in table categories
	categoryId := input.CategoryId
	if cat, _ := h.cService.GetCategoryById(int(categoryId)); cat.ID == 0 {
		response := helper.NotFoundResponse("category")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// insert with error
	if err := h.sService.CreateSubcategory(input); err != nil {
		// set response error
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("subcategory", "create")

	// return
	c.JSON(http.StatusOK, response)
}

func (h *subcategoryHandler) UpdateSubcategory(c *gin.Context) {
	// convert param id to int
	subcategoryId, _ := strconv.Atoi(c.Param("id"))

	// validate if subcategory id doesn't exist
	if subcategory, _ := h.sService.GetSubcategoryById(subcategoryId); subcategory.ID == 0 {
		response := helper.NotFoundResponse("subcategory") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate input
	var input requests.UpdateSubcategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)

		c.JSON(http.StatusBadRequest, response)
		return // break
	}

	// validate categoryId if not exist in table categories
	categoryId := input.CategoryId
	if cat, _ := h.cService.GetCategoryById(int(categoryId)); cat.ID == 0 {
		response := helper.NotFoundResponse("category")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// update category
	err := h.sService.UpdateSubcategory(subcategoryId, input)
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("subcategory", "update")

	c.JSON(http.StatusOK, response)
}

func (h *subcategoryHandler) UpdateSubcategoryStatus(c *gin.Context) {

	subcategoryId, _ := strconv.Atoi(c.Param("id"))

	if err := h.sService.UpdateSubcategoryStatus(subcategoryId); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("subcategory", "update")

	c.JSON(http.StatusOK, response)
}

func (h *subcategoryHandler) DeleteSubcategory(c *gin.Context) {
	// convert param id to int
	subcategoryId, _ := strconv.Atoi(c.Param("id"))

	// validate if category id doesn't exist
	if subcategory, _ := h.sService.GetSubcategoryById(subcategoryId); subcategory.ID == 0 {
		response := helper.NotFoundResponse("subcategory") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// delete category
	h.sService.DeleteSubcategory(subcategoryId)

	// set response
	response := helper.SuccessResponse("subcategory", "delete")

	c.JSON(http.StatusOK, response)
}
