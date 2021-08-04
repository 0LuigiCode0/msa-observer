package roots_handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"x-msa-observer/helper"
	"x-msa-observer/store/mongo/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
)

func (h *handler) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenString := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenString, "Bearer ") {
			ctx = context.WithValue(ctx, helper.CtxKeyValue, fmt.Errorf("token invalid"))
			r := r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if token.Method == jwt.SigningMethodHS256 {
				err := token.Claims.Valid()
				if err == nil {
					return []byte(helper.Secret), nil
				}
				return nil, err
			}
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		})

		if err != nil {
			ctx = context.WithValue(ctx, helper.CtxKeyValue, err)
			r := r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
			user, err := h.MongoStore().UserStore().SelectByID(claims.ID)
			if err != nil {
				ctx = context.WithValue(ctx, helper.CtxKeyValue, err)
				r := r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}
			ctx = context.WithValue(ctx, helper.CtxKeyValue, user)
			r := r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func applyCORS(handler http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(handler)
}
