package model

import "github.com/jinzhu/gorm"

type PayState int

const (
	NotPay     PayState = iota //未支付
	BeingPay                   //正在支付
	FaildPay                   //支付失败
	AlreadyPay                 //已支付
)

type AuditState int

const (
	NotAudit     AuditState = iota //未审核
	BeingAudit                     //正在审核
	FaildAudit                     //审核失败
	AlreadyAudit                   //已审核
)

type Orders struct {
	gorm.Model

	UserID   string `json:"user_id" gorm:"comment:'用户ID'"'`
	AdressID string `json:"adress_id" gorm:"comment:'地址ID'"`

	ProductID  string `json:"product_id" gorm:"comment:'商品ID'"`
	ProductNum int    `json:"product_num" gorm:"comment:'商品数量'"`
	TotalPrice int    `json:"total_price" gorm:"comment:'总价'"`

	PayState   PayState `json:"pay_state" gorm:"comment:'支付状态'""`
	AuditState PayState `json:"audit_state" gorm:"comment:'审核状态'""`

	Product []Product `gorm:"many2many:product_orders"`
}
