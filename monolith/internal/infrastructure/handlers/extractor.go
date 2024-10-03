package handlers

import (
	"net/http"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/middleware"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
)

func ExtractUser(r *http.Request) (cryptography.UserFromClaims, error) {
	ctx := r.Context()
	u, ok := ctx.Value(middleware.UserRequest{}).(cryptography.UserClaims)
	if !ok {
		return cryptography.UserFromClaims{}, errors.TokenExtractUserError
	}
	userID, err := strconv.Atoi(u.ID)
	if err != nil {
		return cryptography.UserFromClaims{}, errors.TokenExtractUserError
	}
	role, err := strconv.Atoi(u.Role)
	if err != nil {
		return cryptography.UserFromClaims{}, errors.TokenExtractUserError
	}

	return cryptography.UserFromClaims{
		ID:     userID,
		Role:   role,
		Groups: nil,
		Layers: nil,
	}, nil
}
