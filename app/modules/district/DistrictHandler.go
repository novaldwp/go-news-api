package district

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/responses"
)

type districtHandler struct {
	service DistrictServiceInterface
}

func NewDistrictHandler(districtService DistrictServiceInterface) *districtHandler {
	return &districtHandler{districtService}
}

// get all districts
func (h *districtHandler) GetDistricts(c *gin.Context) {
	// get query param city_id and convert from string to int
	cityId, _ := strconv.Atoi(c.Query("city_id"))

	// get data from service using param status
	districts, err := h.service.GetDistricts(cityId)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToDistrictResponses(districts)
	response := helper.SuccessResponseWithData("districts", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single district by param id
func (h *districtHandler) GetDistrict(c *gin.Context) {
	// convert param id to int
	districtId, _ := strconv.Atoi(c.Param("id"))

	district, err := h.service.GetDistrictById(districtId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing district
	if district.ID == 0 {
		response := helper.NotFoundResponse("district")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToDistrictResponse(district)
	response := helper.SuccessResponseWithData("district", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}
