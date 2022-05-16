import { AnimatePresence, motion } from "framer-motion"
import { useEffect, useState } from 'react'

import { DATE_CHOICES, DEFAULT_HABIT_LIST } from '../../../common/constants'
import { createUserHabitList, deleteUserHabitList, updateUserHabitList } from '../../../common/services'
import { getDateBiWeekly, getDateMonthly, getDateWeekly } from '../../../common/utils'
import { Button, TextInput, ColorPicker, NumberPicker, ConfirmModal } from '../../common'
import { Close } from '../../../public/svgs'
import css from './habitModalList.module.css'

export default function HabitModalList(props) {
	const { 
		habitList, 
		isOpenHabitModalList, 
		setOpenHabitModalList, 
		habits, 
		setHabits,
		dateSort,
		selectedDates,
		setNotifModalOpen,
		setNotifModalContent,
		habitsToday,
		setHabitsToday
	} = props

	const [habitListState, setHabitListState] = useState(habitList)
	const [oldHabitName, setOldHabitName] = useState(habitList.habit_name)

	const [showConfirmModal, setShowConfirmModal] = useState(false)

	async function updateHabit() {
		try {
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
				setNotifModalContent({msg: "Successfully udpated habit list", error: false})	
				setNotifModalOpen(true)
	
				setOpenHabitModalList(false)
			}
		} catch (err) {
			setNotifModalContent({msg: "An error occurred in updating habit list", error: true})	
			setNotifModalOpen(true)
			setOpenHabitModalList(false)
		}
	}

	async function createHabit() {
		try {
			const res = await createUserHabitList(habitListState)
			if (res.status === 200) {
				let newHabits = []
				if (dateSort === DATE_CHOICES.weekly) {newHabits = getDateWeekly(selectedDates[0])}
				if (dateSort === DATE_CHOICES.biweekly) {newHabits = getDateBiWeekly(selectedDates[0])}
				if (dateSort === DATE_CHOICES.monthly) {newHabits = getDateMonthly(selectedDates[0])}
				const newHabitListState = {...habitListState, habits: [...newHabits]}
				for (let i = 0; i < newHabitListState.habits.length; i++) {
					newHabitListState.habits[i].target_repeat_count = res.data.default_repeat_count 
					newHabitListState.habits[i].habit_name = res.data.habit_name
					newHabitListState.id = res.data.id
				}
				if (habits === null) {
					setHabits([newHabitListState])
				} else {
					setHabits([...habits, newHabitListState])
				}
				setOpenHabitModalList(false)
				setNotifModalContent({msg: "Successfully created habit list", error: false})	
				setNotifModalOpen(true)
			}
		} catch (err) {
			setOpenHabitModalList(false)
			setNotifModalContent({msg: "An error occurred in creating habit list", error: true})	
			setNotifModalOpen(true)
		}
	}
	
	async function deleteHabit() {
		try {
			const res = await deleteUserHabitList({habit_name: habitListState.habit_name})
			if (res.status === 200) {
				let newHabits = []
				for (let i = 0; i < habits.length; i++) {
					if (habits[i].habit_name !== habitListState.habit_name) {
						newHabits.push(habits[i])
					}
				}
				let newHabitsToday = []
				for (let i = 0; i < habitsToday.length; i++) {
					if (habitsToday[i].habit_name !== habitListState.habit_name) {
						newHabitsToday.push(habitsToday[i])
					}
				}
				setHabitsToday(newHabitsToday)
				
				setHabits(newHabits)
				setNotifModalOpen(true)
				setNotifModalContent({msg: "Successfully deleted habit list", error: false})	
				setOpenHabitModalList(false)
				setShowConfirmModal(false)
			}
		} catch (err) {
			setNotifModalOpen(true)
			setNotifModalContent({msg: "An error occurred in deleting habit list", error: true})	
			setShowConfirmModal(false)
			setOpenHabitModalList(false)
		}
	}

	useEffect(() => {
		setHabitListState(habitList)
		setOldHabitName(habitList.habit_name)
	}, [habitList])

	return(
		<AnimatePresence>
			{showConfirmModal &&
				<ConfirmModal
					message={"Are you sure you want to delete this habit?"}
					showConfirmModal={setShowConfirmModal}
					action={deleteHabit}
					key={0}
				/>
			}
			{isOpenHabitModalList && 
			<motion.div
				onClick={() => (setOpenHabitModalList(false))}

				key={1}
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
						<>
							<motion.div className={css.headerWrapper}>
								<motion.div className={css.titleContainer}>
									<motion.h1>{!habitList.id ? "Create Habit" : "Edit Habit"}</motion.h1>
								</motion.div>
								<Button onClick={ ()=> {setOpenHabitModalList(false); setHabitListState(DEFAULT_HABIT_LIST);} }>
									<Close/>
								</Button>
								
							</motion.div>
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
								
							
							<motion.div className={css.rowContainer}>
								{habitList.id &&
									<Button 
										primary={false} 
										onClick={() => {setShowConfirmModal(true)} }
										style={{ backgroundColor: '#ff6262', color: '#fff' }}
									>
										Delete
									</Button>
								}
								<Button 
									primary={false} 
									onClick={() => {
										if (habitListState.id) {updateHabit()}
										if (!habitListState.id) {createHabit()}
									}}
								>
									Save
								</Button>
							</motion.div>
						</>
				</motion.div>
			</motion.div>
			} 
		</AnimatePresence>
	)
}