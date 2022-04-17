package models

import (
	"habit-tracker/helpers"
)


type Habit struct {
	ID                    uint   							`json:"id"`
	Owner_ID              string 							`json:"owner_id"`
	Date_Created          helpers.Datetime		`json:"date_created" gorm:"unique"`
	Habit_Name            string 							`json:"habit_name"`
	Comment               string 							`json:"comment"`
	Target_Repeat_Count 	uint   							`json:"target_repeat_count"`
	Repeat_Count          uint   							`json:"repeat_count"`
}
