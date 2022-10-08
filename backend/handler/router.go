package handler

import (
    "net/http"

    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func InitRouter() http.Handler {
    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })

    router := mux.NewRouter()

    router.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST")
    router.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET")
    router.Handle("/post/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE")

    router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
    router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")
    
    originsOk := handlers.AllowedOrigins([]string{"*"})
    headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

    return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}