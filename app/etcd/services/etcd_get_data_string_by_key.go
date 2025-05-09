package etcdservices

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (s *defaultService) GetDataStringByKey(key string) (string, *fiber.Error) {
	res, err := s.etcdClient.Get(context.TODO(), key)
	if err != nil {
		return "", errors.NewDefaultFiberError(err)
	}

	var value []byte
	for _, v := range res.Kvs {
		value = v.Value
	}

	return string(value), nil
}
