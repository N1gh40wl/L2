package main

import (
	"fmt"
	"log"
)

type CardFacade struct {
	cardNumber   *CardNumber
	securityCode *SecurityCode
	balance      *Balance
}

func newCardFacade(cardNumber string, code int) *CardFacade {
	CardFacade := &CardFacade{
		cardNumber:   newCardNumber(cardNumber),
		securityCode: newSecurityCode(code),
		balance:      newBalance(100),
	}
	return CardFacade
}

func (c *CardFacade) checkBalance(cardNumber string, code int) error {
	err := c.cardNumber.checkNumber(cardNumber)
	if err != nil {
		return err
	}
	err = c.securityCode.checkCode(code)
	if err != nil {
		return err
	}
	c.balance.checkBalance()
	return nil
}

type CardNumber struct {
	number string
}

func newCardNumber(number string) *CardNumber {
	return &CardNumber{
		number: number,
	}
}

func (c *CardNumber) checkNumber(number string) error {
	if c.number != number {
		return fmt.Errorf("Card number is incorrect")
	}
	fmt.Println("Card number verified")
	return nil
}

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("SecurityCode is incorrect")
	}
	fmt.Println("SecurityCode verified")
	return nil
}

type Balance struct {
	amount int
}

func newBalance(amount int) *Balance {
	return &Balance{
		amount: amount,
	}
}

func (s *Balance) checkBalance() error {

	fmt.Println("Your balance is", s.amount)
	return nil
}

func main() {
	card := newCardFacade("1111 1111 1111 1111", 123)
	err := card.checkBalance("1111 1111 1111 1111", 123)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
