import { MediumLogo } from '../../../public/svgs'
import { Loader } from '../../common'
import css from './loaderPage.module.css'

export default function LoaderPage() {
	return (
		<div className={css.wrapper}>
			<div className={css.logoContainer}>
				<MediumLogo/>
			</div>
			<Loader />
		</div>
	)
}