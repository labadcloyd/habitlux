import axios from "axios";
import { BASE_URL } from "../constants";

export const newAxios = axios.create({
	withCredentials: true,
	baseURL: BASE_URL
})