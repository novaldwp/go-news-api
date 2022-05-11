package city

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/responses"
)

type cityHandler struct {
	service CityServiceInterface
}

func NewCityHandler(cityService CityServiceInterface) *cityHandler {
	return &cityHandler{cityService}
}

// get all cities
func (h *cityHandler) GetCities(c *gin.Context) {
	// get query param province_id and convert from string to int
	provinceId, _ := strconv.Atoi(c.Query("province_id"))

	// get data from service using param status
	cities, err := h.service.GetCities(provinceId)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToCityResponses(cities)
	response := helper.SuccessResponseWithData("cities", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single city by param id
func (h *cityHandler) GetCity(c *gin.Context) {
	// convert param id to int
	cityId, _ := strconv.Atoi(c.Param("id"))

	city, err := h.service.GetCityById(cityId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing city
	if city.ID == 0 {
		response := helper.NotFoundResponse("city")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToCityResponse(city)
	response := helper.SuccessResponseWithData("city", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}
