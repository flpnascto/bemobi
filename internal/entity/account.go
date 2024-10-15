package entity

import (
	"errors"
	"sync"
)

type application []float32

type Account struct {
	applicationHistory []application
}

func NewAccount(applicationNumber int) (*Account, error) {
	if applicationNumber <= 0 {
		return nil, errors.New("applicationNumber must be greater than 0")
	}

	app := make(application, applicationNumber)
	for i := range app {
		app[i] = 0
	}

	account := &Account{
		applicationHistory: []application{app},
	}
	return account, nil

}

func (a *Account) MakeApplication(application Investment) error {
	if len(a.applicationHistory) == 0 {
		return errors.New("no applications in history")
	}

	lastApplication := a.applicationHistory[len(a.applicationHistory)-1]

	newApplication := append([]float32(nil), lastApplication...)

	for i := application.left - 1; i < application.right; i++ {
		newApplication[i] += application.amount
	}
	a.applicationHistory = append(a.applicationHistory, newApplication)
	return nil
}

func (a *Account) MaxInvestmentValue() float32 {
	if len(a.applicationHistory) == 0 {
		return 0
	}

	c1 := make(chan float32)
	var wg sync.WaitGroup
	for i := 0; i < len(a.applicationHistory); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			readApplication(a.applicationHistory[i], c1)
		}(i)
	}

	go func() {
		wg.Wait()
		close(c1)
	}()

	var maxValue = float32(0.0)
	for c := range c1 {
		if c > maxValue {
			maxValue = c
		}
	}

	return maxValue
}

func readApplication(data []float32, c chan float32) {
	var maxValue = float32(0.0)
	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}
	channelReceiver(maxValue, c)
}

func channelReceiver(data float32, c chan<- float32) {
	c <- data
}
