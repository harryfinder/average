package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и тд)
type Money int64

// Category представляет собой категорию, в которой был совершён платёж (авто, аптекы,рестораны и тд)
type Category string

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
