import css from './dateChanger.module.css'
import { SwitchBtn, SmallBtn, Button } from "../../common";
import { ArrowLeft, ArrowRight, Plus } from '../../../public/svgs';
import { useState } from 'react';

export default function DateChanger () {

	const [dateSort, setDateSort] = useState()

	return (
		<div className={css.wrapper}>
			<div>
				<div className={css.arrowBtnContainer}>
					<SmallBtn>
						<ArrowLeft/>
					</SmallBtn>
					<SmallBtn>
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