package auth

import (
	"context"
	"fmt"
	"log"
	"onlineShop/external/database"
	"onlineShop/infra/response"
	"onlineShop/internal/config"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@noobe.id", uuid.NewString()),
		Password: "mysecretpassword",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}
func TestRegister_Fail(t *testing.T) {
	t.Run("error email already used", func(t *testing.T) {
		email := fmt.Sprintf("%v@noobe.id", uuid.NewString())
		req := RegisterRequestPayload{
			Email:    email,
			Password: "mysecretpassword",
		}
		err := svc.register(context.Background(), req)
		require.Nil(t, err)

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}

func TestLogin_Success(t *testing.T) {
	email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
	pass := "mh217512"
	req := RegisterRequestPayload{
		Email:    email,
		Password: pass,
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email:    email,
		Password: pass,
	}
	token, err := svc.login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)
	log.Println(token)
}
