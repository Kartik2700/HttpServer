package Entity

type User struct{
	ID int
	Username String
	Password String
	Email String
}

func newUser(userName String, password String, email String)*User{
	user := new(User)
	user.ID = 0
	user.Username = userName
    user.Password = password
    user.Email = email
    return user
}