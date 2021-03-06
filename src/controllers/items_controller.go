package controllers

import (
	"net/http"
	"github.com/juliandev/bookstore_oauth-go/oauth"
	"github.com/juliandev/bookstore_items-api/src/domain/items"
	"github.com/juliandev/bookstore_items-api/src/services"
	"github.com/juliandev/bookstore_items-api/src/utils/http_utils"
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"strings"
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
		// http_utils.RespondError(w, err)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respError := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.RespondError(w, respError)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid_request_body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}
