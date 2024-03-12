package v1

import (
	"context"
	"solidgate-test/config"
	"solidgate-test/internal/controllers/http/requests"
)

type Services struct {
	CardService CardService
}

type CardService interface {
	ValidateCard(ctx context.Context, body requests.Card) (bool, config.ServiceCode, error)
}
