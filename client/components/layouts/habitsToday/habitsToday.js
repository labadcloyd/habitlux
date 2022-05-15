import { motion } from "framer-motion"
import css from './habitsToday.module.css'

export default function HabitsToday(props) {
	const { habitsToday, setIsHabitModalOpen, setCurrentHabit } = props

	return (
		<div className={css.wrapper}>
			<h1>Habits for today</h1>
			<div className={css.habitWrapper}>
				{habitsToday.map((habit, i) => (
					<motion.div
						initial={{opacity: 0}}
						animate={{opacity: 1}}
						exit={{opacity: 0}}
						layout

						whileHover={{ scale:1.05, boxShadow: '0px 0px 15px rgba(0, 0, 0, 0.2)' }}
						whileTap={{ scale:0.95, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
						transition={{ duration: 0.2, delay: 0 }}

						key={i} 
						className={css.habitContainer}
						onClick={() => {setCurrentHabit(habit); setIsHabitModalOpen(true)}}
					>
						<div className={css.iconContainer} style={{backgroundColor: `rgb(${habit.color})`}} />
						<div className={css.habitTitle}>
							<h2>{habit.habit_name}</h2>
							{(habit.repeat_count < 1 || (habit.repeat_count/habit.target_repeat_count !== 1)) ?
								<span
									className={css.notCompletedLabel}
									style={{ color: '#a7a7a7' }}
								>
									Not Completed
								</span>
							:
								<span
									className={css.completedLabel}
									style={{ color: '#60f95d' }}
								>
									Completed
								</span>
							}
						</div>
						<div></div>
					</motion.div>
				))}
			</div>
		</div>
	)
}