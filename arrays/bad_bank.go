package arrays

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(a, b Account, sum float64) Transaction {
	return Transaction{
		From: a.Name,
		To:   b.Name,
		Sum:  sum,
	}
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}

func NewBalanceFor(a Account, t []Transaction) Account {
	return Reduce(t, applyTransaction, a)

}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(acc float64, t Transaction) float64 {
		if t.From == name {
			return acc - t.Sum
		}
		if t.To == name {
			return acc + t.Sum
		}
		return acc
	}

	return Reduce(transactions, adjustBalance, 0.0)
}
