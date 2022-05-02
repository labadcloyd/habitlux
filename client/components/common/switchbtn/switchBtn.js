import { useState } from 'react'
import css from './switchBtn.module.css'

export default function SwitchBtn(props) {
	const { values, setValue } = props

	const [currentValue, setCurrentValue] = useState(values[0] || "")

	function changeValue(clickedValue) {
		setValue(clickedValue)
		setCurrentValue(clickedValue)
	}

	return (
		<div>
			<div className={css.wrapper}>
				{values && values.map((value, i) => (
					<div
						onClick={() => changeValue(value)}
						className={(currentValue === value ? css.container : '')}
						key={i}
					>
						{value}
					</div>
				))}
			</div>
		</div>
	)
}