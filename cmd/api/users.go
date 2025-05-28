package main

import (
	"context"
	"errors"
	"github.com/Tranduy1dol/go-microservice/internal/store"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type userKey string
type followUser struct {
	UserID int64 `json:"user_id"`
}

const userCtx userKey = "user"

// GetUser godoc
//
//	@Summary	Show a user
//	@Description	Get user by ID
//	@Tags	users
//	@Accept	JSON
//	@Produce	JSON
//	@Param			userID	path		int	true	"User ID"
//	@Success		200	{object}	store.User
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/{userID} [get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r.Context())
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

// FollowUser godoc
//
//	@Summary		Follow a user
//	@Description	Follow another user by ID
//	@Tags			users
//	@Accept			JSON
//	@Produce		JSON
//	@Param			userID	path	int	true	"User ID"
//	@Param			payload	body	followUser	true	"User to follow"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/{userID}/follow [put]
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r.Context())

	var payload followUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// UnfollowUser godoc
//
//	@Summary		Unfollow a user
//	@Description	Unfollow another user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int	true	"User ID"
//	@Param			payload	body		followUser	true	"User to unfollow"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/{userID}/unfollow [put]
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowUser := getUserFromContext(r.Context())

	var payload followUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, unfollowUser.ID, payload.UserID); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()
		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(ctx context.Context) *store.User {
	user, _ := ctx.Value(userCtx).(*store.User)
	return user
}
