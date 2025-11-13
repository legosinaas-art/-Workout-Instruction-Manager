package handler

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/m/v2/internal/dto"
	"example.com/m/v2/internal/repository/repo_error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) WorkoutInstructionsCreate(c *gin.Context) {
	var input dto.CreateWorkoutInstructionsRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	createWorkoutInstructionsResponse, err := h.workoutServiceInstructions.Create(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createWorkoutInstructionsResponse)
}
func (h *Handler) WorkoutInstructionsUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input dto.WorkoutInstructionsDto
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Id = id
	updateWorkoutInstructionsResponse, err := h.workoutServiceInstructions.Update(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateWorkoutInstructionsResponse)
}
func (h *Handler) WorkoutInstructionsDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.workoutServiceInstructions.Delete(c, id)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "workout deleted"})
}
func (h *Handler) WorkoutInstructionsGetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	workoutInstructionsGetId, err := h.workoutServiceInstructions.GetById(c, id)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, workoutInstructionsGetId)
}
func (h *Handler) WorkoutInstructionsGetAll(c *gin.Context) {
	workoutInstructionsGetAll, err := h.workoutServiceInstructions.GetAll(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, workoutInstructionsGetAll)
}

// AddExercisesToWorkout - новый метод-обработчик
func (h *Handler) AddExercisesToWorkout(c *gin.Context) {
	workoutId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверный ID тренировки")
		return
	}
	var input dto.AddExercisesToWorkoutRequest
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.workoutServiceInstructions.AddExercises(c.Request.Context(), workoutId, input.Exercises)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
