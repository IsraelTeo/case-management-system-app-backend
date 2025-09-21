package jwt

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GenerateToken genera un JWT firmado con los datos del usuario.
func GenerateToken(u *domain.User, cfg *config.JWT) (string, error) {
	// expAt es el tiempo de expiración del token
	expAt := time.Now().Add(time.Second * time.Duration(cfg.ExpirationTime)).Unix()

	// SecretKey es la clave secreta para firmar el token
	secretKey := cfg.SecretKey

	// iat es la hora en que se emitió el token
	iat := time.Now().Unix()

	// Crea los claims personalizados
	claims := Claims{
		UserID: u.ID,
		Email:  u.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  iat,
			ExpiresAt: expAt,
		},
	}

	// Crea el token con los claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token con la clave secreta
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		slog.Error(
			"Could not sign JWT token",
			slog.String("userID", u.ID.String()),
			slog.Any("error", err),
		)

		return "", fmt.Errorf("could not sign JWT token: %w", err)
	}

	// Devuelve el token firmado
	return tokenString, nil
}

// ExtractUserID obtiene el userID desde el token
func ExtractUserID(c echo.Context, cfg *config.JWT) (uuid.UUID, error) {
	authHeader := getAuthorizationHeader(c)
	if authHeader == "" {
		return uuid.Nil, fmt.Errorf("could not get authorization header")
	}

	tokenStr := extractBearerToken(authHeader)
	if tokenStr == "" {
		return uuid.Nil, fmt.Errorf("could not extract bearer token")
	}

	claims, err := validateToken(tokenStr, cfg)
	if err != nil {
		return uuid.Nil, fmt.Errorf("could not validate token: %w", err)
	}

	// Extrae el userID de los claims
	if claims.UserID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("user id missing in token")
	}

	return claims.UserID, nil
}

// getAuthorizationHeader obtiene el encabezado de autorización de la solicitud.
func getAuthorizationHeader(c echo.Context) string {
	// Obtiene el encabezado "Authorization
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	return authHeader
}

// tokenStr := authHeader[len("Bearer "):]

// extractBearerToken extrae el token Bearer del encabezado de autorización.
func extractBearerToken(authHeader string) string {
	const prefix = "Bearer "
	// Verifica que el encabezado comience con "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		return ""
	}

	// Extrae y devuelve el token después de "Bearer "
	return authHeader[len(prefix):]
}

// validateToken valida el token y devuelve el token si es válido.
func validateToken(tokenStr string, cfg *config.JWT) (*Claims, error) {
	claims := &Claims{}

	// ParseWithClaims permite parsear directo a tu struct
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return validateMethodAndGetSecret(token, cfg)
	})
	if err != nil {
		slog.Error("token parsing failed", slog.Any("error", err))
		return nil, fmt.Errorf("invalid token") // mensaje genérico al cliente
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// validateMethodAndGetSecret asegura que el token use el algoritmo correcto y devuelve la clave para verificarlo.
func validateMethodAndGetSecret(token *jwt.Token, cfg *config.JWT) ([]byte, error) {
	// token.Method es el agoritmo de firma con el que se firmó el token
	// .(*jwt.SigningMethodHMAC) intenta convertirlo al tipo esperado: HMAC
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	// Si la conversión no es exitosa, el método no es válido
	// Sino el token usa HMAC (válido).
	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	// Devuelve la clave secreta (en bytes) al parser de JWT para verificar el token
	return []byte(cfg.SecretKey), nil
}
