import css from './smallBtn.module.css'

export default function SmallBtn(props) {
	const {children, onClick, ...rest} = props

	return (
		<button
			className={css.btn}
			onClick={onClick} 
			{...rest}
		>
			{children} 
		</button>
	)
}