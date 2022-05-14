package tag

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/app/responses"
)

type tagHandler struct {
	service TagServiceInterface
}

func NewTagHandler(tagService TagServiceInterface) *tagHandler {
	return &tagHandler{tagService}
}

// get all tags
func (h *tagHandler) GetTags(c *gin.Context) {
	// set value query parameter
	status := c.Query("status")

	// get data from service using param status
	tags, err := h.service.GetTags(status)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToTagResponses(tags)
	response := helper.SuccessResponseWithData("tags", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get all pagination tags with query parameters e.g => page, limit, order, sort
func (h *tagHandler) Paginate(c *gin.Context) {
	// set app url
	appHOST := c.Request.Host     // www.example.com/
	urlPATH := c.Request.URL.Path // api/v1/handler/function
	appPATH := appHOST + urlPATH  // www.example.com/api/v1/handler/function

	// generate struct helper
	meta := helper.NewPagination()
	query := helper.NewPaginationQuery(c)
	link := helper.NewPaginationLink()

	// get result pagination
	result, err := h.service.Pagination(meta, query, link, appPATH)

	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)
	}

	// convert result
	convertResponse := responses.ToTagResponses(result)
	response := helper.SuccessResponsePaginate("paginate tags", "get", meta, query, link, convertResponse)

	c.JSON(http.StatusOK, response)
}

// get a single tag by param id
func (h *tagHandler) GetTag(c *gin.Context) {
	// convert param id to int
	tagId, _ := strconv.Atoi(c.Param("id"))

	tag, err := h.service.GetTagById(tagId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing tag
	if tag.ID == 0 {
		response := helper.NotFoundResponse("tag")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToTagResponse(tag)
	response := helper.SuccessResponseWithData("tag", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}

func (h *tagHandler) CreateTag(c *gin.Context) {
	// set validate
	var input requests.CreateTagRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		// set response error
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// insert with error
	if err := h.service.CreateTag(input); err != nil {
		// set response error
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("tag", "create")

	// return
	c.JSON(http.StatusOK, response)
}

func (h *tagHandler) UpdateTag(c *gin.Context) {
	// convert param id to int
	tagId, _ := strconv.Atoi(c.Param("id"))

	// validate if tag id doesn't exist
	if tag, _ := h.service.GetTagById(tagId); tag.ID == 0 {
		response := helper.NotFoundResponse("tag") // return not found
		c.JSON(http.StatusNotFound, response)

		return
	}

	// validate input
	var input requests.UpdateTagRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// update tag
	err := h.service.UpdateTag(tagId, input)
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("tag", "update")

	c.JSON(http.StatusOK, response)
}

func (h *tagHandler) UpdateTagStatus(c *gin.Context) {

	tagId, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.UpdateTagStatus(tagId); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("tag", "update")

	c.JSON(http.StatusOK, response)
}

func (h *tagHandler) DeleteTag(c *gin.Context) {
	// convert param id to int
	tagId, _ := strconv.Atoi(c.Param("id"))

	// validate if tag id doesn't exist
	if tag, _ := h.service.GetTagById(tagId); tag.ID == 0 {
		response := helper.NotFoundResponse("tag") // return not found
		c.JSON(http.StatusNotFound, response)

		return
	}

	// delete tag
	h.service.DeleteTag(tagId)

	// set response
	response := helper.SuccessResponse("tag", "delete")

	c.JSON(http.StatusOK, response)
}
