package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	sponsorMoney = 1000
)

func TestMediator(t *testing.T) {

	blogger := NewBlogger(0, 0)
	manager := NewManager(0, 0)
	sponsor := NewSponsor(sponsorMoney, 0)

	СonnectСolleagues(blogger, manager, sponsor)

	sponsor.OrderAd(money)

	expectBloggerMoney := float64(70)
	expectManagerMoney := float64(30)
	expectSponsorMoney := float64(900)
	expectSponsorFame := fame

	bloggerMoney := blogger.GetMoney()
	managerMoney := manager.GetMoney()
	sponsorMoney := sponsor.GetMoney()
	sponsorFame := sponsor.GetFame()

	assert.Equal(t, expectBloggerMoney, bloggerMoney)
	assert.Equal(t, expectManagerMoney, managerMoney)
	assert.Equal(t, expectSponsorMoney, sponsorMoney)
	assert.Equal(t, expectSponsorFame, sponsorFame)
}
