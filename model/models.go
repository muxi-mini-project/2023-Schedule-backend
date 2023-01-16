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
	Year string
	Month string
	Day string
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
	Photo string
}
