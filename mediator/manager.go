package mediator

// Manager ...
type Manager interface {
	Talk()
	TakeFee(money float64)
	GetMoney() (sum float64)
	GetBLoggerMoney() (sum float64)
	SetMediator(mediator Mediator)
	Mediator
}

type manager struct {
	mediator     Mediator
	money        float64
	bloggerMoney float64
}

func (m *manager) SetMediator(mediator Mediator) {
	m.mediator = mediator
}

func (m *manager) Notify(msg string) {
	m.mediator.Notify(msg)
}

func (m *manager) Talk() {
}

func (m *manager) TakeFee(money float64) {
	m.money += 0.3 * money
	m.bloggerMoney += money - m.money
}

func (m *manager) GetMoney() (sum float64) {
	return m.money
}

func (m *manager) GetBLoggerMoney() (sum float64) {
	return m.bloggerMoney
}
