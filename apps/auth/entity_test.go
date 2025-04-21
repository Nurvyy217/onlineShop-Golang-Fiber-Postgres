package auth

import (
	"log"
	"onlineShop/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthEntity(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "mh217512@gmail.com",
			Password: "mh217512",
		}
		err := authEntity.Validate()
		require.Nil(t, err)
	})
	t.Run("email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "mh217512",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})
	t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "hasbiiiii",
			Password: "mh217512",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})
	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "mh217512@gmail.com",
			Password: "",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})
	t.Run("password must have minimum 6 characters", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "mh217512@gmail.com",
			Password: "pass",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalidLength, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "mh217512@gmail.com",
			Password: "mh217512",
		}
		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)
		log.Printf("%+v\n", authEntity)
	})
}
