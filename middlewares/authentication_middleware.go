package middlewares

import (
	"fmt"
	"personal-growth/db/models"
	"personal-growth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		tokenStr := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("TOKEN_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid access token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid claims"})
		}

		rdb := handlers.NewRedis()
		rfToken := c.Cookies("refresh_token")
		val, _ := rdb.GetVal(fmt.Sprintf("actoken_%s_%s", claims["sub"], rfToken[len(rfToken)-6:]))

		if val != tokenStr {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid access token"})
		}

		// Gắn user_id vào context để dùng trong handler
		c.Locals("user_id", claims["sub"])

		return c.Next()
	}
}

func GetProfileHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(string)

		var user *models.User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// Gắn user_id vào context để dùng trong handler
		c.Locals("user", user)
		return c.Next()
	}
}
