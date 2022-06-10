package controllers

import (
	"habit-tracker/helpers"
	"habit-tracker/models"
)

type ReqCreateHabit struct {
	Habit_Name						string	    				`json:"habit_name" validate:"required,min=1,max=32"`
	Habit_List_ID         uint 								`json:"habit_list_id" validate:"required,min=1,max=32"`
	Date_Created          helpers.Datetime		`json:"date_created" validate:"required,min=1,max=32"`
	Comment               string    					`json:"comment"`
	Target_Repeat_Count 	uint      					`json:"target_repeat_count"`
	Repeat_Count          uint      					`json:"repeat_count"`
}

type ReqUpdateHabit struct {
	ID                    uint      					`json:"id" validate:"required"`
	Habit_Name						string	    				`json:"habit_name" validate:"required,min=1,max=32"`
	Habit_List_ID         uint 								`json:"habit_list_id" validate:"required,min=1,max=32"`
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
	Icon_Url							string							`json:"icon_url"`
	Color									string							`json:"color" validate:"required"`
	Default_Repeat_Count	uint								`json:"default_repeat_count"`
}

type ReqUpdateHabitList struct {
	ID                    uint      					`json:"id" validate:"required"`
	Habit_Name						string	    				`json:"habit_name" validate:"required,min=1,max=32"`
	Icon_Url							string							`json:"icon_url" `
	Color									string							`json:"color" validate:"required"`
	Default_Repeat_Count	uint								`json:"default_repeat_count" `
}

type ReqDeleteHabitList struct {
	ID										uint 								`json:"id" validate:"required,min=1,max=32"`
}

type ReqGetUserHabits struct {
	Start_Date						string							`query:"start_date" validate:"required,min=1,max=32"`
	End_Date							string							`query:"end_date" validate:"required,min=1,max=32"`
}


type ResGetUserHabits struct {
	ID                    uint   							`json:"id"`
	Owner_ID              uint 								`json:"owner_id"`
	Habit_Name						string							`json:"habit_name"`
	Icon_Url							string							`json:"icon_url"`
	Color									string							`json:"color"`
	Default_Repeat_Count	uint								`json:"default_repeat_count"`
	Habits								[]models.Habit			`json:"habits"`
}