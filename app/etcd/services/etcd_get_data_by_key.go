package etcdservices

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (s *defaultService) GetDataByKey(key string, output interface{}) *fiber.Error {
	res, err := s.etcdClient.Get(context.TODO(), key)
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}

	var value []byte
	for _, v := range res.Kvs {
		value = v.Value
	}

	err = json.Unmarshal(value, &output)
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}

	return nil
}
