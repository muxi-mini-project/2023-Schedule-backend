package datechange
import(

)
func MonthTrans(m string)(int){
	var month int
	switch m{
	case "January":		month=1
	case "February":	month=2
	case "March":		month=3
	case "April":		month=4
	case "May":			month=5
	case "June":		month=6
	case "July":		month=1
	case "August":		month=2
	case "September":	month=3
	case "October":		month=4
	case "November":	month=5
	case "December":	month=6
	}
	return month
}//把字符串月份改成数字月份

var days=[13]int{0,31,28,31,30,31,30,31,31,30,31,30,31}

func Yesterday(y int,m int,d int)(int,int,int){
	if(d>1){
		d=d-1;
	}else{
		if(m>1){
			m=m-1;
			d=days[m]
		}else{
			y=y-1
			m=12
			d=days[m]
		}
		if m==2&&leapyear(y){
			d=d+1
		}
	}
	return y,m,d
}

func leapyear(y int)bool{
	if(y%4!=0){
		return false//四年一润
	}else if (y%100!=0){
		return true//百年不润
	}else if (y%400!=0){
		return false//四百年又润
	}else{
		return true
	}
}