package weather_test

import (
	"count/weather"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseResponse(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary: "Clouds",
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponseEmpty(t *testing.T) {
	t.Parallel()
	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error, got nil")
	}
}

func TestResponseEmpty(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather_empty.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Fatal("want error, got nil")
	}
}

func TestFormatURL(t *testing.T) {
	t.Parallel()
	c := weather.NewClient("dummy")
	location := "Paris,FR"
	want := "https://api.openweathermap.org/data/2.5/weather?q=Paris,FR&appid=dummy"
	got := c.FormatURL(location)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeather(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/weather.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}))
	defer ts.Close()

	c := weather.NewClient("dummy")
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()

	want := weather.Conditions{
		Summary: "Clouds",
	}
	got, err := c.GetWeather("Paris,FR")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCelsius(t *testing.T) {
	t.Parallel()
	input := weather.Temperature(274.15)
	want := 1.0
	got := input.Celsius()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
