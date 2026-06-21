package serviceauth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/repository/repoauth"
	"github.com/redis/go-redis/v9"
)


type AuthService interface {
	Loginuser(ctx context.Context, data *request.Login) (string,error)
	Signup(ctx context.Context, data *request.Register) error
	Logout(ctx context.Context, userID int, tokenID string) error
}

type ServiceRepoAuth struct {
	red *redis.Client
	jwt_secret, refresh_secret string
	repo repoauth.RepoAuth
}

func ConnectRepo(rds *redis.Client,r repoauth.RepoAuth)  AuthService {
	return &ServiceRepoAuth{
		red: rds,
		jwt_secret: os.Getenv("JWT_SECRET_KEY"),
		refresh_secret: os.Getenv("REFRESH_SECRET_KEY"),
		repo: r,
	}
}

func (s *ServiceRepoAuth) Loginuser(ctx context.Context,req *request.Login) (string,error) {
	data, err := s.repo.Login(req.Email)
	if err != nil {
		return "", err
	}


	if err := helper.DecryptPassword(data.Password, req.Password);err != nil {
		return "",err
	}

	
	token, err:= helper.GenerateToken(s.jwt_secret,data.IdUser, data.Role.CodeRole)
	if err != nil {
		return "",err
	}

	refreshtoken, err:= helper.GenerateToken(s.refresh_secret,data.IdUser, data.Role.CodeRole)
	if err != nil {
		return "",err
	}

	s.red.Set(ctx, "refresh_token:"+strconv.Itoa(int(data.IdUser)), refreshtoken, 7*24*time.Hour)



	return token,nil
}

func (s *ServiceRepoAuth) Signup(ctx context.Context, data *request.Register) error {
	data.Password = helper.HashingPassword(data.Password)
	data.Model.CreatedAt = time.Now().Local()

	if err := s.repo.Register(data); err != nil {
		return err
	}

	return nil
}

func (s *ServiceRepoAuth) Logout(ctx context.Context, userID int, tokenID string) error {
	if err := s.red.Del(ctx, "refresh_token:"+strconv.Itoa(userID)).Err(); err != nil {
		return err
	}

	if err := s.red.Set(ctx, "blacklist_token:"+tokenID, true, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}