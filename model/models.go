package model
import(
	"github.com/dgrijalva/jwt-go"
)

type User struct{
	UID string
	Password string
	//Token string
}

type Claims struct{
	ID string
	jwt.StandardClaims
}

type Date struct{
	Year int
	Month int
	Day int
	UserId string
}

type DatePlus struct{
	Date
	//Done bool
	Memory bool
}

type Schedule struct{
	Date
	Content string
	Done bool
}

type Door struct{
	Date
}

type Photo struct{
	Date
	PhotoUrl string
	InsertTime string
	Number int
}



type Fundmt struct{
	Code int `json:"code"`
	Message string `json:"message"`
}
type SetDate struct{
	Fundmt
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`
}
type Mem struct{
	Fundmt
	IfMemory bool
}
type Token struct{
	Fundmt
	Token string `json:"token"`
}
type Sch struct{
	Fundmt
	Schedule struct{
		Schedule
	} `json:"schedule"`//我感觉这里的数据类型是要改的//明早再说好了
}
type SchAndPh struct{
	Sch
	Photo  struct{
		Photo
	} `json:"photo"`
}
type Url struct{
	Fundmt
	Url string `json:"url"`
}
