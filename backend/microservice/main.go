package main

import (
	"context"
	"direst/repository"
	"direst/resource"
	"direst/service"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/buaazp/fasthttprouter"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
)

func loadEnviromentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	c, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"), option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Fatal("Error getting firestore client")
	}
	return c, nil
}

//CORS ..
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {

	var (
		corsAllowHeaders     = "authorization"
		corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
		corsAllowOrigin      = "*"
		corsAllowCredentials = "true"
	)

	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}

func startHTTPServer(r *fasthttprouter.Router) {
	log.Fatal(fasthttp.ListenAndServe(":"+os.Getenv("APPLICATION_PORT"), CORS(r.Handler)))
}

func main() {
	loadEnviromentVariables()
	ctx := context.Background()
	client, _ := getFirestoreClient(ctx)
	fmt.Println("starting firestore client project id : " + os.Getenv("GOOGLE_PROJECT_ID"))

	openMangaDocumentRef := client.Collection("repository").Doc("openmanga")
	fmt.Println("this project works with next collections : ")
	fmt.Println("/repository/openmanga/users")

	router := fasthttprouter.New()

	openMangacollectionRef := openMangaDocumentRef.Collection("users")
	userRepository := repository.NewUserRepository(ctx, openMangacollectionRef)
	userService := service.NewUserService(userRepository)
	resource.NewUserResource(userService, router)

	resource.NewHealthResource(router)

	fmt.Println("starting application on port : " + os.Getenv("APPLICATION_PORT"))
	startHTTPServer(router)
}
