package transaction

import (
	"context"
	"onlineShop/external/database"
	"onlineShop/internal/config"
	"testing"

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

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU: "f8e0ecf2-b745-40ad-a66c-1d600a464315",
			Amount:     2,
			UserPublicId:      "821f61da-7b15-4cf8-984e-37ffa188db0c",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})
}
