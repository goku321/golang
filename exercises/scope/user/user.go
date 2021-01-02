package user

type User struct{}

func (u *User) Name(name string) {
	if printName != nil {
		printName(name)
	}
}