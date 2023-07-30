package plane

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// harus bikin mock struct dari plane
type MockPassanger struct {
	mock.Mock
}

// implementation
func (m *MockPassanger) GetTotalAdults() int {
	args := m.Called()
	return args.Get(0).(int)
}

func (m *MockPassanger) GetTotalKids() int {
	args := m.Called()
	return args.Get(0).(int)
}

func TestPlanePossibleToTakeOff(t *testing.T) {
	testCases := []struct {
		Name        string
		Expectation bool
	}{
		{
			Name:        "possible take off",
			Expectation: true,
		}, {
			Name:        "impossible take off",
			Expectation: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock := new(MockPassanger)
			newPlane := Plane{
				Passangers: mock,
			}
			if testCase.Expectation {
				mock.On("GetTotalKids").Return(3)
				mock.On("GetTotalAdults").Return(12)
				isPossible := newPlane.IsPossibleToTakeOff()
				assert.Equal(t, testCase.Expectation, isPossible)
			} else {
				mock.On("GetTotalKids").Return(90)
				mock.On("GetTotalAdults").Return(45)
				isNotPossible := newPlane.IsPossibleToTakeOff()
				assert.Equal(t, testCase.Expectation, isNotPossible)
			}
		})
	}
}

func TestInformationPassanger(t *testing.T) {
	mock := new(MockPassanger)
	newPlane := Plane{
		Passangers: mock,
	}

	mock.On("GetTotalAdults").Return(50)
	mock.On("GetTotalKids").Return(10)
	totalAdults, totalKids := newPlane.GetInformationPassanger()
	assert.Equal(t, 50, totalAdults)
	assert.Equal(t, 10, totalKids)
}

func TestTotalPassanger(t *testing.T) {
	passanger := new(Passanger)
	testCases := []struct {
		Name        string
		Expectation int
	}{
		{
			Name:        "adults",
			Expectation: 50,
		}, {
			Name:        "kids",
			Expectation: 10,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Name == "adults" {
				totalAdults := passanger.GetTotalAdults()
				assert.Equal(t, testCase.Expectation, totalAdults)
			} else {
				totalKids := passanger.GetTotalKids()
				assert.Equal(t, testCase.Expectation, totalKids)
			}
		})
	}
}
