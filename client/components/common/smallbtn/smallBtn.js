import { motion } from "framer-motion"
import css from './smallBtn.module.css'

export default function SmallBtn(props) {
	const {children, onClick, ...rest} = props

	return (
		<motion.button
			whileHover={{ scale:1.1, boxShadow: '0px 0px 5px rgba(0, 0, 0, 0.2)' }}
			whileTap={{ scale:0.9, boxShadow: '0px 0px 0px rgba(0, 0, 0, 0.5)' }}
			transition={{ duration: 0.2, delay: 0 }}

			className={css.btn}
			onClick={onClick} 
			{...rest}
		>
			{children} 
		</motion.button>
	)
}