package model
import(
	"github.com/dgrijalva/jwt-go"
)

type User struct{
	UID string
	Password string
	Token string
}
type Claims struct{
	ID string
	jwt.StandardClaims
}
type Date struct{
	Year int
	Month int
	Day int
	Schedules[20] Schedule
	Photos[5] string

}

type Schedule struct{
	Content string
	Done bool
}

type Door struct{
	Year int
	Month int
	Day int
}

