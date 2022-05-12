import moment from 'moment'
import { DATE_CHOICES } from '../../../common/constants'
import { calcBgColor } from '../../../common/utils'
import css from './habitDay.module.css'

export default function HabitDay(props) {
	const {updateCurrentHabit, habit, habitDay, dateSort} = props

	const backgroundColor = calcBgColor(habitDay.repeat_count/habitDay.target_repeat_count, habit.color)
	const ratio = (habitDay.repeat_count/habitDay.target_repeat_count)

	return (
		<h6 
			className={dateSort === DATE_CHOICES.monthly ? css.monthlyDayContainer : css.dayContainer} 
			onClick={() => { updateCurrentHabit({habit, habitDay}) }} 
			style={{backgroundColor: backgroundColor, color: ratio >= 1 && "#fff" }}
		>
			{moment(habitDay.date_created).format("D")}
		</h6>
	)
}