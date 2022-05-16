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

export async function createUserHabitList(habitList) {
	const res = await newAxios.post(
		HABIT_ENDPOINTS.CREATE_HABIT_LIST,
		{ ...habitList }
	)
	return res
}

export async function updateUserHabitList(habitList) {
	const res = await newAxios.put(
		HABIT_ENDPOINTS.UPDATE_HABIT_LIST,
		{ ...habitList }
	)
	return res
}

export async function deleteUserHabitList({habit_name}) {
	const res = await newAxios.delete(
		HABIT_ENDPOINTS.DELETE_HABIT_LIST,
		{ data: { habit_name: habit_name } }
	)
	return res
}