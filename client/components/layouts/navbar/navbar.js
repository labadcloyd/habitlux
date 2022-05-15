import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { GetUser } from '../../../common/services'
import { isLoggedIn, localLogout } from '../../../common/utils'

import { MediumLogo } from '../../../public/svgs'
import css from './navbar.module.css'

export default function Navbar() {
	const [username, setUsername] = useState()
	
	const router = useRouter()

	async function checkIsLoggedIn() {
		if (isLoggedIn() === true) {
			try {
				const res = await GetUser()
				setUsername(res.data.username)
			} catch (err) {
				if (err.response.status === 401) {
					localLogout()
					return router.push('/auth')
				}
			}
		}
	}

	async function logout() {
		await localLogout()
		router.push('/auth')
	}

	useEffect(() => {
		checkIsLoggedIn()
	},[])

	return (
		<div className={css.navbarWrapper}>
			<div className={css.navbarContainer}>

				<div className={css.leftWrapper} onClick={() => {router.push('/')}}>
					<div className={css.logoContainer}>
						<MediumLogo/>
					</div>
					<h1>Habitlux</h1>
				</div>

				<div className={css.rightWrapper}>
					{ username &&
						<h2 onClick={() => {router.push('/dashboard')}}> {username} </h2>
					}
					<h2 onClick={logout}>Logout</h2>
				</div>

			</div>
		</div>
	)
}