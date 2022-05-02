
import css from '../styles/dashboard.module.css'
import { Navbar, DateChanger } from "../components/layouts";

export default function Dashboard() {

	return(
		<div className={css.pageWrapper}>
			<Navbar/>
			<div className={css.pageContainer}>
				<div className={css.rowContainer}>
					<DateChanger/>
				</div>
			</div>
		</div>
	)
}