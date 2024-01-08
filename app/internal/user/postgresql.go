package user

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

func (r repository) Create(ctx context.Context, user User) (uuid string, err error) {
	r.logger.Infof("postgres create")
	err = r.client.QueryRow(ctx, `INSERT INTO public."user" (uuid, username, first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid`,
		user.Uuid,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&uuid)

	return uuid, err
}

func (r repository) FindByUUID(ctx context.Context, uuid string) (User, error) {
	r.logger.Debugf("recieve uuid: %s", uuid)
	var user User

	row := r.client.QueryRow(ctx, `SELECT id, uuid, username, first_name, last_name, email FROM public."user" WHERE uuid=$1`, uuid)
	err := row.Scan(
		&user.Id,
		&user.Uuid,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)
	if err != nil {

	}
	r.logger.Debug(user.Username)
	return user, nil
}

func (r repository) FindOne(ctx context.Context, uuid string) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, user User) error {
	r.logger.Infof("recieve user data: %v", user)

	return nil
}

func (r repository) Delete(ctx context.Context, uuid string) error {
	r.logger.Debugf("recieve uuid: %s", uuid)
	if _, err := r.client.Exec(ctx, `delete from public."user" where uuid=$1`, uuid); err == nil {
		return err
	}
	r.logger.Debugf("user with uuid: %s was deleted", uuid)
	return nil
}
