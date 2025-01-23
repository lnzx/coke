package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
	"github.com/lnzx/coke/api"
	. "github.com/lnzx/coke/handler"
	"gopkg.in/ini.v1"
	"log"
	"math/big"
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
	app.Get("/api/login", Login)

	app.Use(keyauth.New(keyauth.Config{
		Validator: validate, // header:Authorization
	}))

	router := app.Group("/api")
	router.Post("/logout", Logout)
	// users
	MountUserRoutes(router)
	// preauthkeys
	router.Get("/preauthkeys", func(c fiber.Ctx) error {

		return c.SendString("login")
	})
	// nodes
	MountNodeRoutes(router)

	// SPA (Single Page Application)
	app.Get("/*", func(c fiber.Ctx) error {
		return c.SendFile("web/index.html")
	})

	str := "2"
	bigInt := new(big.Int)
	bigInt, success := bigInt.SetString(str, 10) // 第二个参数是进制，10表示十进制
	if success {
		fmt.Println("Big integer:", bigInt)
		fmt.Println("str len:", len(str))
		// 获取该大整数的位数
		bitLength := bigInt.BitLen()
		fmt.Println("大整数占用的位数:", bitLength)

		result := bigInt.Exp(bigInt, big.NewInt(4), nil)
		fmt.Println("result:", result)
	} else {
		fmt.Println("无法解析该字符串为大整数")
	}

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

func Login(c fiber.Ctx) error {
	var param = new(api.Profile)
	var err error
	if err = c.Bind().Body(param); err != nil { // <- here you receive the validation errors
		return c.JSON(fiber.ErrBadRequest)
	}
	if param.Username != _profile.Username || param.Password != _profile.Password {
		return c.JSON(fiber.ErrUnauthorized)
	}
	_token, err = genToken(param.Password)
	if err != nil {
		return c.JSON(fiber.ErrInternalServerError)
	}
	return c.JSON(fiber.Map{"token": _token})
}

func Logout(c fiber.Ctx) error {
	return c.SendString("logout")
}

func validate(_ fiber.Ctx, token string) (bool, error) {
	if token == _token {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func loadConfig() {
	_profile = new(api.Profile)
	if err := ini.MapTo(_profile, "config.ini"); err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
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
	return hex.EncodeToString(hash[:]), nil
}

// 生成一个随机的盐值
func genSalt() ([]byte, error) {
	salt := make([]byte, 16) // 16 字节的盐值
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
