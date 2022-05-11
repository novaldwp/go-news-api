package province

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/responses"
)

type provinceHandler struct {
	service ProvinceServiceInterface
}

func NewProvinceHandler(provinceService ProvinceServiceInterface) *provinceHandler {
	return &provinceHandler{provinceService}
}

// get all provinces
func (h *provinceHandler) GetProvinces(c *gin.Context) {
	// get data from service using param status
	provinces, err := h.service.GetProvinces()

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToProvinceResponses(provinces)
	response := helper.SuccessResponseWithData("provinces", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single province by param id
func (h *provinceHandler) GetProvince(c *gin.Context) {
	// convert param id to int
	provinceId, _ := strconv.Atoi(c.Param("id"))

	province, err := h.service.GetProvinceById(provinceId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing province
	if province.ID == 0 {
		response := helper.NotFoundResponse("province")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToProvinceResponse(province)
	response := helper.SuccessResponseWithData("province", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}
