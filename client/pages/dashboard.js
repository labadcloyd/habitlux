
import { useEffect, useState } from 'react';

import css from '../styles/dashboard.module.css'
import { getAllUserHabits } from '../common/services/habit';
import { DATE_CHOICES } from '../common/constants';

import { 
	addHabitsToDate, 
	getDateBiWeekly, 
	getDateMonthly, 
	getDateWeekly, 
	incrementBiWeekly, 
	incrementMonthly, 
	incrementWeekly, 
	decrementBiWeekly, 
	decrementMonthly, 
	decrementWeekly
} from '../common/utils';
import { 
	Navbar,
	DateChanger, 
	MonthlyHabits, 
	BiWeeklyHabits, 
	WeeklyHabits
} from "../components/layouts";


export default function Dashboard() {
	const [habits, setHabits] = useState(null)
	const [dateSort, setDateSort] = useState(DATE_CHOICES.biweekly)
	const [isLoading, setIsLoading] = useState(true)
	const [selectedDates, setSelectedDates] = useState([])

	async function fetchData() {
		setIsLoading(true)
		let selectedDatesWithHabits = [], newSelectedDates = []
		if (dateSort === DATE_CHOICES.biweekly) {selectedDatesWithHabits = getDateBiWeekly(selectedDates[0])}
		if (dateSort === DATE_CHOICES.weekly) {selectedDatesWithHabits = getDateWeekly(selectedDates[0])}
		if (dateSort === DATE_CHOICES.monthly) {selectedDatesWithHabits = getDateMonthly(selectedDates[0])}
		selectedDatesWithHabits.forEach((dateWithHabit) => {
			newSelectedDates.push(dateWithHabit.date_created)
		})

		setSelectedDates(newSelectedDates)
	
		const res = await getAllUserHabits({ 
			Start_Date: newSelectedDates[0],
			End_Date: newSelectedDates[newSelectedDates.length - 1]
		})
		if (res.data.length < 1) {
			const formattedHabits = addHabitsToDate({habits: res.data, selectedDatesWithHabits: selectedDatesWithHabits })
			setHabits(formattedHabits)
			return setIsLoading(false)
		}
		const formattedHabits = addHabitsToDate({habits: res.data, selectedDatesWithHabits: selectedDatesWithHabits })

		setHabits(formattedHabits)
		setIsLoading(false)
	}

	async function changeDate({increment}) {
		setIsLoading(true)
		let selectedDatesWithHabits = [], newSelectedDates = []
		if (increment) {
			if (dateSort === DATE_CHOICES.biweekly) {selectedDatesWithHabits = incrementBiWeekly(selectedDates[selectedDates.length -1])}
			if (dateSort === DATE_CHOICES.weekly) {selectedDatesWithHabits = incrementWeekly(selectedDates[selectedDates.length -1])}
			if (dateSort === DATE_CHOICES.monthly) {selectedDatesWithHabits = incrementMonthly(selectedDates[selectedDates.length -1])}
		} else if (!increment) {
			if (dateSort === DATE_CHOICES.biweekly) {selectedDatesWithHabits = decrementBiWeekly(selectedDates[0])}
			if (dateSort === DATE_CHOICES.weekly) {selectedDatesWithHabits = decrementWeekly(selectedDates[0])}
			if (dateSort === DATE_CHOICES.monthly) {selectedDatesWithHabits = decrementMonthly(selectedDates[0])}
		}

		selectedDatesWithHabits.forEach((dateWithHabit) => {
			newSelectedDates.push(dateWithHabit.date_created)
		})

		setSelectedDates(newSelectedDates)
	
		const res = await getAllUserHabits({ 
			Start_Date: newSelectedDates[0],
			End_Date: newSelectedDates[newSelectedDates.length - 1]
		})
		if (res.data.length < 1) {
			setHabits(selectedDatesWithHabits)
			return setIsLoading(false)
		}
		const formattedHabits = addHabitsToDate({habits: res.data, selectedDatesWithHabits })
		setHabits(formattedHabits)
		setIsLoading(false)
	}

	useEffect(() => {
		fetchData()
	},[dateSort])

	return(
		<div className={css.pageWrapper}>
			<Navbar/>
			<div className={css.pageContainer}>
				<div className={css.contentContainer}>
					<DateChanger
						dateSort={dateSort}
						setDateSort={setDateSort}
						isLoading={isLoading}
						setIsLoading={setIsLoading}
						changeDate={changeDate}
						selectedDates={selectedDates}
					/>
					{isLoading ?
						<h1>Loading...</h1>
					:
						<>
							{dateSort === DATE_CHOICES.monthly &&
								<MonthlyHabits habits={habits} />
							}
							{dateSort === DATE_CHOICES.biweekly &&
								<BiWeeklyHabits habits={habits} />
							}
							{dateSort === DATE_CHOICES.weekly &&
								<WeeklyHabits habits={habits} />
							}
						</>
					}
				</div>
			</div>
		</div>
	)
}