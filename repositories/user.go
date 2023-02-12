package repositories

type IUserRepository interface {
	UploadUsers()
	GetUser()
}

type UserSQLRepository struct {
}

func InitUserSQLRepository() *UserSQLRepository {
	ur := new(UserSQLRepository)
	return ur
}

func (s *UserSQLRepository) UploadUsers() {

}

func (s *UserSQLRepository) GetUser() {

}

