package models


type HabitList struct {
	ID                    uint   							`json:"id" gorm:"index"`
	Owner_ID              uint 								`json:"owner_id" gorm:"index;"`
	Habit_Name						string							`json:"habit_name" gorm:"unique;index;"`
}
