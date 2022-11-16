package controllers

import (
	"strconv"
	"time"
	"os"
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

		var validate_err []string

		for _, err := range err.(validator.ValidationErrors) {
			validate_err = append(validate_err,err.Field())
		}

		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "wrong input",		
			"validate_error": validate_err,	
		})
	}
	if (data.Password != data.RetypePassword){
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "password is mismatch",
			"validate_error": 0,			
		})
	}


	if(models.UserNameExist(data.UserName)) {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "username has already been registered",
			"validate_error": 0,
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user := models.User{
		UserName: data.UserName,
		Password: string(password),
	}

	user.CreateUser()

	return c.JSON(fiber.Map{
		"message": "success",
		"validate_error": 0,		
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

		var validate_err []string

		for _, err := range err.(validator.ValidationErrors) {
			validate_err = append(validate_err,err.Field())
		}

		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "wrong input",		
			"validate_error": validate_err,		
		})
	}
	
	var user models.User

	if (!models.UserNameExist(data.UserName)){
		c.Status(fiber.StatusNotFound)
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "user not found",
			"validate_error" : 0,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "incorrect password",
			"validate_error": 0,		
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not login",
			"validate_error": 0,		
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
		"validate_error": 0,		
	})
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
