import moment from "moment";
import { DATE_CHOICES } from "../constants/dates";
import { getDateBiWeekly, getDateMonthly, getDateWeekly } from "./date";

export function addHabitsToDate({ habits, datesWithHabits }) {
	const habitsWithAddedDates = []

	habits.forEach((habitlist) => {
		console.log(habitlist)
		const newHabitlist = {
			...habitlist,
			habits: [...datesWithHabits]
		}
		datesWithHabits.forEach((day, index) => {
			for (let i = 0; i < habitlist.habits.length; i++) {
				const current_habit_date = moment(habitlist.habits[i].date_created).format('YYYY-MM-DD')
				const current_loop_date = moment(day.date_created).format('YYYY-MM-DD')
				if ( current_habit_date === current_loop_date) {
					const newCurrentHabit = habitlist.habits[i]
					newCurrentHabit.habit_id = habitlist.habits[i].id
					delete newCurrentHabit.id
					newHabitlist.habits[index] = newCurrentHabit
					break
				}
				newHabitlist.habits[index].habit_name = habitlist.habit_name
			}
		})

		habitsWithAddedDates.push(newHabitlist)
	})

	return habitsWithAddedDates
	
}