// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package store

import (
	"encoding/json"
	"time"
)

type GowebappExercise struct {
	UserID       int64  `json:"user_id"`
	ExerciseName string `json:"exercise_name"`
}

type GowebappImage struct {
	ImageID     int64  `json:"image_id"`
	UserID      int64  `json:"user_id"`
	ContentType string `json:"content_type"`
	ImageData   []byte `json:"image_data"`
}

type GowebappSet struct {
	SetID        int64  `json:"set_id"`
	WorkoutID    int64  `json:"workout_id"`
	ExerciseName string `json:"exercise_name"`
	Weight       int32  `json:"weight"`
	Set1         int64  `json:"set1"`
	Set2         int64  `json:"set2"`
	Set3         int64  `json:"set3"`
}

type GowebappUser struct {
	UserID       int64           `json:"user_id"`
	UserName     string          `json:"user_name"`
	PasswordHash string          `json:"password_hash"`
	Name         string          `json:"name"`
	Config       json.RawMessage `json:"config"`
	CreatedAt    time.Time       `json:"created_at"`
	IsEnabled    bool            `json:"is_enabled"`
}

type GowebappWorkout struct {
	WorkoutID int64     `json:"workout_id"`
	UserID    int64     `json:"user_id"`
	StartDate time.Time `json:"start_date"`
}
