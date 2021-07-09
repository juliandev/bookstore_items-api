package items

import (
	"github.com/juliandev/bookstore_items-api/src/clients/mongodb"
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"errors"
	"fmt"
)

const (
	collectionName = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := mongodb.InsertDocument(i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))

	}
	i.Id = result
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	result, err := mongodb.GetDocument(i.Id)
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	result.Decode(&i)
	if i.Description == (Description{}) {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
	}
	return nil
}
