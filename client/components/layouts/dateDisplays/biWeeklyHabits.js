import { motion } from "framer-motion"
import moment from 'moment'

import { DEFAULT_HABIT_LIST, WEEKDAYS } from '../../../common/constants'
import { calcBgColor } from '../../../common/utils'
import { HabitDay } from '../../common'
import css from './biWeeklyHabits.module.css'

export default function BiWeeklyHabits(props) {
	const { 
		habits, 
		setCurrentHabit, 
		setIsHabitModalOpen,
		setIsHabitModalListOpen,
		setCurrentHabitList,
		dateSort
	} = props
	const biWeeklyDays = [...WEEKDAYS, ...WEEKDAYS]

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

	return (
		<div className={css.pageWrapper}>
			{!habits &&
				<h1>No habits created yet</h1>
			}
			{ habits && 
				<>
					<div className={css.rowWrapper}>
						<div></div>
						<div className={css.contentContainer} id={css.mainWeekTitle}>
							{biWeeklyDays.map((day, i) => (
								<h6 key={i}>{day}</h6>
							))}
						</div>
					</div>

					{habits && habits.map((habit, i) => (
						<motion.div
							className={css.rowWrapper} 
							key={i} 
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
									onClick={() => {updateCurrentHabitList({habit:habit})} }
									style={{backgroundColor: `rgb(${habit.color})` || `rgb(${DEFAULT_HABIT_LIST.color})` }} 
								>
									<h2>
										{habit.habit_name}
									</h2>
								</motion.div>
							</motion.div>
							<motion.div className={css.contentContainer}>
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
											{WEEKDAYS[(moment(habitDay.date_created).day())]}
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