package mediator

const (
	makeAd = "Сделай рекламу"
	doneAd = "Сделал рекламу"
	money  = 100
	fame   = 100500
)

// Mediator ...
type Mediator interface {
	Notify(msg string)
}

// ConcreteMediator ...
type ConcreteMediator struct {
	blogger Blogger
	sponsor Sponsor
	manager Manager
}

func (m *ConcreteMediator) Notify(msg string) {
	if msg == makeAd {
		m.manager.TakeFee(money)
		m.blogger.MakeAd()
	} else if msg == doneAd {
		m.blogger.AddMoney(m.manager.GetBLoggerMoney())
		m.sponsor.AddFame(fame)
	}
}

// СonnectСolleagues ...
func СonnectСolleagues(blogger Blogger, manager Manager, sponsor Sponsor) {
	mediator := NewMediator(blogger, sponsor, manager)

	blogger.SetMediator(mediator)
	manager.SetMediator(mediator)
	sponsor.SetMediator(mediator)
}

// NewSponsor ...
func NewSponsor(money float64, fame int) Sponsor {
	return &sponsor{
		money: money,
		fame:  fame,
	}
}

// NewManager ...
func NewManager(money float64, bloggerMoney float64) Manager {
	return &manager{
		money:        money,
		bloggerMoney: bloggerMoney,
	}
}

// NewBlogger ...
func NewBlogger(money float64, time int) Blogger {
	return &blogger{
		money: money,
		time:  time,
	}
}

// NewMediator ...
func NewMediator(
	blogger Blogger,
	sponsor Sponsor,
	manager Manager,
) Mediator {
	return &ConcreteMediator{
		blogger: blogger,
		sponsor: sponsor,
		manager: manager,
	}
}
