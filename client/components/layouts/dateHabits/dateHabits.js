import { motion } from "framer-motion"
import moment from 'moment'
import { useEffect, useState } from "react"

import { DATE_CHOICES, DEFAULT_HABIT_LIST, WEEKDAYS, WEEKDAYS_INITIALS } from '../../../common/constants'
import { HabitDay } from '../../common'
import css from './dateHabits.module.css'

export default function DateHabits(props) {
	const { 
		habits, 
		setCurrentHabit, 
		setIsHabitModalOpen,
		setIsHabitModalListOpen,
		setCurrentHabitList,
		dateSort,
		selectedDates,
		isLoading
	} = props
	const [selectedDays, setSelectedDays] = useState([])

	function updateCurrentHabit({habit, habitDay}) {
		const newCurrentHabit = {
			...habitDay,
			...habit,
		}
		newCurrentHabit.target_repeat_count = (newCurrentHabit.target_repeat_count || newCurrentHabit.default_repeat_count)
		delete newCurrentHabit.id
		delete newCurrentHabit.habits
		setCurrentHabit(newCurrentHabit)
		setIsHabitModalOpen(true)
	}

	function updateCurrentHabitList({habit}) {
		const newCurrentHabitList = {
			...habit
		}
		setCurrentHabitList(newCurrentHabitList)
		setIsHabitModalListOpen(true)
	}

	useEffect(() => {
		let newSelectedDays = []
		selectedDates.forEach((day) => {
			if (dateSort === DATE_CHOICES.monthly) {
				newSelectedDays.push(WEEKDAYS_INITIALS[moment(day).day()])
			} else {
				newSelectedDays.push(WEEKDAYS[moment(day).day()])
			}
		})
		setSelectedDays(newSelectedDays)
	}, [])

	return (
		<div className={css.pageWrapper}>
			{!habits &&
				<h1>No habits created yet</h1>
			}
			{ (habits && isLoading === false) && 
				<>
					<div
						id={css.mainWeekTitle}
						className={ dateSort === DATE_CHOICES.monthly ? css.monthlyWrapper : css.rowWrapper }
					>
						<div></div>
						<div 
							className={ dateSort === DATE_CHOICES.monthly ? css.monthlyContainer : css.contentContainer }
						>
							{selectedDays && selectedDays.map((day, i) => (
								<h6 key={i}>{day}</h6>
							))}
						</div>
					</div>

					{habits && habits.map((habit, i) => (
						<motion.div
							key={i} 
							className={ dateSort === DATE_CHOICES.monthly ? css.monthlyWrapper : css.rowWrapper }
						>
							<motion.div className={css.rowTitleContainer}>
								<motion.div
									initial={{opacity: 0}}
									animate={{opacity: 1}}
									exit={{opacity: 0}}
									layout

									whileHover={{ scale:1.1, boxShadow: '0px 0px 15px rgba(0, 0, 0, 0.2)' }}
									whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
									transition={{ duration: 0.1, delay: 0 }}

									className={css.rowTitle} 
									id={dateSort === DATE_CHOICES.monthly ? css.monthlyTitle : css.weeklyTitle}
									onClick={() => {updateCurrentHabitList({habit:habit})} }
								>
									<div
										className={css.iconContainer}
										id={dateSort === DATE_CHOICES.monthly ? css.monthlyIcon : css.weeklyIcon}
										style={{backgroundColor: `rgb(${habit.color})` || `rgb(${DEFAULT_HABIT_LIST.color})` }} 
									/>
									<h2>
										{habit.habit_name}
									</h2>
								</motion.div>
							</motion.div>
							<motion.div
								className={ dateSort === DATE_CHOICES.monthly ? css.monthlyContainer : css.contentContainer }
							>
								{habit.habits.map((habitDay, i) => (
									<motion.div
										initial={{opacity: 0}}
										animate={{opacity: 1}}
										exit={{opacity: 0}}
										transition={{ duration: 0.1, delay: `${i*0.03}` }}
										layout

										className={css.dayWrapper} 
										key={i}
									>
										<h6 className={css.dayTitle}>
											{ dateSort === DATE_CHOICES.monthly?
												WEEKDAYS_INITIALS[(moment(habitDay.date_created).day())]
											:
												WEEKDAYS[(moment(habitDay.date_created).day())]
											}
										</h6>
										<HabitDay
											dateSort={dateSort}
											updateCurrentHabit={updateCurrentHabit}
											habit={habit}
											habitDay={habitDay}
										/>
									</motion.div>
								))}
							</motion.div>
						</motion.div>
					))}
				</>
			}
		</div>
	)
}