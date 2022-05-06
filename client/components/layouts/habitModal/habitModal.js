import moment from 'moment'
import { useEffect, useState } from 'react'
import { Close } from '../../../public/svgs'
import { Button, TextArea } from '../../common'
import NumberPicker from '../numberPicker'
import css from './habitModal.module.css'

export default function HabitModal(props) {
	const { habit, openHabitModal, setOpenHabitModal } = props

	const [habitState, setHabitState] = useState(habit)

	function markComplete() {
		setHabitState({...habitState, repeat_count: habitState.target_repeat_count})
	}

	useEffect(() => {
		setHabitState(habit)
	}, [habit])

	return(
		<div className={css.wrapper} style={{display: openHabitModal ? "flex" : "none"}}>
			<div className={css.container}>
				{habitState && 
					<>
						<div className={css.headerWrapper}>
							<div className={css.headerContainer}>
								<div className={css.iconContainer} style={{backgroundColor: habitState.color || "#62A1FF"}}>
								</div>
								<div className={css.titleContainer}>
									<h1>{moment(habitState.date_created).format("MMMM DD, YYYY")}</h1>
									<span>{habitState.habit_name}</span>
								</div>
							</div>
							<Button onClick={ ()=> {setOpenHabitModal(false); setHabitState(null)} }>
								<Close/>
							</Button>
							
						</div>

						<NumberPicker 
							setState={ (value) => { setHabitState({...habitState, repeat_count: value}) } } 
							value={habitState.repeat_count} 
							maxValue={habitState.target_repeat_count || habitState.default_repeat_count}
						>
							This day's Repeat Count
						</NumberPicker>
						<NumberPicker
							id={css.requiredCountContainer}
							value={habitState.target_repeat_count || habitState.default_repeat_count}
							setState={ (value) => { setHabitState({...habitState, target_repeat_count: value}) } } 
						>
							Required Target Repeat Count
						</NumberPicker>
						<TextArea placeholder="Write a comment"/>
						<div className={css.rowContainer}>
							<Button id={css.greenBtn} primary={false} onClick={markComplete}>Mark As Complete</Button>
							<Button primary={false}>Save</Button>
						</div>
					</>
				}
			</div>
		</div>
	)
}