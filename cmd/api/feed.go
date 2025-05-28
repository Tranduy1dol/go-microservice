package main

import (
	"github.com/Tranduy1dol/go-microservice/internal/store"
	"net/http"
)

// GetUserFeed godoc
//
//	@Summary		Get user feed
//	@Description	Get feed for the current user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	store.Post
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user := getUserFromContext(ctx)

	feed, err := app.store.Posts.GetUserFeed(ctx, user.ID, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
