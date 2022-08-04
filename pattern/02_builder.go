package pattern

type Director struct {
	builder Builder
}

type Builder interface {
	AddItem(s string)
	AddDelivery(s string)
	AddPayment(s string)
}

func (d *Director) Construct() {
	d.builder.AddItem("Утюг")
	d.builder.AddDelivery("Улица Пушкина 19")
	d.builder.AddPayment("Оплата картой")
}

type ConcreteBuilder struct {
	product Product
}

func (c *ConcreteBuilder) AddItem(s string) {
	c.product.status += "Заказ продукта: " + s
}

func (c *ConcreteBuilder) AddDelivery(s string) {
	c.product.status += "; Адрес доставки: " + s
}

func (c *ConcreteBuilder) AddPayment(s string) {
	c.product.status += "; Способ оплаты: " + s
}

type Product struct {
	status string
}

func (p *Product) Show() string {
	return p.status
}
