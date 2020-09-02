package mediator

// Blogger ...
type Blogger interface {
	MakeAd()
	AddMoney(money float64)
	GetMoney() (money float64)
	SetMediator(mediator Mediator)
	Mediator
}

type blogger struct {
	mediator Mediator
	money    float64
	time     int
}

func (b *blogger) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *blogger) Notify(msg string) {
	b.mediator.Notify(msg)
}

func (b *blogger) MakeAd() {
	b.time--
	b.mediator.Notify(doneAd)
}

func (b *blogger) AddMoney(money float64) {
	b.money += money
}

func (b *blogger) GetMoney() (money float64) {
	return b.money
}
