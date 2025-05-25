package router

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v5"
	bloghandlers "github.com/meth-suchatchai/kz-blog-api/app/blog/handlers"
	blogrepositories "github.com/meth-suchatchai/kz-blog-api/app/blog/repositories"
	blogservices "github.com/meth-suchatchai/kz-blog-api/app/blog/services"
	calculateservices "github.com/meth-suchatchai/kz-blog-api/app/calculate/services"
	clienthandlers "github.com/meth-suchatchai/kz-blog-api/app/client/handlers"
	clientservices "github.com/meth-suchatchai/kz-blog-api/app/client/services"
	etcdservices "github.com/meth-suchatchai/kz-blog-api/app/etcd/services"
	filerepositories "github.com/meth-suchatchai/kz-blog-api/app/file/repositories"
	fileservices "github.com/meth-suchatchai/kz-blog-api/app/file/services"
	rphandlers "github.com/meth-suchatchai/kz-blog-api/app/role_permission/handlers"
	rprepositories "github.com/meth-suchatchai/kz-blog-api/app/role_permission/repositories"
	rpservices "github.com/meth-suchatchai/kz-blog-api/app/role_permission/services"
	scenehandlers "github.com/meth-suchatchai/kz-blog-api/app/scene/handlers"
	scenerepositories "github.com/meth-suchatchai/kz-blog-api/app/scene/repositories"
	sceneservices "github.com/meth-suchatchai/kz-blog-api/app/scene/services"
	userhandlers "github.com/meth-suchatchai/kz-blog-api/app/user/handlers"
	userrepositories "github.com/meth-suchatchai/kz-blog-api/app/user/repositories"
	userservices "github.com/meth-suchatchai/kz-blog-api/app/user/services"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	"github.com/meth-suchatchai/kz-blog-api/lib/validator"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

func NewRouter(opts *Options) *fiber.App {
	if opts == nil {
		log.Fatal("can't load config")
	}

	cv := validator.CustomValidator{Validator: validator.Validate}
	viewEngine := html.New("views", ".html")

	/* Repositories */
	userRepo := userrepositories.NewRepository(opts.Db)
	roleRepo := rprepositories.NewRepository(opts.Db)
	blogRepo := blogrepositories.NewRepository(opts.Db)
	fileRepo := filerepositories.NewRepository(opts.StorageService)
	sceneRepo := scenerepositories.NewRepository(opts.Db)

	/* Services */
	userService := userservices.NewService(userRepo, opts.TOtp)
	roleService := rpservices.NewService(roleRepo)
	blogService := blogservices.NewService(blogRepo)
	sceneService := sceneservices.NewService(sceneRepo)

	etcdService := etcdservices.NewService(opts.EtcdClient)
	fileService := fileservices.NewService(fileRepo)
	clientService := clientservices.NewService(userRepo, blogRepo, opts.Jwt, opts.Redis)

	userHandler := userhandlers.NewHandler(cv, userService)
	roleHandler := rphandlers.NewHandler(cv, roleService)
	cliHandler := clienthandlers.NewHandler(cv, userService, blogService, etcdService, clientService, opts.Jwt)
	blogHandler := bloghandlers.NewHandler(cv, blogService, fileService)
	sceneHandler := scenehandlers.NewHandler(cv, sceneService)

	pemMiddleware := NewPermission(opts.Db)

	//appz := kuroctxfiber.New(kuroctxfiber.Config{
	//	Name: "kz-blog-api",
	//})

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
		Views:       viewEngine,
		ViewsLayout: "layouts/main",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(func(ctx *fiber.Ctx) error {
		token := ctx.Get("KZ-API", "")
		//log.Printf("token: ", token, opts.Env.Server.AccessToken)
		if token == "" || token != opts.Env.Server.AccessToken {
			return ctx.Status(fiber.StatusForbidden).SendString("who are youuuuu!!!!")
		}
		return ctx.Next()
	})

	app.Use(func(ctx *fiber.Ctx) error {
		isMaintenance := false
		_ = etcdService.GetDataByKey(constant.MaintenanceStage, &isMaintenance)
		if isMaintenance {
			return ctx.Status(fiber.StatusForbidden).SendString("is undermaintain!")
		}
		return ctx.Next()
	})

	app.Static("/", "./public")
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{})
	})

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
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("success")
	})
	app.Get("/metrics", monitor.New(), pemMiddleware.CheckPermission("AUDIT"))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	client := v1.Group("/client")
	{
		client.Post("/login", cliHandler.Login)
		client.Post("/register/admin", cliHandler.RegisterAdmin)
		client.Post("/verify/otp", cliHandler.VerifyOTP)
		client.Get("/blogs", cliHandler.ListBlog)
		client.Get("/blog/:slug", cliHandler.GetBlog)
		client.Put("/blog/views/:slug", cliHandler.UpdateViewBlog)
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
			SuccessHandler: func(ctx *fiber.Ctx) error {
				//log.Println("userContext: ", ctx.UserContext())
				msg := fiber.NewError(401, "Invalid or expired access token")
				token, ok := ctx.Locals("user").(*jwt.Token)
				if !ok {
					log.Printf("Claims token failed")
					return ctx.Status(fiber.StatusUnauthorized).JSON(msg)
				}
				uid, ok := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
				if !ok {
					log.Printf("Claims token failed")
					return ctx.Status(fiber.StatusUnauthorized).JSON(msg)
				}

				if userId, err := strconv.Atoi(uid); err == nil {
					_, authErr := userRepo.GetUserAuthenticationByUserId(uint(userId), token.Raw)
					if authErr != nil {
						log.Printf("GetUserAuthenticationByUserId error: %v", authErr)
						return ctx.Status(fiber.StatusUnauthorized).JSON(msg)
					}
				} else {
					log.Printf("Claims token failed")
					return ctx.Status(fiber.StatusUnauthorized).JSON(msg)
				}

				//Check Latest Token
				return ctx.Next()
			},
		}))
	}

	user := crm.Group("/user")
	{
		user.Get("/profile", userHandler.Profile)
		user.Put("/enabled_otp", userHandler.OTPEnabled)
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

	scene := crm.Group("/scene")
	{
		scene.Post("", pemMiddleware.CheckPermission("CREATE_SCENE"), sceneHandler.CreateScene)
		scene.Put(":id/status", pemMiddleware.CheckPermission("UPDATE_SCENE"), sceneHandler.UpdateStatusScene)
	}

	c := calculateservices.NewService()
	v1.Post("/calc", func(ctx *fiber.Ctx) error {
		//delivery := make(chan kafka.Event, 10000)
		//err = p.Produce(&kafka.Message{
		//	TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		//	Value:          []byte(value)},
		//	delivery_chan,
		//)
		c.SampleService()
		return nil
	})

	return app
}
