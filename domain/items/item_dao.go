package items

import (
	"github.com/juliandev/bookstore_items-api/clients/mongodb"
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"errors"
	"fmt"
)

const (
	collectionName = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := mongodb.InsertDocument(i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))

	}
	i.Id = result
	return nil
}
