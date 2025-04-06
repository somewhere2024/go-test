package services

import (
	"gin--/internal/dao/mysqldb"
	"gin--/internal/models"
	"gin--/internal/utils/logger"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("this is a test"), nil
	}
}

func CreateToken(payload jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte("this is a test"))
}

func Get_Crruent_User(token string) (*jwt.MapClaims, error) {
	user := &jwt.MapClaims{}
	jwtstr, err := jwt.ParseWithClaims(token, user, Secret())
	if err != nil {
		logger.Logger.Error("token is invalid", zap.Error(err))
		return nil, err
	} else if !jwtstr.Valid {
		logger.Logger.Error("token is invalid", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func VerifyUsernamePassword(username, password string) (*models.User, error) {
	user := &models.User{}
	reulte := mysqldb.DB.Where("username = ?", username).First(user)
	err := reulte.Error
	if user.ID == 0 {
		logger.Logger.Error("用户不存在", zap.Error(err))
		return nil, err
	}
	if user.Password != password {
		logger.Logger.Error("密码无效", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func CreateUser(username, password string) (*models.User, error) {
	user := &models.User{ //定义指针类型，因为无法向 struct{} 类型赋值
		Username: username,
		Password: password,
	}
	result := mysqldb.DB.Create(user)

	if result.Error != nil {
		logger.Logger.Error("创建用户失败", zap.Error(result.Error))
		return nil, result.Error
	}
	return user, nil
}
