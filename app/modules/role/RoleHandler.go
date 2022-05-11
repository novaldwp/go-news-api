package role

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/app/responses"
)

type roleHandler struct {
	service RoleServiceInterface
}

func NewRoleHandler(roleService RoleServiceInterface) *roleHandler {
	return &roleHandler{roleService}
}

// get all roles
func (h *roleHandler) GetRoles(c *gin.Context) {
	// set value query parameter
	status := c.Query("status")

	// get data from service using param status
	roles, err := h.service.GetRoles(status)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set success response
	convertResponse := responses.ToRoleResponses(roles)
	response := helper.SuccessResponseWithData("roles", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}

// get a single role by param id
func (h *roleHandler) GetRole(c *gin.Context) {
	// convert param id to int
	roleId, _ := strconv.Atoi(c.Param("id"))

	role, err := h.service.GetRoleById(roleId)

	// check if service return error
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// check existing role
	if role.ID == 0 {
		response := helper.NotFoundResponse("role")
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// set response success
	convertResponse := responses.ToRoleResponse(role)
	response := helper.SuccessResponseWithData("role", "get", convertResponse)

	c.JSON(http.StatusOK, response)
}

func (h *roleHandler) CreateRole(c *gin.Context) {
	// set validate
	var input requests.CreateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		// set response error
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// insert with error
	if err := h.service.CreateRole(input); err != nil {
		// set response error
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("role", "create")

	// return
	c.JSON(http.StatusOK, response)
}

func (h *roleHandler) UpdateRole(c *gin.Context) {
	// convert param id to int
	roleId, _ := strconv.Atoi(c.Param("id"))

	// validate if role id doesn't exist
	if role, _ := h.service.GetRoleById(roleId); role.ID == 0 {
		response := helper.NotFoundResponse("role") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// validate input
	var input requests.UpdateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return // break
	}

	// update role
	err := h.service.UpdateRole(roleId, input)
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	// set response
	response := helper.SuccessResponse("role", "update")

	c.JSON(http.StatusOK, response)
}

func (h *roleHandler) UpdateRoleStatus(c *gin.Context) {

	roleId, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.UpdateRoleStatus(roleId); err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return // break
	}

	response := helper.SuccessResponse("role", "update")

	c.JSON(http.StatusOK, response)
}

func (h *roleHandler) DeleteRole(c *gin.Context) {
	// convert param id to int
	roleId, _ := strconv.Atoi(c.Param("id"))

	// validate if role id doesn't exist
	if role, _ := h.service.GetRoleById(roleId); role.ID == 0 {
		response := helper.NotFoundResponse("role") // return not found
		c.JSON(http.StatusNotFound, response)

		return // break
	}

	// delete role
	h.service.DeleteRole(roleId)

	// set response
	response := helper.SuccessResponse("role", "delete")

	c.JSON(http.StatusOK, response)
}
