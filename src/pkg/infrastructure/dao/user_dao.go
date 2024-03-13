package dao

import (
	"database/sql"
	"go-minimum-crud/src/pkg/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

type UserDao struct {
	DB *sql.DB
}

func (c *UserDao) SelectAll() ([]model.User, error) {
	selected, err := c.DB.Query("SELECT id, name, email FROM user")
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for selected.Next() {
		user := model.User{}
		err = selected.Scan(&user.Id, &user.UserName, &user.UserEmail)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	selected.Close()
	return users, err
}

func (c *UserDao) Select(id model.UserId) (*model.User, error) {
	selected, err := c.DB.Query("SELECT id, name, email FROM user WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	user := model.User{}
	for selected.Next() {
		err = selected.Scan(&user.Id, &user.UserName, &user.UserEmail)
		if err != nil {
			return nil, err
		}
	}
	selected.Close()
	return &user, nil
}

func (c *UserDao) Insert(user model.NewUser) error {
	insert, err := c.DB.Prepare("INSERT INTO user (name, email) VALUES (?, ?)")
	if err != nil {
		return err
	}
	if _, err := insert.Exec(user.UserName, user.UserEmail); err != nil {
		return err
	}
	return nil
}

func (c *UserDao) Update(id model.UserId, user model.NewUser) error {
	insert, err := c.DB.Prepare("UPDATE user SET name=?, email=? WHERE id=?")
	if err != nil {
		return err
	}
	if _, err := insert.Exec(user.UserName, user.UserEmail, id); err != nil {
		return err
	}
	return nil
}

func (c *UserDao) Delete(id model.UserId) error {
	insert, err := c.DB.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		return err
	}
	if _, err := insert.Exec(id); err != nil {
		return err
	}
	return nil
}
