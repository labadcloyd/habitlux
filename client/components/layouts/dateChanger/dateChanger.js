import css from './dateChanger.module.css'
import { SwitchBtn, SmallBtn, Button } from "../../common";
import { ArrowLeft, ArrowRight, Plus } from '../../../public/svgs';
import { useEffect, useState } from 'react';
import moment from 'moment';
import { DATE_CHOICES, DEFAULT_HABIT_LIST } from '../../../common/constants';

export default function DateChanger (props) {
	const {
		dateSort,
		setDateSort,
		changeDate,
		selectedDates,
		setIsHabitModalListOpen,
		setCurrentHabitList
	} = props

	const [dateTitle, setDateTitle] = useState("")

	useEffect(() => {
		const start_date = moment(selectedDates[0]) 
		const end_date = moment(selectedDates[selectedDates.length -1]) 

		if( start_date.year() !== end_date.year() ) {
			return setDateTitle(`${start_date.format('MMMM D, YYYY')} - ${end_date.format('MMMM D, YYYY')}`)
		} else if (start_date.year() === end_date.year()) {
			if (start_date.month() === end_date.month()) {
				return setDateTitle(`${start_date.format('MMMM D')} - ${end_date.format('D, YYYY')}`)
			}
			return setDateTitle(`${start_date.format('MMMM D')} - ${end_date.format('MMMM D, YYYY')}`)
		}
	}, [selectedDates, dateSort])


	return (
		<div className={css.wrapper}>
			<div className={css.container}>
				<div>
					<div className={css.arrowBtnContainer}>
						<SmallBtn onClick={() => {changeDate({increment: false})}}>
							<ArrowLeft/>
						</SmallBtn>
						<div className={css.dateContainerSmall}>
							{dateTitle}
						</div>
						<SmallBtn onClick={() => {changeDate({increment: true})}}>
							<ArrowRight/>
						</SmallBtn>
					</div>
					<div className={css.dateContainer}>
						{dateTitle}
					</div>
					{ dateSort !== DATE_CHOICES.weekly &&
						<SwitchBtn values={["Bi-Weekly", "Monthly"]} setValue={setDateSort}/>
					}
				</div>
				<Button onClick={() => {setIsHabitModalListOpen(true); setCurrentHabitList(DEFAULT_HABIT_LIST);}}>
					<Plus/>
					Add habit
				</Button>
			</div>
		</div>
	)
}