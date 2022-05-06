import { HABIT_ENDPOINTS } from '../constants'
import { newAxios } from '../utils'

export async function getAllUserHabits({ Start_Date, End_Date }) {
	const res = await newAxios.get(HABIT_ENDPOINTS.GET_ALL, { params: { start_date: Start_Date, end_date: End_Date } })
	return res
}