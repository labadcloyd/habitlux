import { Signout } from "../services/auth"

function localLogin() {
	localStorage.setItem('auth', JSON.stringify({isLoggedIn: true}))
}
async function localLogout() {
	localStorage.setItem('auth', JSON.stringify({isLoggedIn: false}))
	await Signout()
}
function isLoggedIn() {
	const auth = JSON.parse(localStorage.getItem('auth'))
	return auth?.isLoggedIn || false
}

export {localLogin, localLogout, isLoggedIn}
