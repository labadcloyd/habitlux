package models

import (
	"habit-tracker/helpers"
)


type Habit struct {
	ID                    uint   							`json:"id" gorm:"index"`
	Owner_ID              uint 								`json:"owner_id" gorm:"index;"`
	Habit_Name		        string 							`json:"habit_name" gorm:"index;not null;default:null"`
	Date_Created          helpers.Datetime		`json:"date_created" gorm:"index;not null;default:null"`
	Comment               string 							`json:"comment"`
	Target_Repeat_Count 	uint   							`json:"target_repeat_count"`
	Repeat_Count          uint   							`json:"repeat_count"`
}
