import axios from "axios";
import { BASE_URL } from "../constants";

// make sure to allow credentials in order to set cookies in each req
const newAxios = axios.create({
	withCredentials: true,
	baseURL: BASE_URL,
})

export default newAxios