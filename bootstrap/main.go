package main

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/kuroshibaz/config"
	"github.com/kuroshibaz/lib/gormdb"
	kzjwt "github.com/kuroshibaz/lib/jwt"
	"github.com/kuroshibaz/lib/kzobjectstorage"
	"github.com/kuroshibaz/lib/taximail"
	"github.com/kuroshibaz/lib/totp"
	"github.com/kuroshibaz/router"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := loadConfig()
	/* Initialize Database */
	db, err := gormdb.ConnectSQL(&cfg.Database)
	if err != nil {
		log.Fatalf("error connect SQL: %v", err)
	}

	err = db.Migrate()
	//db.Seed()
	if err != nil {
		log.Fatal("migrate: ", err)
	}

	rc := resty.New()
	taxiMailService, err := taximail.New(&taximail.Provide{
		ApiKey:      cfg.TaxiMail.ApiKey,
		SecretKey:   cfg.TaxiMail.SecretKey,
		URL:         cfg.TaxiMail.URL,
		SMSTemplate: cfg.TaxiMail.SMSTemplate,
	}, rc)

	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("error etcd: %v", err)
	}
	log.Println(etcdCli.Endpoints())
	defer etcdCli.Close()

	msc, err := kzobjectstorage.NewClient(&kzobjectstorage.Options{
		Endpoint:        cfg.Storage.Endpoint,
		AccessKeyId:     cfg.Storage.AccessKeyId,
		SecretAccessKey: cfg.Storage.SecretKey,
		UseSSL:          false,
		Region:          cfg.Storage.Region,
	})
	if err != nil {
		log.Fatalf("error minio: %v", err)
	}

	bucketCli, err := kzobjectstorage.NewSelectBucket(cfg.Storage.Bucket, cfg.Storage.Endpoint, cfg.Storage.Region, msc.Minio())
	if err != nil {
		log.Fatalf("error choose bucket size: %v", err)
	}
	log.Println(bucketCli)

	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
	})
	err = redisCli.Ping(context.TODO()).Err()
	if err != nil {
		log.Fatalf("redis connection failed: %v", err)
	}

	totpcli := totp.New(totp.Config{AppName: cfg.Server.ApplicationName})

	kzjwt := kzjwt.New(&config.JWT{
		Secret:        cfg.JWT.Secret,
		Issuer:        cfg.JWT.Issuer,
		Domain:        cfg.JWT.Domain,
		Expire:        cfg.JWT.Expire,
		RefreshExpire: cfg.JWT.RefreshExpire,
	})

	app := router.NewRouter(&router.Options{
		Env:             cfg,
		Db:              db,
		Rc:              rc,
		TaximailService: taxiMailService,
		EtcdClient:      etcdCli,
		TOtp:            totpcli,
		Jwt:             kzjwt,
		Redis:           redisCli,
		StorageService:  bucketCli,
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		log.Println("Graceful Shutdown...")
		cancel()
	}()

	/* Running Application */
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
			log.Fatalf("can't start application %v", err)
			cancel()
		}
	}()
	<-ctx.Done()
}

/* loadConfig read/map to environment config */
func loadConfig() *config.Env {
	cfg, err := config.ReadConfig("config")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return cfg
}
