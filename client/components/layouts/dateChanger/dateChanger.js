import css from './dateChanger.module.css'
import { SwitchBtn, SmallBtn, Button } from "../../common";
import { ArrowLeft, ArrowRight, Plus } from '../../../public/svgs';
import { useEffect, useState } from 'react';
import moment from 'moment';

export default function DateChanger (props) {
	const {
		setDateSort,
		changeDate,
		selectedDates,
		setIsHabitModalListOpen
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
	}, [selectedDates])


	return (
		<div className={css.wrapper}>
			<div>
				<div className={css.arrowBtnContainer}>
					<SmallBtn onClick={() => {changeDate({increment: false})}}>
						<ArrowLeft/>
					</SmallBtn>
					<SmallBtn onClick={() => {changeDate({increment: true})}}>
						<ArrowRight/>
					</SmallBtn>
				</div>
				<div className={css.dateContainer}>
					{dateTitle}
				</div>
				<SwitchBtn values={["Bi-Weekly", "Monthly"]} setValue={setDateSort}/>
			</div>
			<Button onClick={() => {setIsHabitModalListOpen(true)}}>
				<Plus/>
				Add habit
			</Button>
		</div>
	)
}