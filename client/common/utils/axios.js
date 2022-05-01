import axios from "axios";
import { BASE_URL } from "../constants";

export const newAxios = axios.create({
	baseURL: BASE_URL
})