package controller

import (
	"frp-panel/middleware"
	"frp-panel/model"
	"frp-panel/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// Login @Summary 登录
// @Schemes
// @Description user login
// @Tags user
// @Accept json
// @Produce json
// @Param message body model.LoginParams true "login info"
// @Success 200
// @Router /api/login/account [post]
func Login(ctx *gin.Context) {
	var loginParams = model.LoginParams{}
	_ = ctx.ShouldBindJSON(&loginParams)

	username := loginParams.Username
	password := loginParams.Password

	if username != "admin" {
		response.LoginResponse(ctx, http.StatusOK, model.LoginResult{
			Status:           "error",
			Type:             "account",
			CurrentAuthority: "guest",
			Token:            "",
		})
		return
	}

	if password != "ant.design" {
		response.LoginResponse(ctx, http.StatusOK, model.LoginResult{
			Status:           "error",
			Type:             "account",
			CurrentAuthority: "guest",
			Token:            "",
		})
		return
	}

	j := middleware.NewJWT()
	token, _ := j.CreateToken(middleware.CustomClaims{
		ID:               "1",
		Name:             "admin",
		RegisteredClaims: jwt.RegisteredClaims{},
	})

	response.LoginResponse(ctx, http.StatusOK, model.LoginResult{
		Status:           "ok",
		Type:             "account",
		CurrentAuthority: "admin",
		Token:            token,
	})

}

func OutLogin(ctx *gin.Context) {
	response.Success(ctx, true)
}

func CurrentUser(ctx *gin.Context) {
	currentUser := model.CurrentUser{
		Name:   "frp-panel",
		Avatar: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Userid: "00000001",
		//Email:       "antdesign@alipay.com",
		//Signature:   "海纳百川，有容乃大",
		//Group:       "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
		Tags:        nil,
		NotifyCount: 12,
		UnreadCount: 11,
		Country:     "China",
		Access:      "admin",
		Geographic: model.Geographic{
			Province: model.Province{
				Label: "浙江省",
				Key:   "330000",
			},
			City: model.City{
				Label: "杭州市",
				Key:   "33010",
			},
		},
		Address: "西湖区工专路 77 号",
		Phone:   "0752-268888888",
	}
	response.Success(ctx, currentUser)
}
