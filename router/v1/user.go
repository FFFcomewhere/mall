package v1

import (
	"exam8/model"
	"exam8/pkg/app"
	"exam8/pkg/app/request"
	"exam8/pkg/e"
	"exam8/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// @Summary 添加用户
// @Description 添加用户
// @Tags 用户接口
// @ID AddUser
// @Accept application/json
// @Produce application/json
// @Param body body request.AddUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users [post]
func AddUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AddUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	//userInfo := model.UserInfo{
	//	UserName: form.UserName,
	//	Password: form.Password,
	//}

	userModel := model.User{
		Model:    gorm.Model{},
		UserName: form.UserName,
		Password: form.Password,

		//UserInfo: userInfo,
		//
		//Adresses:  nil,
		//Categorys: nil,
		//Cart:      model.Cart{},
		//Orders:    nil,
	}

	fmt.Println("userModel", userModel)

	ok, err := user_service.ExistUserByUserName(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if ok {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_USER_NAME, nil)
	}

	user, err := user_service.AddUser(userModel)
	if err != nil {
		fmt.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, e.SUCCESS, user)
}

// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags 用户接口
// @ID AddUser
// @Accept application/json
// @Produce application/json
// @Param body body request.AddUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users [post]
func UpdateUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.UpdeteUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	userModel := model.User{
		Model:    gorm.Model{},
		UserName: form.UserName,
		Password: form.Password,

		//UserInfo: userInfo,
		//
		//Adresses:  nil,
		//Categorys: nil,
		//Cart:      model.Cart{},
		//Orders:    nil,
	}

	ok, err := user_service.ExistUserByUserName(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_USER_NAME, nil)
		return
	}

	user, err := user_service.UpdateUser(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户接口
// @ID AddUser
// @Accept application/json
// @Produce application/json
// @Param body body request.AddUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users [post]
func GetUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.GetUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	//userInfo := model.UserInfo{
	//	UserName: form.UserName,
	//	Password: form.Password,
	//}

	userModel := model.User{
		Model:    gorm.Model{},
		UserName: form.UserName,

		//UserInfo: userInfo,
		//
		//Adresses:  nil,
		//Categorys: nil,
		//Cart:      model.Cart{},
		//Orders:    nil,
	}

	ok, err := user_service.ExistUserByUserName(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if ok {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_USER_NAME, nil)
	}

	user, err := user_service.GetUser(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// @Summary 删除用户信息
// @Description 删除用户信息
// @Tags 用户接口
// @ID AddUser
// @Accept application/json
// @Produce application/json
// @Param body body request.AddUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/users [post]
func DeleteUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.DeleteUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return
	}

	//userInfo := model.UserInfo{
	//	UserName: form.UserName,
	//	Password: form.Password,
	//}

	userModel := model.User{
		Model:    gorm.Model{},
		UserName: form.UserName,

		//UserInfo: userInfo,
		//
		//Adresses:  nil,
		//Categorys: nil,
		//Cart:      model.Cart{},
		//Orders:    nil,
	}

	ok, err := user_service.ExistUserByUserName(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	fmt.Println("ok:", ok)

	if !ok {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	fmt.Println("")
	if err := user_service.DeleteUser(userModel); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 用户登录
// @Description 用户登录
// @Tags 用户接口
// @ID UserLogin
// @Accept application/json
// @Produce application/json
// @Param body body request.LoginUserForm true "body"
// @Success 200 {object} app.Response "success"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form request.LoginUserForm
	)

	httpCode, errCode, errMsg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg)
		return

	}

	userModel := model.User{
		Model:    gorm.Model{},
		UserName: form.UserName,
		Password: form.Password,

		//UserInfo: userInfo,
		//
		//Adresses:  nil,
		//Categorys: nil,
		//Cart:      model.Cart{},
		//Orders:    nil,
	}

	user, err := user_service.UserLogin(userModel)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}

	if user == nil {
		appG.Response(http.StatusBadRequest, e.ERROR_USER_NAME_OR_PASSWORD, nil)
		return
	}

}
