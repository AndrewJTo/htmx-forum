package daos

import (
	"time"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

func (dao Dao) GetUserByEmail(email string) (*model.Users, error) {
	stmt := table.Users.SELECT(
		table.Users.AllColumns.Except(table.Users.Password),
	).WHERE(
		table.Users.Email.EQ(postgres.String(email)),
	).LIMIT(1)

	var dest model.Users

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) GetUserById(id int32) (*model.Users, error) {
	stmt := table.Users.SELECT(
		table.Users.AllColumns.Except(table.Users.Password),
	).WHERE(
		table.Users.ID.EQ(postgres.Int32(id)),
	).LIMIT(1)

	var dest model.Users

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) GetUserPassword(user *model.Users) error {
	stmt := table.Users.SELECT(
		table.Users.Password,
	).WHERE(
		table.Users.ID.EQ(postgres.Int32(user.ID)),
	).LIMIT(1)

	var dest model.Users
	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return err
	}

	user.Password = dest.Password

	return nil
}

func (dao Dao) CreateUser(user *model.Users) (*model.Users, error) {
	user.CreatedAt = time.Now()
	stmt := table.Users.INSERT(
		table.Users.Email, table.Users.Name, table.Users.Password,
		table.Users.IsAdmin, table.Users.BanUntil, table.Users.CreatedAt,
	).MODEL(user)

	var dest model.Users

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) SaveUser(user *model.Users) error {
	stmt := table.Users.UPDATE(
		table.Users.Email, table.Users.Name, table.Users.Password,
		table.Users.IsAdmin, table.Users.BanUntil,
	).MODEL(user).
		WHERE(table.Users.ID.EQ(postgres.Int32(user.ID)))

	_, err := stmt.Exec(dao.DB)

	return err
}
