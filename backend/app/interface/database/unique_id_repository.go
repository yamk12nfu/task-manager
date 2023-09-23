package database

import (
	"context"
	"errors"
	"task-manager/app/utils"
)

type UniqueIDRepository struct {
	SQLHandler
}

func (r *UniqueIDRepository) Issue(ctx context.Context) (string, error) {
	for tries := 0; tries < 5; tries++ {
		id, err := utils.GetRandomString(utils.Length(8), utils.LowerCases(), utils.Numbers())
		if err != nil {
			return "", err
		}

		_, err = r.NamedExec(
			ctx,
			`insert into unique_ids (unique_id) values (:unique_id)`,
			map[string]any{"unique_id": id},
		)

		if r.SqlError(err) == SqlErrDuplicate {
			continue
		}

		return id, err
	}

	return "", errors.New("failed to issue unique id")
}
