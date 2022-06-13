import moment from 'moment'
import { useEffect, useState } from 'react'
import { AnimatePresence, motion } from 'framer-motion'

import css from './habitModal.module.css'
import { Close } from '../../../public/svgs'
import { createUserHabit, updateUserHabit } from '../../../common/services'
import { Button, TextArea, NumberPicker } from '../../common'
import { DEFAULT_HABIT_LIST } from '../../../common/constants'

export default function HabitModal(props) {
	const {
		habit,
		openHabitModal,
		setOpenHabitModal,
		habits,
		setHabits,
		setNotifModalOpen,
		setNotifModalContent
	} = props

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
				if (newHabitList[i].habit_list_id === habitState.habit_list_id) {
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
			setNotifModalOpen(true)
			setNotifModalContent({msg: "Successfully udpated habit", error: false})	
			setOpenHabitModal(false)
		} else {
			setNotifModalOpen(true)
			setNotifModalContent({msg: "An error occurred in updating habit", error: true})	
		}
	}
	async function createHabit() {
		//! add error handling
		//! BUG: Updates the same habit with the same name
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
			setNotifModalOpen(true)
			setNotifModalContent({msg: "Successfully created habit", error: false})	
			setOpenHabitModal(false)
		} else {
			setNotifModalOpen(true)
			setNotifModalContent({msg: "An error occurred in creating habit", error: true})	
		}
	}

	useEffect(() => {
		setHabitState(habit)
	}, [habit])

	return(
		<AnimatePresence>
		{openHabitModal &&
			<motion.div
				onClick={() => (setOpenHabitModal(false))}

				className={css.wrapper}
			> 
				<motion.div 
					initial={{opacity: 1, scale: 0}}
					animate={{opacity: 1, scale: 1}}
					transition={{ type: 'spring', duration: 0.3 }}
					exit={{opacity: 1, scale: 0}}
					layout
					onClick={(e) => e.stopPropagation()}
					className={css.container}
				>
					{habitState && 
						<>
							<div className={css.headerWrapper}>
								<div className={css.headerContainer}>
									<div className={css.iconContainer} style={{backgroundColor: `rgb(${habitState.color})` || `rgb(${DEFAULT_HABIT_LIST.color})`}}>
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
				</motion.div>
			</motion.div>
		}
		</AnimatePresence>
	)
}