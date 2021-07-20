package mysql

import "github.com/DarkSoul94/task_tracker_server/models"

func (r *Repo) toModelGroup(dbGroup dbGroup) *models.Group {
	return &models.Group{
		ID:   dbGroup.ID,
		Name: dbGroup.Name,
	}
}

func (r *Repo) toDbLoginUser(name, passHash string) dbLoginUser {
	return dbLoginUser{
		Name:     name,
		PassHash: passHash,
	}
}

func (r *Repo) toModelUser(dbUser dbUser) models.User {
	return models.User{
		ID:       dbUser.ID,
		Name:     dbUser.Name,
		PassHash: dbUser.PassHash,
	}
}
