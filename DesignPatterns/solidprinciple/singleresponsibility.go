package solidprinciple

type User struct {
	Name  string
	Id    int
	Email string
}

//
func (user *User) Save() {
	//Save user to repository
}

func (user *User) Send() {
	// Send mail to user
}

// Here the above method have different responsibility for the User Struct

// To achive the Single Responsibility Principle we can use like this

type UserRepository interface {
	Save(user *User) error
}

type EmailService interface {
	Send(user *User) error
}


