package routers

import (	
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/shijuvar/go-recipes/ch07/bookmarkapi/controllers"
	"github.com/shijuvar/go-recipes/ch07/bookmarkapi/common"
)

// SetBookmarkRoutes registers routes for bookmark entity.
func SetBookmarkRoutes(router *mux.Router) *mux.Router {
	bookmarkRouter := mux.NewRouter()
	bookmarkRouter.HandleFunc("/bookmarks", controllers.CreateBookmark).Methods("POST")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controllers.UpdateBookmark).Methods("PUT")
	bookmarkRouter.HandleFunc("/bookmarks", controllers.GetBookmarks).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controllers.GetBookmarkByID).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/users/{id}", controllers.GetBookmarksByUser).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controllers.DeleteBookmark).Methods("DELETE")
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(bookmarkRouter),
	))
	return router
}
