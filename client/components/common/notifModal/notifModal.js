import { AnimatePresence, motion } from "framer-motion"
import { useState } from "react"
import { Check, Close, ErrorIcon } from "../../../public/svgs"
import SmallBtn from "../smallBtn"

import css from './notifModal.module.css'

export default function NotifModal(props) {
	const {modalOpen, setModalOpen, content, setContent} = props
	const {msg, error} = content

	function closeModal() {
		setModalOpen(false)
		setContent({msg: "", error: false})
	}

	return (
		<AnimatePresence>
			{modalOpen &&
				<motion.div
					initial={{x: "-10vw"}}
					animate={{x: 0}}
					transition={{ type: 'spring', duration: 0.3 }}
					exit={{x: "100vw"}}

					className={css.wrapper}
					style={{backgroundColor: error ? "#ffc5c5" : "#c6ffc5"}}
				>
					<div style={{backgroundColor: error ? "#ff4646" : "#60f95d"}} className={css.iconContainer}>
						{error? <ErrorIcon/> : <Check/>}
					</div>
					<span>{msg || ""}</span>
					<SmallBtn
						style={{backgroundColor: error ? "#ffc5c5" : "#c6ffc5"}}
						onClick={closeModal}
					>
						<Close/>
					</SmallBtn>
				</motion.div>
			}
		</AnimatePresence>
	)
}