package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

var pmsetOutput = regexp.MustCompile("([0-9]+)%")

type Status struct {
	ChargePercent int
}

func GetPmsetOutput() (string, error) {
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ParsePmsetOutput(output string) (Status, error) {
	matches := pmsetOutput.FindStringSubmatch(output)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", output)
	}

	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{ChargePercent: charge}, nil
}
