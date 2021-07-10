package builderpattern

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	director := new(Director)
	director.SetBuilder(new(CarBuilder))
	car := director.Generate()
	if res := car.Show(); res != fmt.Sprintf("%s %s %s", wheels, chassis, seat) {
		t.Errorf("error wanted: %s, got: %s", fmt.Sprintf("%s %s %s", wheels, chassis, seat), res)
	}
}
