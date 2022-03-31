package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MakMoinee/appInviteService/cmd/webapp/response"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/common"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/models"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/repository/mysqllocal"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/service"
	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/go-chi/cors"
)

type routesHandler struct {
	MysqlService mysqllocal.TokenIntf
}

type IRoutes interface {
	GenerateToken(w http.ResponseWriter, r *http.Request)
	// LoginToken(token string) error
}

// GenerateToken() - generates token
func (svc *routesHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside GenerateToken()")
	errBuilder := response.ErrorResponse{}
	user := models.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Routes:GenerateToken() -> Error in reading the body")
		errBuilder.ErrorCode = http.StatusInternalServerError
		errBuilder.ErrorMessage = common.TOKEN_ERROR
		response.Error(w, errBuilder)
		return

	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Routes:GenerateToken() -> Unmarshal error")
		errBuilder.ErrorCode = http.StatusInternalServerError
		errBuilder.ErrorMessage = common.TOKEN_ERROR
		response.Error(w, errBuilder)
		return
	}

	token, err := service.SendGenerationOfToken(user, svc.MysqlService)
	if err != nil {
		log.Println("Routes:GenerateToken() -> Failed to generate token. Error Stack: " + err.Error())
		errBuilder.ErrorMessage = err.Error()
		errBuilder.ErrorCode = http.StatusInternalServerError
		response.Error(w, errBuilder)
		return
	}
	response.Success(w, token)
}

func (svc *routesHandler) LoginToken(token string) error {
	var err error
	return err
}

// newRoutes() - returns the IRoutes interface
func newRoutes() IRoutes {
	svc := routesHandler{}
	svc.MysqlService = mysqllocal.NewMysqlService()
	return &svc
}

// Set() - sets the routes for the services
func Set(httpService *goserve.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	localRoutes := newRoutes()
	httpService.Router.Use(cors.Handler)
	initiateRoutes(httpService, localRoutes)
}

// initiateRoutes - attach route handler to the http service and sets the http methods to be used.
func initiateRoutes(httpService *goserve.Service, handler IRoutes) {
	httpService.Router.Post(common.GenerateTokenPath, handler.GenerateToken)
}
