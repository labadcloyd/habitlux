import css from './textArea.module.css'

export default function TextArea(props) {
	const {...rest} = props

	return (
		<textarea className={css.wrapper} rows="10" {...rest}></textarea>
	)
}