import { useEffect } from 'react'
import moment from 'moment'
import { WEEKDAYS_INITIALS } from '../../../common/constants'
import css from './monthlyHabits.module.css'

export default function MonthlyHabits(props) {
	const { habits, setCurrentHabit, setIsHabitModalOpen } = props

	function updateCurrentHabit({habit, habitDay}) {
		const newCurrentHabit = {
			...habitDay,
			...habit,
		}
		delete newCurrentHabit.id
		delete newCurrentHabit.habits
		setCurrentHabit(newCurrentHabit)
		setIsHabitModalOpen(true)
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
						{habits[0].habits.map((habitDay, i) => (
							<h6 key={i}>{WEEKDAYS_INITIALS[(moment(habitDay.date_created).day())]}</h6>
						))}
					</div>
				</div>

				{habits.map((habit, i) => (
					<div className={css.rowWrapper} key={i}>
						<div className={css.rowTitle}>
							<div style={{backgroundColor: habit.color || '#62A1FF'}} className={css.iconContainer}>
							</div>
							<h2>
								{habit.habit_name}
							</h2>
						</div>
						<div className={css.contentContainer}>
							{habit.habits.map((habitDay, i) => (
								<div className={css.dayWrapper} key={i}>
									<h6 className={css.dayTitle}>
										{WEEKDAYS_INITIALS[(moment(habitDay.date_created).day())]}
									</h6>
									<h6	className={css.dayContainer} onClick={() => { updateCurrentHabit({habit, habitDay}) }} >
										{moment(habitDay.date_created).format("DD")}
									</h6>
								</div>
							))}
						</div>
					</div>
				))}
			 </>
			}
		</div>
	)
}