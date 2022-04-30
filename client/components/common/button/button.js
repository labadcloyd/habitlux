import style from './button.module.css'

export default function Button(props) {
	const {value, onClick, ...rest} = props

	return (
		<button
			className={style.btn}
			onClick={onClick} 
			{...rest}
		>
			{value} 
		</button>
	)
}