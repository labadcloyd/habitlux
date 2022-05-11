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
						<div className={css.rowWrapper} key={i}>
							<div className={css.rowTitleContainer}>
								<div className={css.rowTitle} onClick={() => {updateCurrentHabitList({habit:habit})} }>
									<div 
										style={{backgroundColor: `rgb(${habit.color})` || `rgb(${DEFAULT_HABIT_LIST.color})` }} 
										className={css.iconContainer}>
									</div>
									<h2>
										{habit.habit_name}
									</h2>
								</div>
							</div>
							<div className={css.contentContainer}>
								{habit.habits.map((habitDay, i) => (
									<div className={css.dayWrapper} key={i}>
										<h6 className={css.dayTitle}>
											{WEEKDAYS[(moment(habitDay.date_created).day())]}
										</h6>
										<HabitDay
											dateSort={dateSort}
											updateCurrentHabit={updateCurrentHabit}
											habit={habit}
											habitDay={habitDay}
										/>
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