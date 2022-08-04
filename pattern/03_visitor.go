package pattern

type Visitor interface {
	OrderPhone(i *Phone) string
	OrderHeadphones(i *Headphones) string
	OrderLaptop(i *Laptop) string
}

type Item interface {
	Accept(v Visitor) string
}

type basket struct {
	items []Item
}

type Buyer struct {
}

func (b *Buyer) OrderPhone(i *Phone) string {
	return "phone ordered ->" + b.OrderPhone(i)
}

func (b *Buyer) OrderHeadphones(i *Headphones) string {
	return "Headphones ordered ->" + b.OrderHeadphones(i)
}

func (b *Buyer) OrderLaptop(i *Laptop) string {
	return "Laptop ordered ->" + b.OrderLaptop(i)
}

type Phone struct {
}

func (p *Phone) BuyPhone() string {
	return "phone bought"
}

type Headphones struct {
}

func (p *Phone) BuyHeadphones() string {
	return "Headphones bought"
}

type Laptop struct {
}

func (p *Phone) BuyLaptop() string {
	return "Laptop bought"
}
