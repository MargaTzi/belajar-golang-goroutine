package helper

import "sync"

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) addBalance(amount int){
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) getBalance() int{
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}