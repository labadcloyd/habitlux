import { AnimatePresence, motion } from "framer-motion"
import Button from "../button"
import css from './confirmModal.module.css'

export default function ConfirmModal(props) {
	const {message, showConfirmModal, action} = props

	return (
		<motion.div className={css.wrapper}>
			<motion.div
				initial={{opacity: 1, scale: 0}}
				animate={{opacity: 1, scale: 1}}
				transition={{ type: 'spring', duration: 0.3 }}
				exit={{opacity: 1, scale: 0}}
				layout
				
				className={css.container}
			>
				<h2>{message}</h2>
				<div className={css.btnContainer}>
					<Button
						onClick={() => {showConfirmModal(false)}}
						primary={false}
					>
						No
					</Button>
					<Button
						onClick={() => {action()}}
						primary={false}
					>
						Yes
					</Button>
				</div>

			</motion.div>
		</motion.div>
	)
}