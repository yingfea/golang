package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Balance int
}

type Banker interface {
	Deposit(amount int) error
	Withdraw(amount int) error
}

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("存入金额必须大于0")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("取款金额必须大于0")
	}

	if a.Balance-amount < 0 {
		return errors.New("账户余额不足")
	}
	a.Balance -= amount
	return nil
}

func main() {
	fmt.Println("--- 练习:银行存取钱 ---")

	account := &Account{Balance: 100}
	fmt.Println("存50")
	var banker Banker = account
	err := banker.Deposit(50)
	if err != nil {
		fmt.Printf("存钱失败: %v\n", err)
	} else {
		fmt.Printf("存钱成功！当前余额: %d\n", account.Balance)
	}

	fmt.Println("--------------------")
	fmt.Println("现在取")
	err = banker.Withdraw(200)
	if err != nil {
		fmt.Printf("取钱失败: %v\n", err)
	} else {
		fmt.Printf("取钱成功！当前余额: %d\n", account.Balance)
	}

}
