package pattern

import "fmt"

type ItemStat struct {
	hasItem      state
	noItem       state
	itemInBasket state
	paid         state

	currentState state

	itemCount int
	itemPrice int
}

func newItemState(itemCount, itemPrice int) *ItemStat {
	i := &ItemStat{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		ItemStat: i,
	}

	noItemState := &noItemState{
		ItemStat: i,
	}

	itemInBasketState := &itemInBasketState{
		ItemStat: i,
	}

	paidState := &paidState{
		ItemStat: i,
	}

	i.setState(hasItemState)
	i.hasItem = hasItemState
	i.noItem = noItemState
	i.itemInBasket = itemInBasketState
	i.paid = paidState
	return i
}

func (i *ItemStat) addItem(count int) error {
	return i.currentState.addItem(count)
}

func (i *ItemStat) addToBasket() error {
	return i.currentState.addToBasket()
}

func (i *ItemStat) payItem(money int) error {
	return i.currentState.payItem(money)
}

func (i *ItemStat) deliverItem() error {
	return i.currentState.deliverItem()
}

func (i *ItemStat) setState(s state) {
	i.currentState = s
}

func (i *ItemStat) incItemCount(count int) {
	i.itemCount = i.itemCount + count
}

type state interface {
	addItem(int) error
	addToBasket() error
	payItem(money int) error
	deliverItem() error
}

type hasItemState struct {
	ItemStat *ItemStat
}

func (i *hasItemState) addItem(count int) error {
	fmt.Println("Добавлено", count, "товаров")
	i.ItemStat.incItemCount(count)
	return nil
}

func (i *hasItemState) addToBasket() error {
	if i.ItemStat.itemCount == 0 {
		i.ItemStat.setState(i.ItemStat.noItem)
		return fmt.Errorf("Товар закончился")
	}
	fmt.Println("Товар добавлен в корзину")
	i.ItemStat.setState(i.ItemStat.itemInBasket)
	return nil
}

func (i *hasItemState) payItem(money int) error {
	return fmt.Errorf("Сначала выберете товар")
}

func (i *hasItemState) deliverItem() error {
	return fmt.Errorf("Сначала добавьте товар")
}

type noItemState struct {
	ItemStat *ItemStat
}

func (i *noItemState) addItem(count int) error {
	fmt.Println("Добавлено", count, "товаров")
	i.ItemStat.setState(i.ItemStat.hasItem)
	i.ItemStat.incItemCount(count)
	return nil
}

func (i *noItemState) addToBasket() error {
	return fmt.Errorf("Товар закончился")
}

func (i *noItemState) payItem(money int) error {
	return fmt.Errorf("Товар закончился")
}

func (i *noItemState) deliverItem() error {
	return fmt.Errorf("Товар закончился")
}

type itemInBasketState struct {
	ItemStat *ItemStat
}

func (i *itemInBasketState) addItem(count int) error {
	return fmt.Errorf("Товар уже добавлен в корзину")
}

func (i *itemInBasketState) addToBasket() error {
	return fmt.Errorf("Товар уже добавлен в корзину")
}

func (i *itemInBasketState) payItem(money int) error {
	if money < i.ItemStat.itemPrice {
		return fmt.Errorf("Не хватает средств, внесите еще", (i.ItemStat.itemPrice - money))
	}
	fmt.Println("Товар оплачен")
	i.ItemStat.setState(i.ItemStat.paid)
	return nil
}

func (i *itemInBasketState) deliverItem() error {
	return fmt.Errorf("Сначала оплатите товар")
}

type paidState struct {
	ItemStat *ItemStat
}

func (i *paidState) addItem(count int) error {
	return fmt.Errorf("Товар уже оплачен")
}

func (i *paidState) addToBasket() error {
	return fmt.Errorf("Товар уже оплачен")
}

func (i *paidState) payItem(money int) error {
	return fmt.Errorf("Товар уже оплачен")
}

func (i *paidState) deliverItem() error {
	i.ItemStat.itemCount = i.ItemStat.itemCount - 1
	if i.ItemStat.itemCount == 0 {
		i.ItemStat.setState(i.ItemStat.noItem)
	} else {
		i.ItemStat.setState(i.ItemStat.hasItem)
	}
	return nil
}
