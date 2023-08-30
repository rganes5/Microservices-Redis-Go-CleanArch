package repository

import (
	"X-TENTIONCREW/auth_svc/pkg/domain"
	"X-TENTIONCREW/auth_svc/pkg/pb"
	"X-TENTIONCREW/auth_svc/pkg/repository/interfaces"
	"X-TENTIONCREW/auth_svc/pkg/utils"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type authRepo struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewauthRepo(db *gorm.DB, redis *redis.Client) interfaces.AuthRepo {
	return &authRepo{
		DB:    db,
		Redis: redis,
	}
}

//WITHOUT REDIS

/*
func (c *authRepo) Register(ctx context.Context, req *pb.RegisterRequest) (int32, error) {
	user := domain.User{
		CreatedAt: time.Now(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}
	err := c.DB.Create(&user).Error
	if err != nil {
		return int32(user.ID), err
	}
	return int32(user.ID), nil
}

func (c *authRepo) GetUser(ctx context.Context, id int32) (utils.Response, error) {
	var user utils.Response
	query := `SELECT * FROM users WHERE id=$1`
	result := c.DB.Raw(query, id).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (c *authRepo) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (utils.Response, error) {
	query := `
		UPDATE users
		SET first_name= $1, last_name=$2, email=$3, phone=$4
		WHERE id = $5
	`
	result := c.DB.Exec(query, req.FirstName, req.LastName, req.Email, req.Phone, req.ID)
	if result.Error != nil {
		return utils.Response{}, result.Error
	}

	return utils.Response{}, nil
}

func (c *authRepo) DeleteUser(ctx context.Context, id int32) error {
	query := `
	DELETE FROM users
	WHERE id= $1
	`
	result := c.DB.Exec(query, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
*/

//WITH REDIS

func (c *authRepo) Register(ctx context.Context, req *pb.RegisterRequest) (int32, error) {
	user := domain.User{
		CreatedAt: time.Now(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}
	err := c.DB.Create(&user).Error
	if err != nil {
		return int32(user.ID), err
	}

	// Cache the user details in Redis
	err = c.cacheUserResponse(user)
	if err != nil {
		fmt.Println("Failed to cache")
	}

	return int32(user.ID), nil
}

func (c *authRepo) GetUser(ctx context.Context, id int32) (utils.Response, error) {
	cachedUser, err := c.getUserFromCache(id)
	if err == nil {
		return cachedUser, nil
	}

	var user domain.User
	query := `
		SELECT id, first_name, last_name, email, phone
		FROM users
		WHERE id = $1
	`
	result := c.DB.Raw(query, id).Scan(&user)
	if result.Error != nil {
		return utils.Response{}, result.Error
	}

	err = c.cacheUserResponse(user)
	if err != nil {
		fmt.Println("Failed to cache")
	}

	response := utils.Response{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	return response, nil
}

func (c *authRepo) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (utils.Response, error) {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, email = $3, phone = $4
		WHERE id = $5
	`
	result := c.DB.Exec(query, req.FirstName, req.LastName, req.Email, req.Phone, req.ID)
	if result.Error != nil {
		return utils.Response{}, result.Error
	}

	// Fetch the updated user from the database
	var user domain.User
	err := c.DB.First(&user, req.ID).Error
	if err != nil {
		return utils.Response{}, err
	}

	err = c.updateUserInCache(user.ID, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		fmt.Println("Failed to cache")
	}

	response := utils.Response{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	return response, nil
}

func (c *authRepo) DeleteUser(ctx context.Context, id int32) error {
	var user domain.User
	err := c.DB.First(&user, id).Error
	if err != nil {
		return err
	}

	result := c.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	err = c.removeUserFromCache(id)
	if err != nil {
		fmt.Println("Failed to remove cache")
	}

	return nil
}

// CACHE FUNCTIONS

func (c *authRepo) cacheUser(user domain.User) error {
	key := fmt.Sprintf("user:%d", user.ID)
	err := c.Redis.HMSet(context.Background(), key, map[string]interface{}{
		"FirstName": user.FirstName,
		"LastName":  user.LastName,
		"Email":     user.Email,
		"Phone":     user.Phone,
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *authRepo) getUserFromCache(id int32) (utils.Response, error) {
	key := fmt.Sprintf("user:%d", id)
	val, err := c.Redis.HGetAll(context.Background(), key).Result()
	if err != nil {
		return utils.Response{}, err
	}

	user := utils.Response{
		FirstName: val["FirstName"],
		LastName:  val["LastName"],
		Email:     val["Email"],
		Phone:     val["Phone"],
	}

	return user, nil
}

func (c *authRepo) updateUserInCache(id int32, firstName, lastName, email, phone string) error {
	key := fmt.Sprintf("user:%d", id)
	err := c.Redis.HMSet(context.Background(), key, map[string]interface{}{
		"FirstName": firstName,
		"LastName":  lastName,
		"Email":     email,
		"Phone":     phone,
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *authRepo) removeUserFromCache(id int32) error {
	key := fmt.Sprintf("user:%d", id)
	err := c.Redis.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *authRepo) cacheUserResponse(user domain.User) error {
	err := c.cacheUser(user) // Use the user details from domain.User to cache
	if err != nil {
		fmt.Println("Failed to cache user:", err)
	}
	return err
}

// func (c *authRepo) cacheUserResponse(user domain.User) error {
// 	response := utils.Response{
// 		ID:        user.ID,
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 		Phone:     user.Phone,
// 	}
// 	return c.cacheUser(response)
// }
