import Link from 'next/link'
import { useEffect, useState } from "react"

import styles from '../styles/auth.module.css'
import { MediumLogo, SmallArrow } from "../public/svgs"
import { TextInput, Button } from "../components/common"
import { RippleLoader } from '../public/loaders'

export default function Auth() {
	const [isSignup, setIsSignup] = useState(false)
	const [isLoading, setIsLoading] = useState(false)
	const [username, setUsername] = useState("")
	const [password, setPassword] = useState("")
	const [errorMsgs, setErrorMsgs] = useState([])

	function handleChangeForm() {
		setUsername("")
		setPassword("")
		setIsLoading(true)
		setIsSignup((prevState) => !prevState)

		setTimeout(() => {
			setIsLoading(false)
		}, 500);
	}

	function validateForm() {
		const currentErrors = []
		if (username.length < 3) {
			currentErrors.push("Username length must not be less than 3")
		}
		if (password.length < 10) {
			currentErrors.push("Password length must not be less than 10")
		}
		if (username.length > 32 || password.length > 32 ) {
			currentErrors.push("Username or password length must not be morethan 32")
		}
		setErrorMsgs(currentErrors)
	}

	function submitData () {
		const data = {
			username,
			password
		}
		if (isSignup) {
		}
	}

	useEffect(() => {
		validateForm()
	}, [username, password])

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
							<h2>{!isSignup ? "Sign in to your account" : "Register an account"}</h2>
							
							{(errorMsgs.length > 0 && isSignup) &&
								<div className={styles.errorWrapper}>
										{errorMsgs.map((value) => (
											<h4>{value}</h4>
										))}
								</div>
							}
							
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
								{!isSignup ? "Register an account" : "Sign in to your account"}
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