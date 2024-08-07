package dictionary

import (
	"httpHandler/models"
	"net/http"
)

type ErrorsClient struct {
	Internal models.Response
	BadArgs  models.Response
}

var errClient = ErrorsClient{
	Internal: models.Response{
		HttpCode: http.StatusInternalServerError,
		Body: models.ErrClient{
			Code:    101,
			Message: "internal server error",
		},
	},

	BadArgs: models.Response{
		HttpCode: http.StatusBadRequest,
		Body: models.ErrClient{
			Code:    102,
			Message: "bad arguments",
		},
	},
}
