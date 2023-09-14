package clientservices

import (
	"crypto/md5"
	"encoding/hex"
	blogrepositories "github.com/kuroshibaz/app/blog/repositories"
	userrepositories "github.com/kuroshibaz/app/user/repositories"
	kzjwt "github.com/kuroshibaz/lib/jwt"
	"github.com/kuroshibaz/lib/taximail"
	"github.com/redis/go-redis/v9"
)

type defaultService struct {
	userRepository userrepositories.Repository
	blogRepository blogrepositories.Repository
	auth           kzjwt.AuthJWT
	mail           taximail.Client
	rdc            *redis.Client
}

func NewService(userRepository userrepositories.Repository, blogRepository blogrepositories.Repository, jwt kzjwt.AuthJWT, mail taximail.Client, rdc *redis.Client) Service {
	return &defaultService{userRepository: userRepository, blogRepository: blogRepository, auth: jwt, mail: mail, rdc: rdc}
}

func (svc *defaultService) encryptedHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
