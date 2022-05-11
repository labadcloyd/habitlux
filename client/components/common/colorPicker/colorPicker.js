import { useEffect, useRef, useState } from 'react'
import { DEFAULT_COLOR_LIST } from '../../../common/constants/defaults'
import { useOutsideAlerter } from '../../../common/utils'
import css from './colorPicker.module.css'

export default function ColorPicker(props) {
	const {value, setValue, children} = props

	const colorPickerRef = useRef(null)

	const [isOpen, setIsOpen] = useState(false)
	const [currentColor, setCurrentColor] = useState(value || "#62A1FF")

	
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
		<div className={css.wrapper}>
			<div className={css.titleWrapper}>
				<div 
					className={css.colorContainer} 
					style={{backgroundColor: currentColor}} 
					onClick={() => (setIsOpen(true))}
				/>
				<div className={css.titleContainer}>{children}</div>
			</div>

			{isOpen &&
				<div className={css.colorPickerWrapper} ref={colorPickerRef}>
					{DEFAULT_COLOR_LIST.map((color, i) => (
						<div
							key={i}
							className={css.colorContainer} 
							style={{backgroundColor: color}}
							onClick={() => handleChange(color)}						
						/>
					))}
				</div>
			}
		</div>
	)
}