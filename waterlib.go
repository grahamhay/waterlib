package waterlib

import (
	"fmt"
	"strings"
)

func CalculateWaterConsumption(FromReading string, UntilReading string) string {
	// TODO Validate the inputs

	myArgs := strings.Split(FromReading, "-")
	fmt.Println(myArgs[0], myArgs[1])
	return "00000,000"
}
