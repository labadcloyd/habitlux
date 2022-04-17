package controllers

import (
	"habit-tracker/helpers"
)

type ReqHabit struct {
	ID                    uint      					`json:"id"`
	Owner_ID              string    					`json:"owner_id"`
	Date_Created          helpers.Datetime		`json:"date_created"`
	Habit_Name            string    					`json:"habit_name" validate:"required,min=1,max=32"`
	Comment               string    					`json:"comment"`
	Target_Repeat_Count 	uint      					`json:"target_repeat_count"`
	Repeat_Count          uint      					`json:"repeat_count"`
}