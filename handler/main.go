package handler

import (
	"github.com/valyala/fasthttp"
	"httpHandler/initialization"
	"httpHandler/libs"
	"httpHandler/routing"
)

func Run(servicePort string, agents initialization.Agents, config libs.Configuration) error {
	handler := func(ctx *fasthttp.RequestCtx) {
		response := routing.Route(ctx.Response.Header.Header(),
			ctx.Response.Body(),
			string(ctx.Method()),
			string(ctx.Path()),
			agents,
			config)

		ctx.SetStatusCode(response.HttpCode)

		for header := range response.Headers {
			ctx.Response.Header.Add(header, response.Headers[header])
		}

		ctx.Response.SetBody(response.Body)
	}

	err := fasthttp.ListenAndServe("[::]:"+servicePort, handler)
	if err != nil {
		return err
	}

	return nil
}
