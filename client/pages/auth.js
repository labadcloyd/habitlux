import { useState } from "react"
import { TextInput } from "../components/common"
import Button from "../components/common/button"
import MediumLogo from "../public/svgs/mediumLogo"
import styles from '../styles/auth.module.css'

export default function Auth() {
	const [isLogin, setIsLogin] = useState(false)
	const [isLoading, setIsLoading] = useState(false)
	const [username, setUsername] = useState("")
	const [password, setPassword] = useState("")

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
				<div className={styles.logoContainer}>
					<MediumLogo/>
				</div>

				<div className={styles.formWrapper}>
					<h2>Sign in to your account</h2>
					<TextInput
						value={username}
						setValue={setUsername}
						title={"Username"}
					/>
					<TextInput
						value={password}
						setValue={setPassword}
						title={"Password"}
						type="password"
					/>
					<Button value={"Continue"} onClick={submitData} />
				</div>
			</div>
		</div>
	)
}