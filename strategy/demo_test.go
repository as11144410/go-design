package strategy

import "testing"

func TestStrategy_Do(t *testing.T) {
	c := Context{}
	strategy1 := &Strategy1{}
	c.Strategy = strategy1
	c.Strategy.Do()

	strategy2 := &Strategy2{}
	c.Strategy = strategy2
	c.Strategy.Do()
}
