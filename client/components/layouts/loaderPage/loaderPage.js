import { motion } from "framer-motion"
import { MediumLogo } from '../../../public/svgs'
import { Loader } from '../../common'
import css from './loaderPage.module.css'

export default function LoaderPage() {
	return (
		<motion.div
			className={css.wrapper}
			
			initial={{opacity: 0}}
			animate={{opacity: 1}}
			exit={{opacity: 0}}
			transition={{ duration: 0.3, delay: 0 }}
			layout
		>
			<div className={css.logoContainer}>
				<MediumLogo/>
			</div>
			<Loader />
		</motion.div>
	)
}