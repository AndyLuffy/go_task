package task_3

import (
	"gorm.io/gorm"
)

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。*/

type Account struct {
	gorm.Model
	Name    string
	Balance float64
}
type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func Create(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	accounts := []Account{
		{Balance: 1000, Name: "A"},
		{Balance: 1000, Name: "B"},
	}
	db.Create(&accounts)

}

func TransAmount(db *gorm.DB) {

	db.Transaction(func(tx *gorm.DB) error {
		A := Account{}
		B := Account{}

		db.Where("name = ?", "A").Find(&A)
		if A.Balance > 100 {
			db.Model(&A).Update("Balance", A.Balance-100)
			db.Where("name = ?", "B").Find(&B)
			db.Model(&B).Update("Balance", B.Balance+100)
			tran := Transaction{
				FromAccountId: A.ID,
				ToAccountId:   B.ID,
				Amount:        100,
			}
			db.Create(&tran)
		}

		return nil
	})
}
