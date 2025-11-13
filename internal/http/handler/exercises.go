package handler

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/m/v2/internal/dto"
	"example.com/m/v2/internal/repository/repo_error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ExercisesCreate(c *gin.Context) {
	var input dto.CreateExercisesRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	createExercisesResponse, err := h.workoutServiceExercises.Create(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createExercisesResponse)
}
func (h *Handler) ExercisesUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input dto.ExercisesDto
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Id = id
	updateExercisesRequest, err := h.workoutServiceExercises.Update(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateExercisesRequest)
}
func (h *Handler) ExercisesDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.workoutServiceExercises.Delete(c, id)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "exercise deleted"})
}
func (h *Handler) ExercisesGetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	getByIdExercisesRequest, err := h.workoutServiceExercises.GetById(c, id)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, getByIdExercisesRequest)
}
func (h *Handler) ExercisesGetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("perPage", "10")
	searchStr := c.Query("search")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid page parameter")
		return
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid perPage parameter")
		return
	}
	req := dto.GetExercisesPaginationRequest{
		Pagination: dto.Pagination{
			Page:    page,
			PerPage: perPage,
		},
		Filter: dto.Filter{
			Search: searchStr,
		},
	}

	paginatedResponse, err := h.workoutServiceExercises.GetAll(c, req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, paginatedResponse)
}
