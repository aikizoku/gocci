package firebaseauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/unrolled/render"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// ContextKey ... ContextKeyの型定義
type ContextKey string

// UserIDContextKey ... UserIDのContextKey
const UserIDContextKey ContextKey = "user_id"

// ClaimsContextKey ... ClaimsのContextKey
const ClaimsContextKey ContextKey = "claims"

// Middleware ... JSONRPC2に準拠したミドルウェア
type Middleware struct {
	Svc Service
}

// Handle ... Firebase認証をする
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		userID, claims, err := m.Svc.Authentication(ctx, r)
		if err != nil {
			m.renderError(ctx, w, http.StatusForbidden, "authentication: "+err.Error())
			return
		}

		rctx := r.Context()
		rctx = context.WithValue(rctx, UserIDContextKey, userID)
		rctx = context.WithValue(rctx, ClaimsContextKey, claims)

		log.Debugf(ctx, "UserID: %s", userID)
		log.Debugf(ctx, "Claims: %v", claims)

		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

func (m *Middleware) renderError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Warningf(ctx, msg)
	render.New().Text(w, status, fmt.Sprintf("%d %s", status, msg))
}

// NewMiddleware ... Middlewareを作成する
func NewMiddleware(svc Service) *Middleware {
	return &Middleware{
		Svc: svc,
	}
}
