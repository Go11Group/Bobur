package postgres

import (
	"database/sql"
	"time"
	"user_profile/models"
)

type UserProfileRepository struct {
	Db *sql.DB
}

func NewUserProfileRepositoryRepository(db *sql.DB) *UserProfileRepository {
	return &UserProfileRepository{Db: db}
}

func (repo UserProfileRepository) CreateUserProfile(user_profile *models.UserProfile) (interface{}, error) {
	return repo.Db.Exec("insert into user_profiles(user_id,full_name,bio,role,location,avatar_url,website,created_at)", user_profile.UserId, user_profile.Fullname, user_profile.Bio, user_profile.Role, user_profile.Location, user_profile.AvatarUrl, user_profile.Website, time.Now())
}
func (repo UserProfileRepository) UpdateUserProfile(user_profile *models.UserProfile, id string) (interface{}, error) {
	return repo.Db.Exec("update user_profiles set user_id=$1,full_name=$2,bio=$3,role=$4,,location=$5,avatar_url=$6,website=$7,updated_at=$8 where id=$9 and deleted_at=0)", user_profile.UserId, user_profile.Fullname, user_profile.Bio, user_profile.Role, user_profile.Location, user_profile.AvatarUrl, user_profile.Website, time.Now(), id)
}
func (repo UserProfileRepository) DeleteUserProfile(id string) (interface{}, error) {
	return repo.Db.Exec("update user_profiles set deleted_at=$1 where id=$2 and deleted_at is null)", time.Now(), id)
}
func (repo UserProfileRepository) GetUserProfileById(id string) (*models.UserProfile, error) {
	rows, err := repo.Db.Query("select user_id,full_name,bio,role,location,avatar_url,website from user_profiles  where id=$1 and deleted_at is null)", id)
	if err != nil {
		return nil, err
	}
	var user_profile models.UserProfile
	for rows.Next() {
		err := rows.Scan(&user_profile.UserId, &user_profile.Fullname, &user_profile.Bio, &user_profile.Role, &user_profile.Location, &user_profile.AvatarUrl, &user_profile.Website)
		if err != nil {
			return nil, err
		}
		return &user_profile, nil
	}
	return nil, err
}
func (repo UserProfileRepository) GetUserProfile(user_profile *models.UserProfile) (*[]models.UserProfile, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(user_profile.UserId) > 0 {
		params["user_id"] = user_profile.UserId
		filter += " and user_id = :user_id "

	}
	if len(user_profile.Fullname) > 0 {
		params["full_name"] = user_profile.Fullname
		filter += " and full_name = :full_name "

	}
	if len(user_profile.Bio) > 0 {
		params["bio"] = user_profile.Bio
		filter += " and bio = :bio "

	}
	if len(user_profile.Role) > 0 {
		params["role"] = user_profile.Role
		filter += " and role = :role "

	}
	if len(user_profile.Location) > 0 {
		params["location"] = user_profile.Location
		filter += " and location = :location "

	}
	if len(user_profile.AvatarUrl) > 0 {
		params["avatar_url"] = user_profile.AvatarUrl
		filter += " and avatar_url = :avatar_url "

	}
	if len(user_profile.Website) > 0 {
		params["website"] = user_profile.Website
		filter += " and website = :website "

	}

	if user_profile.Limit > 0 {
		params["limit"] = user_profile.Limit
		limit = ` LIMIT :limit`

	}
	if user_profile.Offset > 0 {
		params["offset"] = user_profile.Offset
		limit = ` OFFSET :offset`

	}
	query := "select id,user_id,full_name ,bio, role, location, avatar_url, website from user_profiles  where  deleted_at is null"

	query = query + filter + limit + offset
	query, arr = ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	var user_profiles []models.UserProfile
	for rows.Next() {
		var user_profile models.UserProfile
		err := rows.Scan(&user_profile.Id, &user_profile.UserId, &user_profile.Fullname, &user_profile.Bio, &user_profile.Role, &user_profile.Location, &user_profile.AvatarUrl, &user_profile.Website)
		if err != nil {
			return nil, err
		}
		user_profiles = append(user_profiles, user_profile)
	}
	return &user_profiles, err
}
