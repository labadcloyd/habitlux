import moment from 'moment'
import { useEffect, useState } from 'react'
import { DATE_CHOICES, DEFAULT_HABIT_LIST } from '../../../common/constants'
import { createUserHabitList, updateUserHabitList } from '../../../common/services'
import { getDateBiWeekly, getDateMonthly, getDateWeekly } from '../../../common/utils'

import { Close } from '../../../public/svgs'
import { Button, TextInput, ColorPicker, NumberPicker } from '../../common'
import css from './habitModalList.module.css'

export default function HabitModalList(props) {
	const { 
		habitList, 
		isOpenHabitModalList, 
		setOpenHabitModalList, 
		habits, 
		setHabits,
		dateSort
	} = props

	const [habitListState, setHabitListState] = useState(habitList)
	const [oldHabitName, setOldHabitName] = useState(habitList.habit_name)

	async function updateHabit() {
		const newHabitListState = {...habitListState}
		delete newHabitListState.habits
		const res = await updateUserHabitList(newHabitListState)
	
		if (res.status === 200) {
			let newHabitList = [...habits]
			for (let i = 0; i < newHabitList.length; i++) {
				if (newHabitList[i].habit_name === oldHabitName) {
					newHabitList[i].habit_name = res.data.habit_name
					newHabitList[i].color = res.data.color
					newHabitList[i].default_repeat_count = res.data.default_repeat_count
					newHabitList[i].icon_url = res.data.icon_url
					for( let j = 0; j < newHabitList[i].habits.length; j++) {
						if (newHabitList[i].habits[j].habit_name === oldHabitName) {

							newHabitList[i].habits[j].target_repeat_count = res.data.default_repeat_count 
							newHabitList[i].habits[j].habit_name = res.data.habit_name 
						}
					}
				}
			}
			setHabits(newHabitList)
		}

	}
	async function createHabit() {
		const res = await createUserHabitList(habitListState)
		
		if (res.status === 200) {
			let newHabits = []
			if (dateSort === DATE_CHOICES.weekly) {newHabits = getDateWeekly()}
			if (dateSort === DATE_CHOICES.biweekly) {newHabits = getDateBiWeekly()}
			if (dateSort === DATE_CHOICES.monthly) {newHabits = getDateMonthly()}
			const newHabitListState = {...habitListState, habits: [...newHabits]}
			for (let i = 0; i < newHabitListState.habits.length; i++) {
				newHabitListState.habits[i].target_repeat_count = res.data.default_repeat_count 
				newHabitListState.habits[i].habit_name = res.data.habit_name
			}
			setHabits([...habits ,newHabitListState])
		}
	}

	useEffect(() => {
		setHabitListState(habitList)
		setOldHabitName(habitList.habit_name)
	}, [habitList])

	return(
		<div className={css.wrapper} style={{display: isOpenHabitModalList ? "flex" : "none"}}>
			<div className={css.container}>
				{isOpenHabitModalList && 
					<>
						<div className={css.headerWrapper}>
							<div className={css.titleContainer}>
								<h1>{!habitList.id ? "Create Habit" : "Edit Habit"}</h1>
							</div>
							<Button onClick={ ()=> {setOpenHabitModalList(false); setHabitListState(DEFAULT_HABIT_LIST);} }>
								<Close/>
							</Button>
							
						</div>
						<TextInput
							value={habitListState.habit_name || ""}
							setValue={ (value) => { setHabitListState({...habitListState, habit_name: value}) } }
							placeholder="Habit name"
						/>
						<NumberPicker
							id={css.requiredCountContainer}
							value={habitListState.default_repeat_count || 0}
							setState={ (value) => { setHabitListState({...habitListState, default_repeat_count: value}) } } 
						>
							Default Target Repeat Count
						</NumberPicker>
						<ColorPicker
							value={habitListState.color || ""}
							setValue={(value) => { setHabitListState({...habitListState, color: value}) }}
						> 
						Default Color
						</ColorPicker> 
							
						
						<div className={css.rowContainer}>
							<Button 
								primary={false} 
								onClick={() => {
									if (habitListState.id) {updateHabit()}
									if (!habitListState.id) {createHabit()}
								}}
							>
								Save
							</Button>
						</div>
					</>
				}
			</div>
		</div>
	)
}