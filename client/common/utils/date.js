import moment from 'moment'

function getDateMonthly(dateValue) {
	let dateformat = "YYYY-MM-DD";
	let date = moment(dateValue || moment()).startOf("month"),
		weeklength=date.daysInMonth(), 
		result=[];
	while(weeklength--) {
		const habit = {
			date_created: date.format(dateformat),
			repeatCount: 0,
			requiredRepeatCount: 0,
			comment: ''
		}
		result.push(habit);
		date.add(1,"day")
	}
	return result;
}
function incrementMonthly(lastDateValue) {
	const newDate = moment(lastDateValue || moment()).add(1, "months")
	return getDateMonthly(newDate)
}
function decrementMonthly(firstDateValue) {
	const newDate = moment(firstDateValue || moment()).subtract(1, "months")
	return getDateMonthly(newDate)
}



function getDateBiWeekly(dateValue) {
	let dateformat = "YYYY-MM-DD";
	let date = moment(dateValue || moment()).startOf("week"), weeklength=14, result=[];
	while(weeklength--) {
		const habit = {
			date_created: date.format(dateformat),
			repeatCount: 0,
			requiredRepeatCount: 0,
			comment: ''
		}
		result.push(habit);
		date.add(1,"day")
	}
	return result;
}
function incrementBiWeekly(lastDateValue) {
	const newDate = moment(lastDateValue || moment()).add(1, "week")
	return getDateBiWeekly(newDate)
}
function decrementBiWeekly(firstDateValue) {
	const newDate = moment(firstDateValue || moment()).subtract(2, "weeks")
	return getDateBiWeekly(newDate)
}


function getDateWeekly(dateValue) {
	let dateformat = "YYYY-MM-DD";
	let date = moment(dateValue || moment()).startOf("week"), weeklength=7, result=[];
	while(weeklength--) {
		const habit = {
			date_created: date.format(dateformat),
			repeatCount: 0,
			requiredRepeatCount: 0,
			comment: ''
		}
		result.push(habit);
		date.add(1,"day")
	}
	return result;
}
function incrementWeekly(lastDateValue) {
	const newDate = moment(lastDateValue || moment()).add(1, "weeks")
	return getDateWeekly(newDate)
}
function decrementWeekly(firstDateValue) {
	const newDate = moment(firstDateValue || moment()).subtract(1, "weeks")
	return getDateWeekly(newDate)
}

export {
	getDateMonthly, 
	incrementMonthly, 
	decrementMonthly, 
	getDateBiWeekly, 
	incrementBiWeekly,
	decrementBiWeekly,
	getDateWeekly,
	incrementWeekly,
	decrementWeekly
}