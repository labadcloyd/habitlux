import Link from 'next/link'
import { useState } from "react"

import styles from '../styles/auth.module.css'
import { MediumLogo, SmallArrow } from "../public/svgs"
import { TextInput, Button } from "../components/common"
import { RippleLoader } from '../public/loaders'

export default function Auth() {
	const [isLogin, setIsLogin] = useState(false)
	const [isLoading, setIsLoading] = useState(false)
	const [username, setUsername] = useState("")
	const [password, setPassword] = useState("")

	function handleChangeForm() {
		setUsername("")
		setPassword("")
		setIsLoading(true)
		setIsLogin((prevState) => !prevState)

		setTimeout(() => {
			setIsLoading(false)
		}, 500);
	}

	function submitData () {
		const data = {
			username,
			password
		}
		if (isLogin) {
		}
	}

	return (
		<div className={styles.pageWrapper}>
			<div className={styles.pageContainer}>
				<a className={styles.logoContainer} href="/">
					<MediumLogo/>
				</a>

				<div className={styles.formWrapper}>
					{isLoading?
						<div className={styles.loadingWrapper}>
							<RippleLoader/>
						</div>
					:
						<div className={styles.formContainer}>
							<h2>{!isLogin ? "Sign in to your account" : "Register an account"}</h2>
							<TextInput
								value={username}
								setValue={setUsername}
								title={"Username"}
								minLength="3"
								required
							/>
							<div>
								<TextInput
									value={password}
									setValue={setPassword}
									title={"Password"}
									type="password"
									minLength="10"
									required
								/>
								<p className={styles.changeFormBtn} onClick={handleChangeForm}>
								{!isLogin ? "Register an account" : "Sign in to your account"}
									<SmallArrow/>
								</p>
							</div>
							<Button value={"Continue"} onClick={submitData} />
						</div>
					}
				</div>
			</div>
		</div>
	)
}