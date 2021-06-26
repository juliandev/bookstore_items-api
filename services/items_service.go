package services

import (
	"github.com/juliandev/bookstore_items-api/domain/items"
	"github.com/juliandev/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create (items.Item) (*items.Item, *rest_errors.RestErr)
	Get (string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct {}

func (s *itemsService) Create(items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewNotFoundError("implement me")
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewNotFoundError("implement me")
}
