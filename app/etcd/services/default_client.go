package etcdservices

import clientv3 "go.etcd.io/etcd/client/v3"

type defaultService struct {
	etcdClient *clientv3.Client
}

func NewService(etcdClient *clientv3.Client) Service {
	return &defaultService{etcdClient: etcdClient}
}
