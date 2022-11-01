package controllers

import (
	"fmt"
	"strconv"
	"time"
	"os"

	"blog-cms/database"
	"blog-cms/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"

)

var validate *validator.Validate

func Register(c *fiber.Ctx) error {
	// var data *map[string]string

	data := new(struct {
		UserName string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required,min=8"`
		RetypePassword string `form:"retype-password"`

	})

	if err := c.BodyParser(data); err != nil {
		return err
	}

	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var message_err []string

		for _, err := range err.(validator.ValidationErrors) {
			message_err = append(message_err,err.Field())
		}

		return c.JSON(fiber.Map{
			"message": message_err,			
		})
	}
	if (data.Password != data.RetypePassword){
		return c.JSON(fiber.Map{
			"message": "password is mismatch",			
		})
	}

	var count int64

	database.DBGorm.Model(&models.User{}).Where("user_name = ?", data.UserName ).Count(&count)

	if(count>0) {
		return c.JSON(fiber.Map{
			"message": "username has already been registered",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user := models.User{
		UserName: data.UserName,
		Password: string(password),
	}

	database.DBGorm.Create(&user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login(c *fiber.Ctx) error {

	SecretKey := os.Getenv("SecretKey")


	var data struct {
		UserName string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	validate = validator.New()
	err := validate.Struct(&data)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var message_err []string

		for _, err := range err.(validator.ValidationErrors) {
			message_err = append(message_err,err.Field())
		}

		// from here you can create your own error messages in whatever language you wish
		return c.JSON(fiber.Map{
			"message": message_err,			
		})
	}
	
	var user models.User

	database.DBGorm.First(&user, "user_name = ?", data.UserName)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	SecretKey := os.Getenv("SecretKey")

	
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DBGorm.Where("id = ?", claims.Issuer).First(&user)

	fmt.Println(user.UserName)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
