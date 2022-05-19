package models


type HabitList struct {
	ID                    uint   							`json:"id" gorm:"index"`
	Owner_ID              uint 								`json:"owner_id" gorm:"index;"`
	Habit_Name						string							`json:"habit_name" gorm:"index;"`
	Icon_Url							string							`json:"icon_url"`
	Color									string							`json:"color"`
	Default_Repeat_Count	uint								`json:"default_repeat_count"`
}
