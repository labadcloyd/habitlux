import { useEffect, useState } from "react"
import { useRouter } from 'next/router'

import styles from '../styles/auth.module.css'
import { MediumLogo, SmallArrow } from "../public/svgs"
import { TextInput, Button } from "../components/common"
import { RippleLoader } from '../public/loaders'

import { Signin, Signup } from '../common/services'
import { VerifyToken } from "../common/services/auth"

export default function Auth() {
  const router = useRouter()

	const [isSignup, setIsSignup] = useState(false)
	const [isLoading, setIsLoading] = useState(true)
	const [username, setUsername] = useState("")
	const [password, setPassword] = useState("")
	const [errorMsgs, setErrorMsgs] = useState([])

	function handleChangeForm() {
		setUsername("")
		setPassword("")
		setErrorMsgs([])
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
		return currentErrors.length
	}

	async function submitData () {
		const data = { username, password }
		if (isSignup) {
			const errLength = validateForm()
			if (errLength > 0) { return }

			const res = await Signup(data)
			if (res.status !== 200) {
				const currentErrors = []
				if (res.status === 400) {
					if (res.data.Number === 1062) {
						currentErrors.push("Username already taken")
					} else {
						currentErrors.push("Invalid password")
					}
				} else {
					currentErrors.push("Server error, please try again.")
				}
				setErrorMsgs(currentErrors)
			}
		}
		if (!isSignup) {
			const errLength = validateForm()
			if (errLength > 0) { return }

			const res = await Signin(data)
			console.log(res)
			if (res.status !== 200) {
				const currentErrors = []
				currentErrors.push("Username and password do not match")
				setErrorMsgs(currentErrors)
			}
		}
	}

	async function verifySession() {
		const res = await VerifyToken()
		if (res.status === 200) {
			return router.push('/dashboard')
		}
		setIsLoading(false)
	}

	useEffect(() => {
		verifySession()
	}, [])

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
							
							{(errorMsgs.length > 0) &&
								<div className={styles.errorWrapper}>
										{errorMsgs.map((value, i) => (
											<h4 key={i}>{value}</h4>
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