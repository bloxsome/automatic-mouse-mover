package mousemover

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestMover struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestMover))
}

//Run once before each test
func (suite *TestMover) SetupTest() {
	instance = nil
}

func (suite *TestMover) TestSingleton() {
	t := suite.T()
	mouseMover1 := GetInstance()
	mouseMover1.Start()

	time.Sleep(time.Millisecond * 500)

	mouseMover2 := GetInstance()
	assert.True(t, mouseMover2.isRunning(), "instance should not have started")
}
func (suite *TestMover) TestAppStartAndStop() {
	t := suite.T()
	mouseMover := GetInstance()
	mouseMover.Start()
	time.Sleep(time.Millisecond * 500) //wait for app to start
	assert.True(t, mouseMover.isRunning(), "app should have started")

	mouseMover.Quit()
	time.Sleep(time.Millisecond * 500) //wait for app to stop
	assert.False(t, mouseMover.isRunning(), "app should have stopped")
}
