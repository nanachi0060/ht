package routing

import (
	"httpHandler/controllers"
	"httpHandler/initialization"
	"httpHandler/libs"
	"httpHandler/models"
	"net/http"
)

func Route(headers, body []byte,
	method, path string,
	agents initialization.Agents,
	config libs.Configuration) models.Response {

	switch path {

	case "/t1":
		controllers.T1()

	case "/t2":
		controllers.T2()

	case "/t3":
		controllers.T3()

	default:
		return models.Response{
			HttpCode: http.StatusNotFound,
			Body:     nil,
			Headers:  nil,
		}
	}

	return models.Response{}
}
