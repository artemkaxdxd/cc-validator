package service

import (
	"context"
	"solidgate-test/config"
	"solidgate-test/internal/controllers/http/requests"
)

type CardService struct{}

func NewCardService() CardService {
	return CardService{}
}

func (s CardService) ValidateCard(ctx context.Context, card requests.Card) (bool, config.ServiceCode, error) {
	err := card.ValidateNumber()
	if err != nil {
		return false, config.ErrToServiceCode(err), err
	}

	err = card.ValidateDate()
	return err == nil, config.ErrToServiceCode(err), err
}
