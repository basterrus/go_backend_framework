package roles

import (
	"context"
	"github.com/basterrus/go_backend_framework/pkg/client"
	"github.com/basterrus/go_backend_framework/pkg/logging"
)

type repository struct {
	client client.PgClient
	logger *logging.Logger
}

func NewStorage(client client.PgClient, logger logging.Logger) Storage {
	return &repository{
		client: client,
		logger: &logger,
	}
}

func (r repository) CreateRole(ctx context.Context, role Role) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetAllRoles(ctx context.Context) ([]Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetRoleByID(ctx context.Context, id string) (Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, role Role) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	r.logger.Debugf("[Delete Role] receive role id: %s", id)
	tx, err := r.client.Begin(ctx)
	if err != nil {
		r.logger.Debugf("[Delete Role] error to begin transaction: %s", err)
	}
	if _, err := tx.Exec(ctx, `delete from public."roles" where id=$1`, id); err != nil {
		tx.Rollback(ctx)
		return err
	}

	r.logger.Debugf("[Delete Role] user with uuid: %s was deleted", id)

	return tx.Commit(ctx)
}

//func (r repository) Create(ctx context.Context, user User) (uuid string, err error) {
//	r.logger.Infof("[POSTGRES client] create user item")
//	err = r.client.QueryRow(ctx, `INSERT INTO public."user" (uuid, username, first_name, last_name, email, password_hash)
//										VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid`,
//		user.Uuid,
//		user.Username,
//		user.FirstName,
//		user.LastName,
//		user.Email,
//		user.Password,
//	).Scan(&uuid)
//
//	return uuid, err
//}
//
//func (r repository) FindByUUID(ctx context.Context, uuid string) (User, error) {
//	r.logger.Debugf("[POSTGRES client] recieve uuid: %s", uuid)
//	var user User
//
//	row := r.client.QueryRow(ctx, `SELECT id, uuid, username, first_name, last_name, email FROM public."user" WHERE uuid=$1`, uuid)
//	err := row.Scan(
//		&user.Id,
//		&user.Uuid,
//		&user.Username,
//		&user.FirstName,
//		&user.LastName,
//		&user.Email,
//	)
//	if err != nil {
//
//	}
//	r.logger.Debug(user.Username)
//	return user, nil
//}
//
//func (r repository) FindOne(ctx context.Context, uuid string) (User, error) {
//	r.logger.Debugf("[FindOne User] recieve user uuid: %s", uuid)
//	tx, err := r.client.Begin(ctx)
//	if err != nil {
//		r.logger.Debugf("[Delete User] error to begin transaction: %s", err)
//	}
//
//	if _, err := tx.Exec(ctx, ``, uuid); err != nil {
//
//	}
//
//	return User{}, nil
//}
//
//func (r repository) Update(ctx context.Context, user User) error {
//	r.logger.Infof("recieve user data: %v", user)
//
//	return nil
//}
