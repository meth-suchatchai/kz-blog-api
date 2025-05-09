package command

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/meth-suchatchai/kz-blog-api/config"
	"github.com/meth-suchatchai/kz-blog-api/lib/gormdb"
	kzjwt "github.com/meth-suchatchai/kz-blog-api/lib/jwt"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzline"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzobjectstorage"
	"github.com/meth-suchatchai/kz-blog-api/lib/taximail"
	"github.com/meth-suchatchai/kz-blog-api/lib/totp"
	"github.com/meth-suchatchai/kz-blog-api/router"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server(cfg *config.Env) {
	/* Initialize Database */
	db, err := gormdb.ConnectSQL(&cfg.Database)
	if err != nil {
		log.Fatalf("error connect SQL: %v", err)
	}

	err = db.Migrate()
	//db.Seed()
	if err != nil {
		log.Fatalf("error migrate: %v", err)
	}

	rc := resty.New()
	taxiMailService, err := taximail.New(&taximail.Provide{
		ApiKey:      cfg.TaxiMail.ApiKey,
		SecretKey:   cfg.TaxiMail.SecretKey,
		URL:         cfg.TaxiMail.URL,
		SMSTemplate: cfg.TaxiMail.SMSTemplate,
	}, rc)

	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{cfg.ETCD.Hostname},
		DialTimeout: time.Duration(cfg.ETCD.Timeout) * time.Second,
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
	var isMinioConnected = false
	var bucketCli kzobjectstorage.StorageBucket
	if msc != nil {
		isMinioConnected = true
	}

	go func() {
		for {
			if isMinioConnected {
				break
			}
			time.Sleep(time.Minute)
			bucketCli, bucketErr := kzobjectstorage.NewSelectBucket(cfg.Storage.Bucket, cfg.Storage.Endpoint, cfg.Storage.Region, msc.Minio())
			if bucketErr != nil {
				log.Fatalf("error choose bucket size: %v", bucketErr)
			}
			log.Println(bucketCli)
		}
	}()

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

	lineCli := kzline.NewLineNotification(cfg.Line.BotApi, cfg.Line.LineApi, cfg.Line.AccessToken, rc)
	//log.Println(lineCli.PushMessage(kzline.PushMessageRequest{
	//	Message:              "Test",
	//	NotificationDisabled: false,
	//}))

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
		LineService:     lineCli,
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
			log.Fatalf("can't start command %v", err)
			cancel()
		}
	}()
	<-ctx.Done()
}
