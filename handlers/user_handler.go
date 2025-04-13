package handlers

import (
	"time"

	"github.com/aryanpnd/Tasknest/db"
	"github.com/aryanpnd/Tasknest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Hash password before storing
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not hash password"})
	}
	user.Password = string(hash)

	if result := db.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.Redirect("/login", 302)
}

func ShowRegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var foundUser models.User
	if result := db.DB.Where("username = ?", user.Username).First(&foundUser); result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Compare password hashes
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create JWT token
	claims := jwt.MapClaims{
		"username": foundUser.Username,
		"user_id":  foundUser.ID, // Store user ID in the token
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	// ✅ Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		HTTPOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
	})

	// return c.JSON(fiber.Map{
	// 	"token":   signedToken,
	// 	"message": "Login successful",
	// })

	return c.Redirect("/", fiber.StatusSeeOther)
}

func ShowLoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
