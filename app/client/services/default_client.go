package clientservices

import (
	"crypto/md5"
	"encoding/hex"
	blogrepositories "github.com/meth-suchatchai/kz-blog-api/app/blog/repositories"
	userrepositories "github.com/meth-suchatchai/kz-blog-api/app/user/repositories"
	kzjwt "github.com/meth-suchatchai/kz-blog-api/lib/jwt"
	"github.com/meth-suchatchai/kz-blog-api/lib/taximail"
	"github.com/redis/go-redis/v9"
)

type defaultService struct {
	userRepository userrepositories.Repository
	blogRepository blogrepositories.Repository
	auth           kzjwt.AuthJWT
	mail           taximail.Client
	rdc            *redis.Client
}

func NewService(userRepository userrepositories.Repository, blogRepository blogrepositories.Repository, jwt kzjwt.AuthJWT, rdc *redis.Client) Service {
	return &defaultService{userRepository: userRepository, blogRepository: blogRepository, auth: jwt, rdc: rdc}
}

func (svc *defaultService) encryptedHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
