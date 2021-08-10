package mysql

import (
	"encoding/json"
	"strconv"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/pkg/logger"
)

func (r *Repo) toModelGroup(dbGroup dbGroup) models.Group {
	mGroup := models.Group{
		ID:   dbGroup.ID,
		Name: dbGroup.Name,
	}
	temp := make(map[string][]string)
	err := json.Unmarshal(dbGroup.Permissions, &temp)
	if err != nil {
		logger.LogError(ErrReadGroup.Error(), "user_manager/repo/mysql", strconv.FormatUint(dbGroup.ID, 10), err)
	}
	mGroup.Permissions = temp
	return mGroup
}

func (r *Repo) toDbLoginUser(name, passHash string) dbLoginUser {
	return dbLoginUser{
		Name:     name,
		PassHash: passHash,
	}
}

func (r *Repo) toModelUser(dbUser dbUser) models.User {
	var group models.Group
	if dbUser.GroupID != 0 {
		group = models.Group{ID: dbUser.GroupID}
	} else {
		group = models.Group{}
	}
	return models.User{
		ID:       dbUser.ID,
		Name:     dbUser.Name,
		PassHash: dbUser.PassHash,
		Group:    &group,
	}
}
