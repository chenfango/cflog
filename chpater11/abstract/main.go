package main

import (
	"fmt"
)

// 定义一个结构体

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

// 方法
func(account *Account) Savemoney(money float64, pwd string) {
	// 看下输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 看看存款金额是否正确
	if money <= 0{
		fmt.Println("你输入的金额不正确")
		return 
	}
 
	account.Balance += money
	fmt.Println("存款成功")

}

// 取款
func(account *Account) WithDraw(money float64, pwd string) {
	// 看下输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 看看取款金额是否正确
	if money <= 0 || money> account.Balance { 
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功~")

}

// 查询余额
func(account *Account) Query(pwd string) {
	// 看下输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账户为=%v 余额=%v\n", account.AccountNo, account.Balance)
}

func main(){
	// 测试一把
	account := &Account{
		AccountNo: "gs1111",
		Pwd: "1111",
		Balance: 100.0,
	}

	for {
	// 这里可以做的更加灵活
		account.Query("1111")
		account.Savemoney(10.0, "1111")
		account.Query("1111")
		account.WithDraw(20.0, "1111")
		account.Query("1111")
	}


}