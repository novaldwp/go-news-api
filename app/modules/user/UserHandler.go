package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/modules/city"
	"github.com/novaldwp/go-news-api/app/modules/district"
	"github.com/novaldwp/go-news-api/app/modules/province"
	"github.com/novaldwp/go-news-api/app/modules/role"
	"github.com/novaldwp/go-news-api/app/modules/subdistrict"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/app/responses"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	uService UserServiceInterface
	rService role.RoleServiceInterface
	pService province.ProvinceServiceInterface
	cService city.CityServiceInterface
	dService district.DistrictServiceInterface
	sService subdistrict.SubdistrictServiceInterface
}

func NewUserHandler(uService UserServiceInterface, rService role.RoleServiceInterface, pService province.ProvinceServiceInterface, cService city.CityServiceInterface, dService district.DistrictServiceInterface, sService subdistrict.SubdistrictServiceInterface) *userHandler {
	return &userHandler{uService, rService, pService, cService, dService, sService}
}

// get all users
func (h *userHandler) GetUsers(c *gin.Context) {
	// set value query parameter
	status := c.Query("status")

	// get data from service using param status
	users, err := h.uService.GetUsers(status)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToUserResponses(users)
	response := helper.SuccessResponseWithData("users", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get all pagination tags with query parameters e.g => page, limit, order, sort
func (h *userHandler) Paginate(c *gin.Context) {
	// set app url
	appHOST := c.Request.Host     // www.example.com/
	urlPATH := c.Request.URL.Path // api/v1/handler/function
	appPATH := appHOST + urlPATH  // www.example.com/api/v1/handler/function

	// generate struct helper
	meta := helper.NewPagination()
	query := helper.NewPaginationQuery(c)
	link := helper.NewPaginationLink()

	// get result pagination
	result, err := h.uService.Pagination(meta, query, link, appPATH)

	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)
	}

	// convert result
	convertResponse := responses.ToUserResponses(result)
	response := helper.SuccessResponsePaginate("paginate users", "get", meta, query, link, convertResponse)

	c.JSON(http.StatusOK, response)
}

// get a single user by param id
func (h *userHandler) GetUser(c *gin.Context) {
	// convert param id to int
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := h.uService.GetUserById(userId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing user
	if user.ID == 0 {
		response := helper.NotFoundResponse("user")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToUserResponse(user)
	response := helper.SuccessResponseWithData("user", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	// set validate
	var input requests.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		// set response error
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// validate role user
	roleId := input.RoleId
	if roleU, _ := h.rService.GetRoleById(int(roleId)); roleU.ID == 0 {
		response := helper.NotFoundResponse("role")
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// insert with error
	if err := h.uService.CreateUser(input); err != nil {
		// set response error
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("user", "create")

	// return
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	// convert param id to int
	userId, _ := strconv.Atoi(c.Param("id"))

	// validate if user id doesn't exist
	if user, _ := h.uService.GetUserById(userId); user.ID == 0 {
		response := helper.NotFoundResponse("user") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate input
	var input requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// validate role user
	roleId := input.RoleId
	if roleU, _ := h.rService.GetRoleById(int(roleId)); roleU.ID == 0 {
		response := helper.NotFoundResponse("role")
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// update user
	err := h.uService.UpdateUser(userId, input)
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("user", "update")

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUserStatus(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	// validate if user id doesn't exist
	user, _ := h.uService.GetUserById(userId)
	if user.ID == 0 {
		response := helper.NotFoundResponse("user") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	if err := h.uService.UpdateUserStatus(user); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("user", "update")

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUserDetail(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	// get user by id
	user, _ := h.uService.GetUserById(userId)

	// error handling if user not found
	if user.ID == 0 {
		response := helper.NotFoundResponse("user") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate input request
	var input requests.UpdateUserDetailRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// validate province id
	if prov, _ := h.pService.GetProvinceById(int(input.ProvinceID)); prov.ID == 0 {
		response := helper.NotFoundResponse("province") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate city id
	if cty, _ := h.cService.GetCityById(int(input.CityID)); cty.ID == 0 {
		response := helper.NotFoundResponse("city") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate district id
	if dist, _ := h.dService.GetDistrictById(int(input.DistrictID)); dist.ID == 0 {
		response := helper.NotFoundResponse("district") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate subdistrict id
	if subd, _ := h.sService.GetSubdistrictById(int(input.SubdistrictID)); subd.ID == 0 {
		response := helper.NotFoundResponse("subdistrict") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// update user detail
	if err := h.uService.UpdateUserDetail(userId, input); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("user detail", "update")

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUserPassword(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	// get user by id
	user, _ := h.uService.GetUserById(userId)

	// error handling if user not found
	if user.ID == 0 {
		response := helper.NotFoundResponse("user") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate input request
	var input requests.UpdatePasswordUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// check current password from request with password on database
	checkHashedPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.CurrentPassword))

	if checkHashedPassword != nil {
		response := helper.ErrorCheckPassword()
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// process service update password user
	if err := h.uService.UpdateUserPassword(user, input); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("user", "update")

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	// convert param id to int
	userId, _ := strconv.Atoi(c.Param("id"))

	// validate if user id doesn't exist
	if user, _ := h.uService.GetUserById(userId); user.ID == 0 {
		response := helper.NotFoundResponse("user") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// delete user
	h.uService.DeleteUser(userId)

	// set response
	response := helper.SuccessResponse("user", "delete")

	c.JSON(http.StatusOK, response)
}
