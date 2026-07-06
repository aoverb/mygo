package bankaccount
import (
    "sync"
)

type Account struct {
    Amount int64
    Mutex sync.Mutex
    Opened bool
}

func Open(amt int64) *Account {
    if amt < 0 {
        return nil
    }
    acc := &Account{
        Amount: amt,
        Opened: true,
    }
    return acc
}

func (a *Account) Balance() (bal int64, ok bool) {
    if a == nil {
        return 0, false
    }
    a.Mutex.Lock()
    defer a.Mutex.Unlock()
    if !a.Opened {
        return 0, false
    }
	return a.Amount, true
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
    if a == nil {
        return 0, false
    }
    a.Mutex.Lock()
    defer a.Mutex.Unlock()
    if !a.Opened || a.Amount + amt < 0 {
        return 0, false
    }
	a.Amount += amt
    return a.Amount, true
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
    if a == nil || amt < 0 {
        return 0, false
    }
    a.Mutex.Lock()
    defer a.Mutex.Unlock()
    if !a.Opened || a.Amount < amt {
        return 0, false
    }
	a.Amount -= amt
    return a.Amount, true
}

func (a *Account) Close() (pay int64, ok bool) {
    if a == nil {
        return 0, false
    }
    a.Mutex.Lock()
    defer a.Mutex.Unlock()
    if !a.Opened {
        return 0, false
    }
	a.Opened = false
    pay = a.Amount
    a.Amount = 0
    return pay, true
}
