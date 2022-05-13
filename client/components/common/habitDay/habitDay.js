import { motion } from "framer-motion"
import moment from 'moment'

import { DATE_CHOICES } from '../../../common/constants'
import { calcBgColor } from '../../../common/utils'
import css from './habitDay.module.css'

export default function HabitDay(props) {
	const {updateCurrentHabit, habit, habitDay, dateSort} = props

	const backgroundColor = calcBgColor(habitDay.repeat_count/habitDay.target_repeat_count, habit.color)
	const ratio = (habitDay.repeat_count/habitDay.target_repeat_count)

	return (
		<motion.h6 
			whileHover={{ scale:1.1, boxShadow: '0px 0px 15px rgba(0, 0, 0, 0.2)' }}
			whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
			transition={{ duration: 0.2, delay: 0 }}

			className={dateSort === DATE_CHOICES.monthly ? css.monthlyDayContainer : css.dayContainer} 
			onClick={() => { updateCurrentHabit({habit, habitDay}) }} 
			style={{backgroundColor: backgroundColor, color: ratio >= 1 && "#fff" }}
		>
			{moment(habitDay.date_created).format("D")}
		</motion.h6>
	)
}