package main

import "errors"

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

func (u *UserList) Find(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, errors.New("not found")
}

type UserListProxy struct {
	Db UserList
	Cache UserList
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.Cache.Find(id)
	if err == nil {
		return user, nil
	}

	return u.Db.Find(id)
}
