import moment from 'moment'
import { useEffect, useState } from 'react'

import css from './habitModal.module.css'
import { Close } from '../../../public/svgs'
import { createUserHabit, updateUserHabit } from '../../../common/services'
import { Button, TextArea, NumberPicker } from '../../common'

export default function HabitModal(props) {
	const { habit, openHabitModal, setOpenHabitModal, habits, setHabits } = props

	const [habitState, setHabitState] = useState(habit)

	function markComplete() {
		setHabitState({...habitState, repeat_count: habitState.target_repeat_count})
	}
	async function updateHabit() {
		//! add error handling
		const res = await updateUserHabit(habitState)
		if (res.status === 200) {
			let newHabitList = [...habits]
			for (let i = 0; i < newHabitList.length; i++) {
				if (newHabitList[i].habit_name === habitState.habit_name) {
					for( let j = 0; j < newHabitList[i].habits.length; j++) {
						if (
							moment(newHabitList[i].habits[j].date_created).format("YYYY-MM-DD") === 
							moment(habitState.date_created).format("YYYY-MM-DD")
						) {
							newHabitList[i].habits[j] = {...res.data, habit_id: res.data.id}
						}
					}
				}
			}
			setHabits(newHabitList)
		}
	}
	async function createHabit() {
		//! add error handling
		const res = await createUserHabit(habitState)
		if (res.status === 200) {
			let newHabitList = [...habits]
			for (let i = 0; i < newHabitList.length; i++) {
				if (newHabitList[i].habit_name === habitState.habit_name) {
					for( let j = 0; j < newHabitList[i].habits.length; j++) {
						if (
							moment(newHabitList[i].habits[j].date_created).format("YYYY-MM-DD") === 
							moment(habitState.date_created).format("YYYY-MM-DD")
						) {
							newHabitList[i].habits[j] = {...res.data, habit_id: res.data.id}
						}
					}
				}
			}
			setHabits(newHabitList)
		}
	}

	useEffect(() => {
		console.log(habit)
		setHabitState(habit)
	}, [habit])

	return(
		<div className={css.wrapper} style={{display: openHabitModal ? "flex" : "none"}}>
			<div className={css.container}>
				{habitState && 
					<>
						<div className={css.headerWrapper}>
							<div className={css.headerContainer}>
								<div className={css.iconContainer} style={{backgroundColor: habitState.color || "#62A1FF"}}>
								</div>
								<div className={css.titleContainer}>
									<h1>{moment(habitState.date_created).format("MMMM DD, YYYY")}</h1>
									<span>{habitState.habit_name}</span>
								</div>
							</div>
							<Button onClick={ ()=> {setOpenHabitModal(false); setHabitState(null)} }>
								<Close/>
							</Button>
							
						</div>

						<NumberPicker 
							setState={ (value) => { setHabitState({...habitState, repeat_count: value}) } } 
							value={habitState.repeat_count} 
							maxValue={habitState.target_repeat_count || habitState.default_repeat_count}
						>
							This day's Repeat Count
						</NumberPicker>
						<NumberPicker
							id={css.requiredCountContainer}
							value={habitState.target_repeat_count}
							setState={ (value) => { setHabitState({...habitState, target_repeat_count: value}) } } 
						>
							Required Target Repeat Count
						</NumberPicker>
						<TextArea
							placeholder="Write a comment"
							value={habitState.comment}
							updateValue={(value) => { setHabitState({...habitState, comment: value})  }}
						/>
						<div className={css.rowContainer}>
							<Button 
								id={css.greenBtn} 
								primary={false} 
								onClick={markComplete}
							>
								Mark As Complete
							</Button>
							<Button 
								primary={false} 
								onClick={() => {
									if (habitState.habit_id) {updateHabit()}
									if (!habitState.habit_id) {createHabit()}
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