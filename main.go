package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/lnzx/coke/api"
	. "github.com/lnzx/coke/handler"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

var _profile *api.Profile
var _token string

func init() {
	loadConfig()
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder:     sonic.Marshal,
		JSONDecoder:     sonic.Unmarshal,
		StructValidator: &structValidator{validate: validator.New()},
	})

	// /api/login endpoint does not require authentication
	app.Post("/api/login", Login)

	router := app.Group("/api", func(c fiber.Ctx) error {
		path := c.Path()
		fmt.Println("api group path:", path)

		token := c.Get(fiber.HeaderAuthorization)
		fmt.Println("_token:", _token)
		fmt.Println("token:", token)

		return c.Next()
	})
	router.Post("/logout", Logout)

	// users
	MountUserRoutes(router)

	// nodes
	MountNodeRoutes(router)

	// preauthkeys
	router.Get("/preauthkeys", func(c fiber.Ctx) error {

		return c.SendString("login")
	})

	app.Use("/", static.New("", static.Config{
		FS: os.DirFS("web/dist"),
	}))

	app.Get("/*", func(c fiber.Ctx) error {
		return c.SendFile("web/dist/index.html")
	})

	// Start the server
	log.Fatal(app.Listen(_profile.Addr + ":" + _profile.Port))
}

func Login(c fiber.Ctx) error {
	var param = new(api.Profile)
	var err error
	if err = c.Bind().Body(param); err != nil { // <- here you receive the validation errors
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if param.Username != _profile.Username || param.Password != _profile.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	_token, err = genToken(param.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": _token})
}

func Logout(c fiber.Ctx) error {
	return c.SendString("logout")
}

func loadConfig() {
	_profile = new(api.Profile)
	if err := ini.MapTo(_profile, "config.ini"); err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", _profile)
}

type structValidator struct {
	validate *validator.Validate
}

// Validate needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func genToken(username string) (string, error) {
	// 生成一个盐值
	salt, err := genSalt()
	if err != nil {
		return "", err
	}
	// 拼接用户名、盐值，时间戳使用逗号作为分隔符
	token := []byte(fmt.Sprintf("%s,%x,%d", username, salt, time.Now().UnixNano()))
	hash := sha256.Sum256(token)
	// 返回生成的哈希值的十六进制字符串
	return "tk" + hex.EncodeToString(hash[16:]), nil
}

// 生成一个随机的盐值
func genSalt() ([]byte, error) {
	salt := make([]byte, 16) // 16 字节的盐值
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
