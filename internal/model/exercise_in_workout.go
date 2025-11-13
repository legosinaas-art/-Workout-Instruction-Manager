package model

type ExerciseInWorkout struct {
	Id              int             `json:"id"`
	ExerciseDetails ExerciseDetails `json:"details"`
	OrderNum        int             `json:"orderNum"`
}

type ExerciseDetails struct {
	Weight *int `json:"weight"`
	Reps   *int `json:"reps"`
}
