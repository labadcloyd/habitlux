import { MediumLogo } from '../../../public/svgs'
import css from './navbar.module.css'

export default function Navbar() {
	return (
		<div className={css.navbarWrapper}>
			<div className={css.navbarContainer}>
				<div className={css.logoContainer}>
					<MediumLogo/>
				</div>
				<h1>Habitlux</h1>
			</div>
		</div>
	)
}