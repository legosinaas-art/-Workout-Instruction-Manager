package handler

import (
	"context"

	"example.com/m/v2/internal/dto"
	"example.com/m/v2/internal/model"
	"github.com/gin-contrib/cors" // Этот импорт важен
	"github.com/gin-gonic/gin"
)

type workoutServiceInstructions interface {
	Create(ctx context.Context, req dto.CreateWorkoutInstructionsRequest) (dto.CreateWorkoutInstructionsResponse, error)
	Update(ctx context.Context, req dto.WorkoutInstructionsDto) (dto.WorkoutInstructionsDto, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (dto.WorkoutInstructionsDto, error)
	GetAll(ctx context.Context) ([]dto.WorkoutInstructionsDto, error)
	AddExercises(ctx context.Context, workoutInstructionsId int, exercises []model.ExerciseInWorkout) error
}

type workoutServiceExercises interface {
	Create(ctx context.Context, req dto.CreateExercisesRequest) (dto.CreateExercisesResponse, error)
	Update(ctx context.Context, req dto.ExercisesDto) (dto.ExercisesDto, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (dto.ExercisesDto, error)
	GetAll(ctx context.Context, rew dto.GetExercisesPaginationRequest) (dto.PaginatedResponse[dto.ExercisesDto], error)
}

type Handler struct {
	workoutServiceInstructions workoutServiceInstructions // rename
	workoutServiceExercises    workoutServiceExercises
}

func NewHandler(workoutServiceInstructions workoutServiceInstructions, workoutServiceExercises workoutServiceExercises) *Handler {
	return &Handler{workoutServiceInstructions: workoutServiceInstructions, workoutServiceExercises: workoutServiceExercises}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Эта часть — самая главная. Она разрешает браузеру общаться с сервером.
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))

	// Главная страница (путь к проекту)
	router.GET("/", func(c *gin.Context) {
		c.File("E:/dev/WorkManager/index.html")
	})

	api := router.Group("/api")
	{
		workout_instructions := api.Group("/instructions")
		{
			workout_instructions.POST("", h.WorkoutInstructionsCreate)
			workout_instructions.PUT("/:id", h.WorkoutInstructionsUpdate)
			workout_instructions.DELETE("/:id", h.WorkoutInstructionsDelete)
			workout_instructions.GET("/:id", h.WorkoutInstructionsGetById)
			workout_instructions.GET("", h.WorkoutInstructionsGetAll)
			workout_instructions.POST("/:id/exercises", h.AddExercisesToWorkout)
		}
		workout_exercises := api.Group("/exercises")
		{
			workout_exercises.POST("", h.ExercisesCreate)
			workout_exercises.PUT("/:id", h.ExercisesUpdate)
			workout_exercises.DELETE("/:id", h.ExercisesDelete)
			workout_exercises.GET("/:id", h.ExercisesGetById)
			workout_exercises.GET("", h.ExercisesGetAll)
		}
	}
	return router
}
