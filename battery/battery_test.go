package battery_test

import (
	"count/battery"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParsePmsetOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := battery.Status{
		ChargePercent: 47,
	}

	got, err := battery.ParsePmsetOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetPmsetOutput(t *testing.T) {
	t.Parallel()
	text, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}

	status, err := battery.ParsePmsetOutput(text)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Charge: %d%%", status.ChargePercent)
}
