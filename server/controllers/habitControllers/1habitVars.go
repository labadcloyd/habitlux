package controllers

import (
	"habit-tracker/helpers"
)

type ReqCreateHabit struct {
	Habit_Name						string	    				`json:"habit_name" validate:"required,min=1,max=32"`
	Date_Created          helpers.Datetime		`json:"date_created" validate:"required,min=1,max=32"`
	Comment               string    					`json:"comment"`
	Target_Repeat_Count 	uint      					`json:"target_repeat_count"`
	Repeat_Count          uint      					`json:"repeat_count"`
}

type ReqUpdateHabit struct {
	ID                    uint      					`json:"id" validate:"required"`
	Habit_Name						string	    				`json:"habit_name" validate:"required,min=1,max=32"`
	Date_Created          helpers.Datetime		`json:"date_created" validate:"required,min=1,max=32"`
	Comment               string    					`json:"comment"`
	Target_Repeat_Count 	uint      					`json:"target_repeat_count"`
	Repeat_Count          uint      					`json:"repeat_count"`
}

type ReqDeleteHabit struct {
	ID                    uint      					`json:"id" validate:"required"`
}

type ReqCreateHabitList struct {
	Habit_Name						string  						`json:"habit_name" validate:"required,min=1,max=32"`
}

type ReqUpdateHabitList struct {
	ID                    uint      					`json:"id" validate:"required"`
	Habit_Name						string	    				`json:"habit_name" validate:"required,min=1,max=32"`
}

type ReqDeleteHabitList struct {
	Habit_Name						string  						`json:"habit_name" validate:"required,min=1,max=32"`
}

type ReqGetUserHabits struct {
	Start_Date						helpers.Datetime		`json:"start_date"`
	End_Date							helpers.Datetime		`json:"end_date"`
}