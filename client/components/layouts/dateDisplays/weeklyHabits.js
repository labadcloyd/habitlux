import moment from 'moment'
import { WEEKDAYS } from '../../../common/constants'
import css from './weeklyHabits.module.css'

export default function WeeklyHabits(props) {
	const { habits } = props

	return (
		<div className={css.pageWrapper}>
			<div className={css.rowWrapper}>
				<div></div>
				<div className={css.contentContainer} id={css.mainWeekTitle}>
					{WEEKDAYS.map((day, i) => (
						<h6 key={i}>{day}</h6>
					))}
				</div>
			</div>

			{habits && habits.map((habit, i) => (
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
								<h6 className={css.dayContainer}>{moment(habitDay.date_created).format("DD")}</h6>
							</div>
						))}
					</div>
				</div>
			))}
		</div>
	)
}