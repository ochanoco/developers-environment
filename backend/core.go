package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type OchanocoDirector = func(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool
type OchanocoModifyResponse = func(proxy *OchanocoProxy, req *http.Response, c *gin.Context) bool
type OchanocoProxyWebPage = func(proxy *OchanocoProxy, c *gin.RouterGroup)

type OchanocoProxy struct {
	Directors       []OchanocoDirector
	ModifyResponses []OchanocoModifyResponse
	ProxyWebPages   []OchanocoProxyWebPage
	Engine          *gin.Engine
	Database        *Database
}

func NewOchancoProxy(
	r *gin.Engine,
	directors []OchanocoDirector,
	modifyResponses []OchanocoModifyResponse,
	ProxyWebPages []OchanocoProxyWebPage,
	database *Database,
) OchanocoProxy {
	proxy := OchanocoProxy{}

	proxy.Directors = directors
	proxy.ModifyResponses = modifyResponses
	proxy.ProxyWebPages = ProxyWebPages
	proxy.Database = database

	proxy.Engine = r

	specialPath := r.Group("/ochanoco")
	for _, webPage := range proxy.ProxyWebPages {
		webPage(&proxy, specialPath)
	}

	r.NoRoute(func(c *gin.Context) {
		director := func(req *http.Request) {
			proxy.Director(req, c)
		}

		modifyResp := func(resp *http.Response) error {
			return proxy.ModifyResponse(resp, c)
		}

		proxy := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResp,
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return proxy
}
