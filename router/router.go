package router

import (
	"github.com/go-resty/resty/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	bloghandlers "github.com/kuroshibaz/app/blog/handlers"
	blogrepositories "github.com/kuroshibaz/app/blog/repositories"
	blogservices "github.com/kuroshibaz/app/blog/services"
	clienthandlers "github.com/kuroshibaz/app/client/handlers"
	clientservices "github.com/kuroshibaz/app/client/services"
	etcdservices "github.com/kuroshibaz/app/etcd/services"
	filerepositories "github.com/kuroshibaz/app/file/repositories"
	fileservices "github.com/kuroshibaz/app/file/services"
	rphandlers "github.com/kuroshibaz/app/role_permission/handlers"
	rprepositories "github.com/kuroshibaz/app/role_permission/repositories"
	rpservices "github.com/kuroshibaz/app/role_permission/services"
	userhandlers "github.com/kuroshibaz/app/user/handlers"
	userrepositories "github.com/kuroshibaz/app/user/repositories"
	userservices "github.com/kuroshibaz/app/user/services"
	"github.com/kuroshibaz/config"
	"github.com/kuroshibaz/lib/gormdb"
	kzjwt "github.com/kuroshibaz/lib/jwt"
	"github.com/kuroshibaz/lib/kzobjectstorage"
	"github.com/kuroshibaz/lib/taximail"
	"github.com/kuroshibaz/lib/totp"
	"github.com/kuroshibaz/lib/validator"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

type Options struct {
	Env             *config.Env
	Db              *gormdb.DB
	Rc              *resty.Client
	TaximailService taximail.Client
	EtcdClient      *clientv3.Client
	TOtp            totp.Client
	Jwt             kzjwt.AuthJWT
	Redis           *redis.Client
	StorageService  kzobjectstorage.StorageBucket
}

func NewRouter(opts *Options) *fiber.App {
	if opts == nil {
		log.Fatal("can't load config")
	}

	cv := validator.CustomValidator{Validator: validator.Validate}

	app := fiber.New(fiber.Config{
		AppName: "kz-blog",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).JSON(e)
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	/* Repositories */
	userRepo := userrepositories.NewRepository(opts.Db)
	roleRepo := rprepositories.NewRepository(opts.Db)
	blogRepo := blogrepositories.NewRepository(opts.Db)
	fileRepo := filerepositories.NewRepository(opts.StorageService)

	/* Services */
	userService := userservices.NewService(userRepo)
	roleService := rpservices.NewService(roleRepo)
	blogService := blogservices.NewService(blogRepo)

	etcdService := etcdservices.NewService(opts.EtcdClient)
	fileService := fileservices.NewService(fileRepo)
	clientService := clientservices.NewService(userRepo, blogRepo, opts.Jwt, opts.TaximailService, opts.Redis)

	userHandler := userhandlers.NewHandler(cv, userService)
	roleHandler := rphandlers.NewHandler(cv, roleService)
	cliHandler := clienthandlers.NewHandler(cv, userService, blogService, etcdService, clientService, opts.Jwt)
	blogHandler := bloghandlers.NewHandler(cv, blogService, fileService)

	pemMiddleware := NewPermission(opts.Db)

	//app.Post("/upload-test", func(ctx *fiber.Ctx) error {
	//	form, err := ctx.MultipartForm()
	//	if err != nil {
	//		return err
	//	}
	//
	//	files := form.File[constant.FormFileKey]
	//
	//	var paths []string
	//	for _, file := range files {
	//
	//		path, vErr := fileService.FileUpload(file, "blog")
	//		if vErr != nil {
	//			return vErr
	//		}
	//		paths = append(paths, path)
	//	}
	//	return ctx.SendString("File upload successfully " + strings.Join(paths, ","))
	//})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	app.Static("/public/images", "./public/images")

	client := v1.Group("/client")
	{
		client.Post("/login", cliHandler.Login)
		client.Post("/register/admin", cliHandler.RegisterAdmin)
		client.Post("/verify/otp", cliHandler.VerifyOTP)
		client.Get("/blogs", cliHandler.ListBlog)
		client.Get("/blog/:slug", cliHandler.GetBlog)
		client.Get("/blog/views/:slug", cliHandler.UpdateViewBlog)
	}

	crm := v1.Group("/crm")
	{
		crm.Use(jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{
				Key: []byte(opts.Env.JWT.Secret),
			},
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				log.Println(err)
				vErr := fiber.NewError(401, "Invalid or expired access token")
				return ctx.Status(fiber.StatusUnauthorized).JSON(vErr)
			},
		}))
	}

	user := crm.Group("/user")
	{
		user.Get("/profile", userHandler.Profile)
	}

	role := crm.Group("/role")
	{
		role.Get("/list_permission", pemMiddleware.CheckPermission("READ_ROLE_PERMISSION"), roleHandler.ListRolePermission)
		role.Post("", pemMiddleware.CheckPermission("CREATE_ROLE"), roleHandler.CreateRole)
		//role.Put("/user/assign/:id")
	}

	pem := crm.Group("/permission")
	{
		pem.Post("", pemMiddleware.CheckPermission("CREATE_PERMISSION"), roleHandler.CreatePermission)
	}

	blog := crm.Group("/blog")
	{
		blog.Get("", pemMiddleware.CheckPermission("READ_BLOG"), blogHandler.ListBlog)
		blog.Get(":id", pemMiddleware.CheckPermission("READ_BLOG"), blogHandler.GetBlog)
		blog.Post("", pemMiddleware.CheckPermission("CREATE_BLOG"), blogHandler.CreateBlog)
		blog.Put(":id", pemMiddleware.CheckPermission("UPDATE_BLOG"), blogHandler.UpdateBlog)
		blog.Delete(":id", pemMiddleware.CheckPermission("DELETE_BLOG"), blogHandler.DeleteBlog)
	}

	return app
}
