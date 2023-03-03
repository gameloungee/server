package handler

import (
	"net/http"

	"github.com/gameloungee/server/config"
	"github.com/gameloungee/server/internal"
	"github.com/gameloungee/server/internal/account"
	"github.com/gameloungee/server/internal/response"
	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(c *gin.Context) {
	acc, err := internal.Unmarshal[account.Account](c.Request.Body)

	if err != nil {
		response.AbortWith(http.StatusBadRequest, "Parsing error with JSON sent", err.Error(), c)
	}

	err = acc.ToDatabase()

	if err != nil && config.New().AppMode == config.PROD_MOD {
		response.AbortWith(http.StatusInternalServerError, "An error occurred when adding an object to the database", response.DETAILS_HIDDEN, c)
	} else if err != nil && config.New().AppMode == config.DEVELOP_MODE {
		response.AbortWith(http.StatusInternalServerError, "An error occurred when adding an object to the database", err.Error(), c)
	}

	response.Send(http.StatusCreated, "User successfully created", "", c)
}
