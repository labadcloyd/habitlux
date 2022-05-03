
import css from '../styles/dashboard.module.css'
import { Navbar, DateChanger, BiWeeklyHabits } from "../components/layouts";
import { useEffect, useState } from 'react';
import { getDateBiWeekly } from '../common/utils';

export default function Dashboard() {
	const [habits, setHabits] = useState([])

	useEffect(() => {
		const newHabit = {
			name: 'testing',
			habits: getDateBiWeekly()
		}
		
		setHabits([newHabit])

		console.log(newHabit)
	},[])
	return(
		<div className={css.pageWrapper}>
			<Navbar/>
			<div className={css.pageContainer}>
				<div className={css.contentContainer}>
					<DateChanger/>
					<BiWeeklyHabits habits={habits} />
				</div>
			</div>
		</div>
	)
}