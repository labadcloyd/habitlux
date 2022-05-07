import css from './textArea.module.css'

export default function TextArea(props) {
	const {updateValue, value, ...rest} = props

	function updateCurrentValue(e) {
		updateValue(e.target.value)
	}

	return (
		<textarea 
			onChange={updateCurrentValue}
			className={css.wrapper}
			value={value}
			rows="10"
			{...rest}
		/>
	)
}