import { motion } from "framer-motion"
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
			<div className={css.wrapper} >
				{values && values.map((value, i) => (
					<motion.div
						whileHover={{ scale:1.1, boxShadow: '0px 0px 15px rgba(0, 0, 0, 0.2)' }}
						whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
						transition={{ duration: 0.2, delay: 0 }}

						onClick={() => changeValue(value)}
						className={(currentValue === value ? css.container : '')}
						key={i}
					>
						{value}
					</motion.div>
				))}
			</div>
		</div>
	)
}