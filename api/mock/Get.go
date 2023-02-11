package mock

import (
	"github.com/aws/aws-lambda-go/events"
)

var Get = events.APIGatewayV2HTTPRequest{
	Version:               "2.0",
	RouteKey:              "$default",
	RawPath:               "/",
	RawQueryString:        "",
	Headers:               map[string]string{},
	QueryStringParameters: map[string]string{},
	PathParameters:        map[string]string{},
	RequestContext: events.APIGatewayV2HTTPRequestContext{
		HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
			Method: "GET",
		},
	},
}
