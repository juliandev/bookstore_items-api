package controllers

import (
	"net/http"
	"github.com/juliandev/bookstore_oauth-go/oauth"
	"github.com/juliandev/bookstore_items-api/domain/items"
	"github.com/juliandev/bookstore_items-api/services"
	"fmt"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	 Create(w http.ResponseWriter, r *http.Request)
	 Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO
		return
	}

	item := items.Item {
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
	}

	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
