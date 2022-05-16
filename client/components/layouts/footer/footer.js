import css from './footer.module.css'

export default function Footer() {
	return (
		<div className={css.wrapper}>
			<div className={css.container}>
				<h6>© HabitLux 2022</h6>
				<h6>
					Designed and developed by: 
					<a href='https://github.com/labadcloyd' target="_blank"> Cloyd Abad</a>
				</h6>
			</div>
		</div>
	)
}