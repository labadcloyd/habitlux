import css from './button.module.css'

export default function Button(props) {
	const {children, onClick, primary, ...rest} = props

	return (
		<button
			className={css.btn}
			id={(primary ? css.btnPrimary : css.btnSecondary)}
			onClick={onClick} 
			{...rest}
		>
			{children} 
		</button>
	)
}