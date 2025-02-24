package transport

import (
	"context"
	"net/http"
)

type ITransport interface {
	Http(ctx context.Context) http.Handler
	WebSkt(ctx context.Context) http.Handler
}
