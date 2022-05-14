import css from './loader.module.css'

export default function Loader() {
	return (
		<div className={css.loadingBox}>
			<div className={css.loader}></div>
		</div>
	)
}