
import { useEffect, useState } from 'react';
import { useRouter } from 'next/router'

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
	decrementWeekly,
	getTodaysHabits,
	localLogout,
	isLoggedIn
} from '../common/utils';
import { 
	Navbar,
	DateChanger, 
	DateHabits, 
	HabitModal,
	HabitModalList,
	Footer
} from "../components/layouts";
import { NotifModal } from '../components/common';
import LoaderPage from '../components/layouts/loaderPage/loaderPage';
import { AnimatePresence } from 'framer-motion';
import HabitsToday from '../components/layouts/habitsToday';


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
	const [habitsToday, setHabitsToday] = useState([])

	const router = useRouter()

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
			if (err.response.status === 401) {
				localLogout()
				return router.push('/auth')
			}
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
			return setIsLoading(false)
		} catch(err) {
			if (err.response.status === 401) {
				localLogout()
				return router.push('/auth')
			}
			setNotifModalOpen(true)
			setNotifModalContent({msg: "An error occurred in fetching the data", error: true})	
			return setIsLoading(false)
		}
	}

	async function checkIfLoggedIn() {
		if (isLoggedIn() === false) {
			await localLogout()
			router.push('/auth')
		}
	}

	useEffect(() => {
		checkIfLoggedIn()
	},[])

	useEffect(() => {
		fetchData()
	},[dateSort])

	useEffect(() => {
		if (habits !== null) {
			//setting current habits
			const newHabitsToday = getTodaysHabits(habits)
			if (newHabitsToday.length < 1) {
				return
			}
			setHabitsToday(newHabitsToday)
		}
	},[habits])


	return(
		<div className={css.pageWrapper}>
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
				habitsToday={habitsToday}
				setHabitsToday={setHabitsToday}
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
						setIsLoading={setIsLoading}
						setCurrentHabitList={setCurrentHabitList}
						dateSort={dateSort}
						setDateSort={setDateSort}
						changeDate={changeDate}
						selectedDates={selectedDates}
						setIsHabitModalListOpen={setIsHabitModalListOpen}
					/>
					{isLoading ?
						<LoaderPage />
					:
						<AnimatePresence>
							<DateHabits 
								isLoading={isLoading}
								dateSort={dateSort}
								habits={habits}
								selectedDates={selectedDates}
								setCurrentHabit={setCurrentHabit}
								setIsHabitModalOpen={setIsHabitModalOpen}
								setIsHabitModalListOpen={setIsHabitModalListOpen}
								setCurrentHabitList={setCurrentHabitList}
								key={0}
							/>
							{(habits !== null || habits?.length > 0) &&
								<HabitsToday
									habitsToday={habitsToday} 
									setCurrentHabit={setCurrentHabit}
									setIsHabitModalOpen={setIsHabitModalOpen}
									key={1}
								/>
							}
							<Footer/>
						</AnimatePresence>
					}
				</div>
			</div>
		</div>
	)
}