package main

import (
	"context"
	"errors"
	"github.com/Tranduy1dol/go-microservice/internal/store"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type postKey string

const postCtx postKey = "post"

type createPostPayload struct {
	Title   string   `json:"title" validate:"required, max=100"`
	Content string   `json:"content" validate:"required, max=1000"`
	Tags    []string `json:"tags"`
}

type updatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty, max=100"`
	Content *string `json:"content" validate:"omitempty, max=1000"`
}

// CreatePost godoc
//
//	@Summary		Create a post
//	@Description	Create a new post
//	@Tags			posts
//	@Accept			JSON
//	@Produce		JSON
//	@Param			payload	body	createPostPayload	true	"Post payload"
//	@Success		201	{object}	store.Post
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router	/posts [post]
func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  1,
	}
	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetPost godoc
//
//	@Summary		Show a post
//	@Description	Get post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Success		201	{object}	store.Post
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/posts/{postID} [get]
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	comments, err := app.store.Comments.GetByPostId(r.Context(), post.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// DeletePost godoc
//
//	@Summary		Delete a post
//	@Description	Delete post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Success		204	{object}	nil
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/posts/{postID} [delete]
func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err = app.store.Posts.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePost godoc
//
//	@Summary		Update a post
//	@Description	Update post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Param			payload	body		updatePostPayload	true	"Update post payload"
//	@Success		200	{object}	store.Post
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/posts/{postID} [patch]
func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	var payload updatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}
	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if err := app.store.Posts.Update(r.Context(), post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()
		post, err := app.store.Posts.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromContext(r *http.Request) *store.Post {
	post := r.Context().Value(postCtx).(*store.Post)
	return post
}
