import newAxios from "./axios";
import { 
	getDateMonthly, 
	incrementMonthly, 
	decrementMonthly, 
	getDateBiWeekly, 
	incrementBiWeekly,
	decrementBiWeekly,
	getDateWeekly,
	incrementWeekly,
	decrementWeekly
} from "./date"
import { addHabitsToDate, getTodaysHabits } from "./format";
import { useOutsideAlerter, calcBgColor } from "./helper";
import { localLogin, localLogout, isLoggedIn} from "./auth";

export {
	newAxios,
	getDateMonthly, 
	incrementMonthly, 
	decrementMonthly, 
	getDateBiWeekly, 
	incrementBiWeekly,
	decrementBiWeekly,
	getDateWeekly,
	incrementWeekly,
	decrementWeekly,
	addHabitsToDate,
	useOutsideAlerter,
	calcBgColor,
	getTodaysHabits,
	localLogin,
	localLogout,
	isLoggedIn
}