import { motion } from "framer-motion"
import css from './button.module.css'

export default function Button(props) {
	const {children, onClick, primary, ...rest} = props

	return (
		<motion.button
			whileHover={{ scale:1.1, boxShadow: '0px 0px 15px rgba(0, 0, 0, 0.2)' }}
			whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
			layout

			className={css.btn}
			id={(primary ? css.btnPrimary : css.btnSecondary)}
			onClick={onClick} 
			{...rest}
		>
			{children} 
		</motion.button>
	)
}