package account

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string, email string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Email       string               `json:"email"`
	SocialID    string               `json:"social_id"`
	PhotoURL    string               `json:"photoURL"`
	LocationURL string               `json:"locationURL"`
	LastUpdated *timestamp.Timestamp `json:"last_updated"`
}

type accountService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string, email string) (*Account, error) {
	a := &Account{
		Name:  name,
		ID:    ksuid.New().String(),
		Email: email,
	}
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.ListAccounts(ctx, skip, take)
}
