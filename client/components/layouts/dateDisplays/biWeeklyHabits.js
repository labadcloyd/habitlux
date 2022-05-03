import moment from 'moment'
import { WEEKDAYS } from '../../../common/constants'
import css from './biWeeklyHabits.module.css'

export default function BiWeeklyHabits(props) {
	const { habits } = props

	const biWeeklyDays = [...WEEKDAYS, ...WEEKDAYS]

	return (
		<div className={css.pageWrapper}>
			<div className={css.rowWrapper}>
				<div></div>
				<div className={css.contentContainer}>
					{biWeeklyDays.map((day, i) => (
						<h6 key={i}>{day}</h6>
					))}
				</div>
			</div>

			{habits && habits.map((habit, i) => (
				<div className={css.rowWrapper} key={i}>
					<div className={css.rowTitle}> {habit.name} </div>
					<div className={css.contentContainer}>
						{habit.habits.map((habitDay, i) => (
							<h6 key={i} className={css.dayContainer}>{moment(habitDay.date).format("DD")}</h6>
						))}
					</div>
				</div>
			))}
		</div>
	)
}