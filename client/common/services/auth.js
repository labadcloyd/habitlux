import { AUTH_ENDPOINTS } from '../constants'
import { newAxios } from '../utils/axios'

export async function Signin({ username, password }) {
	try {
		const res = await newAxios.post(AUTH_ENDPOINTS.SIGNIN, { username, password })
		return res
	} catch (err) {
		return err.response
	}
}

export async function Signup({ username, password }) {
	try {
		const res = await newAxios.post(AUTH_ENDPOINTS.SIGNUP, { username, password })
		return res
	} catch (err) {
		return err.response
	}
}

export async function VerifyToken() {
	try {
		const res = await newAxios.get(AUTH_ENDPOINTS.VERIFYTOKEN)
		return res
	} catch (err) {
		return err.response
	}
}