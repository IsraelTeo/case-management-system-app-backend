package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/jwt"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
)

/* Login autentica a un usuario usando su username y password.
* Si las credenciales son correctas, genera y retorna un JWT de acceso.
* Caso contrario, devuelve un error.*/
func (s *Service) Login(ctx context.Context, request *payload.LoginRequest) (string, error) {
	user, err := s.repository.GetByUsername(ctx, request.Username)
	if err != nil {
		slog.Error("user not found or database error",
			slog.String("username", request.Username),
			slog.Any("error", err),
		)
		return "", fmt.Errorf("usuario o contraseña incorrectos")
	}

	if err := comparePassword(request.Password, user.Password); err != nil {
		return "", fmt.Errorf("usuario o contraseña incorrectos")
	}

	accessToken, err := jwt.GenerateToken(user, s.cfg)
	if err != nil {
		slog.Error("error generating JWT token",
			slog.String("username", request.Username),
			slog.Any("error", err),
		)
		return "", fmt.Errorf("error generando el token de acceso")
	}

	return accessToken, nil
}
