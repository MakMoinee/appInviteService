package mysqllocal

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MakMoinee/appInviteService/internal/appInviteService/common"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/models"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlService struct {
	DBName           string
	DBUser           string
	DBPassword       string
	ConnectionString string
	Db               *sql.DB
	DbDriver         string
}

type TokenIntf interface {
	// GenerateToken() - generates token from the credentials
	GenerateToken(user models.User) (models.Token, error)
	GetUser(user models.User) (models.User, error)
}

func NewMysqlService() TokenIntf {
	svc := mysqlService{}
	svc.Set()
	return &svc
}

func (svc *mysqlService) Set() {
	svc.DBName = common.DB_NAME
	svc.DBPassword = common.MYSQL_PASSWORD
	svc.DBUser = common.MYSQL_USERNAME
	svc.DbDriver = common.DB_DRIVER
	svc.ConnectionString = svc.DBUser + ":" + svc.DBPassword + "@" + common.CONNECTION_STRING + svc.DBName
	svc.Db = svc.openDBConnection()
	defer svc.Db.Close()
}

func (svc *mysqlService) GetUser(user models.User) (models.User, error) {
	log.Println("Inside Mysqllocal:GetUser()")
	userFromDB := models.User{}
	query := fmt.Sprintf(common.GetUserQuery, user.Username, user.Password)
	svc.Db = svc.openDBConnection()
	result, err := svc.Db.Query(query)
	if err != nil {
		return userFromDB, err
	}
	defer svc.Db.Close()

	for result.Next() {
		err := result.Scan(
			&userFromDB.UserID,
			&userFromDB.Username,
			&userFromDB.Password,
			&userFromDB.UserType,
			&userFromDB.CreatedAt,
			&userFromDB.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return userFromDB, err
		}
	}
	defer result.Close()
	return user, err
}

func (svc *mysqlService) GenerateToken(user models.User) (models.Token, error) {
	list := models.Token{}
	var err error

	return list, err
}

func (svc *mysqlService) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}
