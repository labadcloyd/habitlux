package models

type HabitList struct {
	ID                   uint   `json:"id"`
	Owner_ID             uint   `json:"owner_id"`
	Habit_Name           string `json:"habit_name"`
	Icon_Url             string `json:"icon_url"`
	Color                string `json:"color"`
	Default_Repeat_Count uint   `json:"default_repeat_count"`
}
