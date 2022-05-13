
import { useEffect, useRef, useState } from 'react';

import css from '../styles/dashboard.module.css'
import { getAllUserHabits } from '../common/services/habit';
import { DATE_CHOICES, DEFAULT_HABIT_LIST } from '../common/constants';

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
	DateHabits, 
	HabitModal,
	HabitModalList
} from "../components/layouts";
import { NotifModal } from '../components/common';


export default function Dashboard() {
	const [isLoading, setIsLoading] = useState(true)
	const [isHabitModalOpen, setIsHabitModalOpen] = useState(false)
	const [isHabitModalListOpen, setIsHabitModalListOpen] = useState(false)

	const [isNotifModalOpen, setNotifModalOpen] = useState(false)
	const [notifModalContent, setNotifModalContent] = useState({msg: "", error: false})

	const [dateSort, setDateSort] = useState(DATE_CHOICES.biweekly)
	const [habits, setHabits] = useState(null)
	const [selectedDates, setSelectedDates] = useState([])
	const [currentHabit, setCurrentHabit] = useState()
	const [currentHabitList, setCurrentHabitList] = useState(DEFAULT_HABIT_LIST)

	const pageWidthRef = useRef(null)

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
	
		try {
			const res = await getAllUserHabits({ 
				Start_Date: newSelectedDates[0],
				End_Date: newSelectedDates[newSelectedDates.length - 1]
			})
			if (res.data.length < 1) {
				setHabits(null)
				return setIsLoading(false)
			}
			const formattedHabits = addHabitsToDate({habits: res.data, datesWithHabits: [...selectedDatesWithHabits] })
			setHabits(formattedHabits)
			setIsLoading(false)
		} catch(err) {
			setNotifModalOpen(true)
			setNotifModalContent({msg: "An error occurred in fetching the data", error: true})	
			return setIsLoading(false)
		}
		
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
			setHabits(null)
			return setIsLoading(false)
		}
		const formattedHabits = addHabitsToDate({habits: res.data, datesWithHabits: [...selectedDatesWithHabits] })
		setHabits(formattedHabits)
		setIsLoading(false)
	}

	function handleWidthChange() {
		if (pageWidthRef.current.offsetWidth < 700) {
			if (dateSort !== DATE_CHOICES.biweekly) {
				setDateSort(DATE_CHOICES.biweekly)
			}
		}
	}

	useEffect(() => {
		fetchData()
	},[dateSort])

	useEffect(() => {
    if(pageWidthRef.current) {
			handleWidthChange()
			window.addEventListener('resize', handleWidthChange )
		}
    return ()=> { window.removeEventListener('resize', handleWidthChange ) }
	},[])

	return(
		<div className={css.pageWrapper} ref={pageWidthRef}>
			<Navbar/>
			<HabitModalList
				habitList={currentHabitList}
				isOpenHabitModalList={isHabitModalListOpen}
				setOpenHabitModalList={setIsHabitModalListOpen}
				habits={habits}
				setHabits={setHabits}
				dateSort={dateSort}
				selectedDates={selectedDates}
				setNotifModalOpen={setNotifModalOpen}
				setNotifModalContent={setNotifModalContent}
			/>
			<HabitModal 
				habit={currentHabit}
				openHabitModal={isHabitModalOpen}
				setOpenHabitModal={setIsHabitModalOpen}
				habits={habits}
				setHabits={setHabits}
				setNotifModalOpen={setNotifModalOpen}
				setNotifModalContent={setNotifModalContent}
			/>
			<NotifModal
				modalOpen={isNotifModalOpen}
				setModalOpen={setNotifModalOpen}
				content={notifModalContent}
				setContent={setNotifModalContent}
			/>
			<div className={css.pageContainer}>
				<div className={css.contentContainer}>
					<DateChanger
						setCurrentHabitList={setCurrentHabitList}
						dateSort={dateSort}
						setDateSort={setDateSort}
						changeDate={changeDate}
						selectedDates={selectedDates}
						setIsHabitModalListOpen={setIsHabitModalListOpen}
					/>
					{isLoading ?
						<h1>Loading...</h1>
					:
						<DateHabits 
							dateSort={dateSort}
							habits={habits} 
							selectedDates={selectedDates}
							setCurrentHabit={setCurrentHabit} 
							setIsHabitModalOpen={setIsHabitModalOpen}
							setIsHabitModalListOpen={setIsHabitModalListOpen}
							setCurrentHabitList={setCurrentHabitList}
						/>
					}
				</div>
			</div>
		</div>
	)
}