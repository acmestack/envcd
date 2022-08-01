package util

import (
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// BuildContext build plugin context
//  @param params params
//  @return *context.Context context
//  @return error error
func BuildContext(ginCtx *gin.Context) (*context.Context, error) {
	ctx := &context.Context{
		Uri:         ginCtx.Request.RequestURI,
		Method:      ginCtx.Request.Method,
		Headers:     buildContextHeaders(ginCtx),
		ContentType: ginCtx.ContentType(),
		Cookies:     buildContextCookies(ginCtx),
		Body:        buildRequestBody(ginCtx),
		HttpRequest: ginCtx.Request,
	}
	if ctx != nil {
		return ctx, nil
	}
	return nil, nil
}

// ParseContext parse context to envcd data
//  @param ctx context
//  @return *data.EnvcdData data
//  @return error error
func ParseContext(ctx *context.Context) (*data.EnvcdData, error) {
	// TODO parse context to envcd data
	return nil, nil
}

// buildContextHeaders build plugin context headers
//  @param ginCtx gin context
//  @return map[string]interface{} ret
func buildContextHeaders(ginCtx *gin.Context) map[string]interface{} {
	maps := make(map[string]interface{})
	for k, v := range ginCtx.Request.Header {
		maps[k] = v
	}
	return maps
}

// buildContextCookies build plugin context cookies
//  @param ginCtx gin context
//  @return map[string]interface{} ret
func buildContextCookies(ginCtx *gin.Context) map[string]interface{} {
	maps := make(map[string]interface{})
	for k, v := range ginCtx.Request.Cookies() {
		maps[string(rune(k))] = v
	}
	return maps
}

// buildRequestBody build request body
//  @param ginCtx gin context
//  @return string request body
func buildRequestBody(ginCtx *gin.Context) string {
	all, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return ""
	}
	return string(all)
}
