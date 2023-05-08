package repository

type UserDao struct {
}

func (dao *UserDao) Insert() {

}

func (dao *UserDao) Update() {

}

func (dao *UserDao) Delete() {

}

func (dao *UserDao) Query() {

}

// wire Provider
func NewUserDao() *UserDao {
	return &UserDao{}
}
