package mysql

import (
	"github.com/DarkSoul94/task_tracker_server/models"
)

func (r *Repo) toModelGroup(dbGroup dbGroup) models.Group {

	mGroup := models.Group{
		ID:          dbGroup.ID,
		Name:        dbGroup.Name,
		Permissions: dbGroup.Permissions,
	}
	return mGroup
}

func (r *Repo) toModelUser(dbUser *dbUser) models.User {
	var group models.Group
	if dbUser.GroupID != 0 {
		group = models.Group{ID: dbUser.GroupID}
	} else {
		group = models.Group{}
	}
	return models.User{
		ID:         dbUser.ID,
		Email:      dbUser.Email,
		Name:       dbUser.Name,
		Group:      &group,
		Department: dbUser.Department,
	}
}

func (r *Repo) toDBUser(mUser *models.User) dbUser {
	user := dbUser{
		ID:         mUser.ID,
		Email:      mUser.Email,
		Name:       mUser.Name,
		Department: mUser.Department,
	}
	if mUser.Group != nil {
		user.GroupID = mUser.Group.ID
	}
	return user
}
