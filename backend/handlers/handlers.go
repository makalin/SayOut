package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/makalin/Sayo/auth"
	"github.com/makalin/Sayo/models"
)

type Handler struct {
	db *models.Database
}

func NewHandler(db *models.Database) *Handler {
	return &Handler{db: db}
}

// Auth middleware
func (h *Handler) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	c.Locals("user_id", claims.UserID)
	return c.Next()
}

// User handlers
func (h *Handler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Hash password
	hash, err := auth.HashPassword(user.PasswordHash)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	user.PasswordHash = hash

	// Insert user
	result, err := h.db.Exec(`
		INSERT INTO users (username, email, password_hash)
		VALUES (?, ?, ?)
	`, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)

	// Generate token
	token, err := auth.GenerateToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	err := h.db.QueryRow(`
		SELECT id, username, email, password_hash
		FROM users
		WHERE email = ?
	`, credentials.Email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if !auth.CheckPasswordHash(credentials.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := auth.GenerateToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

// Say handlers
func (h *Handler) CreateSay(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	var say models.Say
	if err := c.BodyParser(&say); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	say.UserID = userID

	result, err := h.db.Exec(`
		INSERT INTO says (user_id, content, is_private)
		VALUES (?, ?, ?)
	`, say.UserID, say.Content, say.IsPrivate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create say",
		})
	}

	id, _ := result.LastInsertId()
	say.ID = int(id)

	return c.JSON(say)
}

func (h *Handler) GetSays(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	isPrivate := c.Query("private") == "true"

	query := `
		SELECT s.*, u.username, u.avatar_url
		FROM says s
		JOIN users u ON s.user_id = u.id
		WHERE s.is_private = ?
		ORDER BY s.created_at DESC
	`

	if isPrivate {
		query = `
			SELECT s.*, u.username, u.avatar_url
			FROM says s
			JOIN users u ON s.user_id = u.id
			WHERE s.user_id = ? OR s.user_id IN (
				SELECT friend_id FROM friends WHERE user_id = ? AND status = 'accepted'
			)
			ORDER BY s.created_at DESC
		`
	}

	rows, err := h.db.Query(query, isPrivate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch says",
		})
	}
	defer rows.Close()

	var says []models.Say
	for rows.Next() {
		var say models.Say
		var user models.User
		err := rows.Scan(
			&say.ID, &say.UserID, &say.Content, &say.IsPrivate,
			&say.CreatedAt, &say.UpdatedAt,
			&user.Username, &user.AvatarURL,
		)
		if err != nil {
			continue
		}
		say.User = &user
		says = append(says, say)
	}

	return c.JSON(says)
}
