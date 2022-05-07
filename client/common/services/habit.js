import moment from 'moment'
import { HABIT_ENDPOINTS } from '../constants'
import { newAxios } from '../utils'

export async function getAllUserHabits({ Start_Date, End_Date }) {
	const res = await newAxios.get(HABIT_ENDPOINTS.GET_ALL, { params: { start_date: Start_Date, end_date: End_Date } })
	return res
}

export async function updateUserHabit(habit) {
	const res = await newAxios.put(
		HABIT_ENDPOINTS.UPDATE_HABIT,
		{ ...habit, id: habit.habit_id, date_created: moment(habit.date_created).format("YYYY-MM-DD") }
	)
	return res
}
export async function createUserHabit(habit) {
	const res = await newAxios.post(
		HABIT_ENDPOINTS.CREATE_HABIT,
		{ ...habit, date_created: moment(habit.date_created).format("YYYY-MM-DD") }
	)
	return res
}