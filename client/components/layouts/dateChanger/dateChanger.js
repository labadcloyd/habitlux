import css from './dateChanger.module.css'
import { SwitchBtn, SmallBtn, Button } from "../../common";
import { ArrowLeft, ArrowRight, Plus } from '../../../public/svgs';
import { useState } from 'react';

export default function DateChanger (props) {
	const {
		setIsLoading,
		dateSort,
		setDateSort,
		isLoading,
		incrementDate,
		decrementDate
	} = props



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
					March 2022
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