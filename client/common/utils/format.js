import moment from "moment";
import { DATE_CHOICES } from "../constants/dates";
import { getDateBiWeekly, getDateMonthly, getDateWeekly } from "./date";

export function addHabitsToDate({ habits, selectedDatesWithHabits }) {
	const habitsWithAddedDates = []

	habits.forEach((habitlist) => {
		const newHabitlist = {
			...habitlist,
			habits: selectedDatesWithHabits
		}
		selectedDatesWithHabits.forEach((day, index) => {
			for (let i = 0; i < habitlist.habits.length; i++) {
				const current_habit_date = moment(habitlist.habits[i].date_created).format('YYYY-MM-DD')
				const current_loop_date = moment(day.date_created).format('YYYY-MM-DD')
				if ( current_habit_date === current_loop_date) {
					newHabitlist.habits[index] = habitlist.habits[i]
					break
				}
			}
		})

		habitsWithAddedDates.push(newHabitlist)
	})

	return habitsWithAddedDates
	
}