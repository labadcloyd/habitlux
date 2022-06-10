package models

import (
	"habit-tracker/helpers"
)


type Habit struct {
	ID                    uint   							`json:"id"`
	Owner_ID              uint 								`json:"owner_id"`
	Habit_List_ID         uint 								`json:"habit_list_id"`
	Habit_Name		        string 							`json:"habit_name"`
	Date_Created          helpers.Datetime		`json:"date_created"`
	Comment               string 							`json:"comment"`
	Target_Repeat_Count 	uint   							`json:"target_repeat_count"`
	Repeat_Count          uint   							`json:"repeat_count"`
}
