import { useEffect, useRef, useState } from 'react'
import { AnimatePresence, motion } from "framer-motion"

import { DEFAULT_COLOR_LIST } from '../../../common/constants/defaults'
import { useOutsideAlerter } from '../../../common/utils'
import css from './colorPicker.module.css'

export default function ColorPicker(props) {
	const {value, setValue, children} = props

	const colorPickerRef = useRef(null)

	const [isOpen, setIsOpen] = useState(false)
	const [currentColor, setCurrentColor] = useState(value || "98 161 255")

	
	function handleClickOutside(event) {
		if (colorPickerRef.current && !colorPickerRef.current.contains(event.target)) {
			setIsOpen(false)
		}
	}
	useOutsideAlerter({ref: colorPickerRef, handleEvent: handleClickOutside})

	function handleChange(clickedColor) {
		setCurrentColor(clickedColor)
		setValue(clickedColor)
		setIsOpen(false)
	}

	useEffect(() => {
		setCurrentColor(value)
	},[value])

	return (
		<motion.div className={css.wrapper}>
			<motion.div className={css.titleWrapper}>
				<motion.div
					whileHover={{ scale:1.1, boxShadow: '0px 0px 15px rgba(0, 0, 0, 0.2)' }}
					whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}

					className={css.colorContainer} 
					style={{backgroundColor: `rgb(${currentColor})`}} 
					onClick={() => (setIsOpen(true))}
				/>
				<motion.div className={css.titleContainer}>{children}</motion.div>
			</motion.div>

			{isOpen &&
				<AnimatePresence>
					<motion.div 
						initial={{opacity: 1, scale: 0}}
						animate={{opacity: 1, scale: 1}}
						transition={{ type: 'spring', duration: 0.3 }}
						exit={{opacity: 1, scale: 0}}
						layout
					
						className={css.colorPickerWrapper} ref={colorPickerRef}
					>
						{DEFAULT_COLOR_LIST.map((color, i) => (
							<motion.div
								whileHover={{ scale:1.1, boxShadow: '0px 0px 8px rgba(0, 0, 0, 0.2)' }}
								whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
								
								key={i}
								className={css.colorContainer} 
								style={{backgroundColor: `rgb(${color})`}}
								onClick={() => handleChange(color)}						
							/>
						))}
					</motion.div>
				</AnimatePresence>
			}
		</motion.div>
	)
}