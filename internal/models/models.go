package models
import (
	"time"
)

type Student struct{
	DNI		string
	Name	string
	Email 	string
	StudentCode		string
}

type Users struct{
	ID		int
	Name	string
	Email	string
	Password	string
	Role	string
}

type Courses struct{
	ID		int
	Name	string
	Code	string
	Group	int
	Professor	string	
} 

type Sessions struct{
	ID			int
	CourseID   int
	Date		time.Time
	StartTime	time.Time
	EndTime	time.Time
}

type Attendance struct{
	ID	int
	Present bool
	TakenBy int
	DNI		string
}

type Points struct{
	ID		int
	StudentDNI	string
	Amount		int
	SessionID	int
}