package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_User
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Procedure json
// @Param User body models.UserCreate true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateUser(c *gin.Context) {
	var createUser models.UserCreate
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "Error in Should Bind Json in User", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.User().Create(c.Request.Context(), &createUser)
	if err != nil {
		h.handlerResponse(c, "error in storage.userCreate", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "error in storage.userGetByID", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "create user response", http.StatusCreated, resp)
}



// GetByID User godoc
// @ID get_by_id_User
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
func (h *Handler) GetByIDUser(c *gin.Context) {
	var id string
	id = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "invalid id", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "error in storage GetByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "getById user response", http.StatusOK, resp)
}



// GetList User godoc
// @ID get_list_User
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListUser(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list user offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list user limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "storage GetList User", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list user", http.StatusOK, resp)
}




// Update User godoc
// @ID update_User
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param User body models.UserUpdate true "UpdateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateUser models.UserUpdate
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.handlerResponse(c, "error User should bind json", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.strg.User().Update(c.Request.Context(), &updateUser)
	if err != nil {
		h.handlerResponse(c, "storage.User.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.User.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User resposne", http.StatusAccepted, resp)
}




// Delete User godoc
// @ID delete_User
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteUser(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.User().Delete(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User resposne", http.StatusNoContent, nil)
}
