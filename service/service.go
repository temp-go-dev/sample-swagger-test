package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikkeloscar/gin-swagger/api"
	"github.com/temp-go-dev/sample-swagger/models"
	"github.com/temp-go-dev/sample-swagger/restapi/operations/user_api"
)

// インターフェース　restapi/api.go の中に定義される

// Svc インターフェースを実装するための構造体
type Svc struct{}

// GetUserByUserID 業務ロジックの実装
func (m *Svc) GetUserByUserID(ctx *gin.Context, params *user_api.GetUserByUserIDParams) *api.Response {
	fmt.Println("GetUserByUserID!!!!!")

	// なにかしらの業務ロジックを書く
	userid := params.UserID
	fmt.Println(userid)

	return &api.Response{
		Code: http.StatusOK,
		Body: &models.User{
			ID:   9999,
			Name: "OK!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
		},
	}
}

// Healthy ヘルスチェックのメソッド？
func (m *Svc) Healthy() bool {
	fmt.Println("health")
	return true
}
