package pattern
import (
	"fmt"
)

type client struct {
	isShoppingCart bool
	isNightstand   bool
	isMeatballs    bool
	isPaymentDone  bool
}

type department interface {
	execute(*client)
	setNext(department)
}

type shoppingCart struct {
	next department
}

func (s *shoppingCart) execute(c *client) {
	if c.isShoppingCart {
		fmt.Println("Shop cart taken")
		s.next.execute(c)
		return
	}
	fmt.Println("Shop cart already taken")
	c.isShoppingCart = true
	s.next.execute(c)
}

func (s *shoppingCart) setNext(next department) {
	s.next = next
}

type furniture struct {
	next department
}

func (f *furniture) execute(c *client) {
	if c.isShoppingCart {
		fmt.Println("nightstand bought")
		f.next.execute(c)
		return
	}
	fmt.Println("nightstand already bought")
	c.isNightstand = true
	c.isPaymentDone = false
	f.next.execute(c)
}

func (f *furniture) setNext(next department) {
	f.next = next
}

type cafe struct {
	next department
}

func (ca *cafe) execute(c *client) {
	if c.isMeatballs {
		fmt.Println("meatballs bought")
		ca.next.execute(c)
		return
	}
	fmt.Println("meatballs already bought")
	c.isMeatballs = true
	c.isPaymentDone = false
	ca.next.execute(c)
}

func (ca *cafe) setNext(next department) {
	ca.next = next
}

type cashbox struct {
	next department
}

func (ca *cashbox) execute(c *client) {
	if c.isPaymentDone {
		fmt.Println("Payment already done")
		ca.next.execute(c)
		return
	}
	fmt.Println("payment done")
	c.isPaymentDone = true
	ca.next.execute(c)
}

func (ca *cashbox) setNext(next department) {
	ca.next = next
}
