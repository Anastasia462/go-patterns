package mediator

// Sponsor ...
type Sponsor interface {
	AddMoney(money float64)
	AddFame(fame int)
	OrderAd(money float64)
	GetMoney() (sum float64)
	GetFame() (sum int)
	SetMediator(mediator Mediator)
	Mediator
}

type sponsor struct {
	mediator Mediator
	money    float64
	fame     int
}

func (s *sponsor) Notify(msg string) {
	s.mediator.Notify(msg)
}

func (s *sponsor) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *sponsor) AddMoney(money float64) {
	s.money += money
}

func (s *sponsor) AddFame(fame int) {
	s.fame += fame
}

func (s *sponsor) OrderAd(money float64) {
	s.money -= money
	s.mediator.Notify(makeAd)
}

func (s *sponsor) GetMoney() (sum float64) {
	return s.money
}

func (s *sponsor) GetFame() (sum int) {
	return s.fame
}
