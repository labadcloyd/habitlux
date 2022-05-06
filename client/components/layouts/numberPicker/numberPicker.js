import { useEffect, useState } from 'react'
import { ArrowLeft, ArrowRight } from '../../../public/svgs'
import { SmallBtn } from '../../common'
import css from './numberPicker.module.css'

export default function NumberPicker(props) {
	const {children, onClick, value, maxValue, setState, ...rest} = props

	const [currentCount, setCurrentCount] = useState(value)

	function updateCount (value){
		if ((currentCount+value) < 0) { return }
		if (maxValue) {
			if ((currentCount+value) > maxValue || currentCount > maxValue) {
				setCurrentCount(maxValue)
				return setState(maxValue)
			}
		}
		setCurrentCount(currentCount+value)
		setState(currentCount+value)
	}

	useEffect(() => {
		setCurrentCount(value)
	}, [value])

	return (
		<div className={css.wrapper} {...rest}>
			<div className={css.btnsWrapper}>
				<SmallBtn onClick={() => {updateCount(-1)}}>
					<ArrowLeft/>
				</SmallBtn>
				<div className={css.counterContainer}>
					{currentCount}
				</div>
				<SmallBtn onClick={() => {updateCount(+1)}}>
					<ArrowRight/>
				</SmallBtn>
			</div>
			<h6>
				{children}
			</h6>
		</div>
	)
}