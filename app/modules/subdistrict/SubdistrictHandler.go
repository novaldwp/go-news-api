package subdistrict

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/responses"
)

// TODO => SUBDISTRICT HANDLER >>> SUBDISTRICT ROUTE >>> FINISH
type subdistrictHandler struct {
	service SubdistrictServiceInterface
}

func NewSubdistrictHandler(districtService SubdistrictServiceInterface) *subdistrictHandler {
	return &subdistrictHandler{districtService}
}

// get all districts
func (h *subdistrictHandler) GetSubdistricts(c *gin.Context) {
	// get query param city_id and convert from string to int
	districtId, _ := strconv.Atoi(c.Query("district_id"))

	// get data from service using param status
	subdistricts, err := h.service.GetSubdistricts(districtId)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToSubdistrictResponses(subdistricts)
	response := helper.SuccessResponseWithData("subdistricts", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single district by param id
func (h *subdistrictHandler) GetSubdistrict(c *gin.Context) {
	// convert param id to int
	districtId, _ := strconv.Atoi(c.Param("id"))

	subdistrict, err := h.service.GetSubdistrictById(districtId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing district
	if subdistrict.ID == 0 {
		response := helper.NotFoundResponse("subdistrict")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToSubdistrictResponse(subdistrict)
	response := helper.SuccessResponseWithData("subdistrict", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}
