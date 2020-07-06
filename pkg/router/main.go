package router

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Environments -
type Environments struct {
	Port string
}

type (
	// EndPoint -
	EndPoint struct {
		Name    string
		Method  string
		Handler func(ctx *Context) (int, *Response)
		Group   *gin.RouterGroup
	}

	// Response -
	Response struct {
		Status  string      `json:"status,omitempty"` // OK, Error, Warning
		Data    interface{} `json:"data,omitempty"`
		Error   string      `json:"error,omitempty"`
		Message string      `json:"message,omitempty"`
	}

	// Context -
	Context struct {
		Params    map[string]interface{}
		ExtraBody interface{}
		Body      []byte
		Headers   interface{}
		Queries   map[string]interface{}
	}
)

// Setup -
func Setup() *gin.Engine {
	rout := gin.Default()
	return rout
}

// EnableHandlers -
func EnableHandlers(p *EndPoint) {
	switch p.Method {
	case "GET":
		p.Group.GET(p.Name, func(c *gin.Context) { interceptor(c, p) })
	case "POST":
		p.Group.POST(p.Name, func(c *gin.Context) { interceptor(c, p) })
	case "PUT":
		p.Group.PUT(p.Name, func(c *gin.Context) { interceptor(c, p) })
	case "DELETE":
		p.Group.DELETE(p.Name, func(c *gin.Context) { interceptor(c, p) })
	}
}

func interceptor(c *gin.Context, p *EndPoint) {
	setHeaders(c, p.Method)

	ctx := &Context{}
	ctx.Params = make(map[string]interface{})
	ctx.Queries = make(map[string]interface{})

	for _, v := range c.Params {
		ctx.Params[v.Key] = v.Value
	}

	for k, v := range c.Request.URL.Query() {
		ctx.Queries[k] = v[0]
	}

	b, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	ctx.Body = b
	c.JSON(p.Handler(ctx))
}

func setHeaders(c *gin.Context, m string) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, "+m)
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Authorization")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}

// NewResposeError -
func NewResposeError(err string, errCode string) *Response {
	return &Response{Status: "error", Message: err, Error: errCode}
}

// NewResponseSuccess -
func NewResponseSuccess(data interface{}) *Response {
	return &Response{Status: "OK", Data: data}
}

// NewResponseSuccessWithMessage -
func NewResponseSuccessWithMessage(data interface{}, msg string) *Response {
	return &Response{Status: "OK", Data: data, Message: msg}
}
