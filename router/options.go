package router

import (
	"github.com/go-resty/resty/v2"
	"github.com/meth-suchatchai/kz-blog-api/config"
	"github.com/meth-suchatchai/kz-blog-api/lib/gormdb"
	kzjwt "github.com/meth-suchatchai/kz-blog-api/lib/jwt"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzobjectstorage"
	"github.com/meth-suchatchai/kz-blog-api/lib/totp"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Options struct {
	Env            *config.Env
	Log            *logrus.Logger
	EtcdClient     *clientv3.Client
	Rc             *resty.Client
	Redis          *redis.Client
	Db             gormdb.Client
	TOtp           totp.Client
	Jwt            kzjwt.AuthJWT
	StorageService kzobjectstorage.StorageBucket
	//LineService     kzline.LineNotification
	//TaximailService taximail.Client
}
