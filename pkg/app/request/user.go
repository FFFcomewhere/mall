package request

type LoginUserForm struct {
	UserName string `json:"user_name" form:"user_name" vaild:"Required"` //用户名
	Password string `json:"password" form:"password" vaild:"MinSize(8)"` //密码
}

type AddUserForm struct {
	UserName        string `json:"user_name" form:"user_name" vaild:"Required"`         //用户名
	Password        string `json:"password" form:"password" vaild:"MinSize(8)"`         //密码
	ConfirmPassword string `json:"confirm_password" form:"password" vaild:"MinSize(8)"` //密码
}

type UpdeteUserForm struct {
	UserName string `json:"user_name" form:"user_name" vaild:"Required"` //用户名
	Password string `json:"password" form:"password" vaild:"MinSize(8)"` //密码
}

type GetUserForm struct {
	UserName string `json:"user_name" form:"user_name" vaild:"Required"` //用户名
}

type DeleteUserForm struct {
	UserName string `json:"user_name" form:"user_name" vaild:"Required"` //用户名
}
