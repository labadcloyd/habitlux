import css from './dateChanger.module.css'
import { SwitchBtn, SmallBtn, Button } from "../../common";
import { ArrowLeft, ArrowRight, Plus } from '../../../public/svgs';
import { useEffect, useState } from 'react';
import moment from 'moment';

export default function DateChanger (props) {
	const {
		setIsLoading,
		dateSort,
		setDateSort,
		isLoading,
		incrementDate,
		decrementDate,
		selectedDates
	} = props

	const [dateTitle, setDateTitle] = useState("")

	useEffect(() => {
		const start_date = moment(selectedDates[0]) 
		const end_date = moment(selectedDates[selectedDates.length -1]) 

		if( start_date.year() !== end_date.year() ) {
			setDateTitle(`${start_date.format('MMMM d, YYYY')} - ${end_date.format('MMMM d, YYYY')}`)
		}
	}, [selectedDates])


	return (
		<div className={css.wrapper}>
			<div>
				<div className={css.arrowBtnContainer}>
					<SmallBtn onClick={decrementDate}>
						<ArrowLeft/>
					</SmallBtn>
					<SmallBtn onClick={incrementDate}>
						<ArrowRight/>
					</SmallBtn>
				</div>
				<div className={css.dateContainer}>
					{dateTitle}
				</div>
				<SwitchBtn values={["Bi-Weekly", "Monthly"]} setValue={setDateSort}/>
			</div>
			<Button>
				<Plus/>
				Add habit
			</Button>
		</div>
	)
}