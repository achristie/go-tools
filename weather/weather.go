package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://api.openweathermap.org"

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}

type Client struct {
	key        string
	BaseURL    string
	HTTPClient *http.Client
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

type Conditions struct {
	Summary     string
	Temperature Temperature
}

func NewClient(apikey string) *Client {
	return &Client{
		BaseURL: baseURL,
		key:     apikey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var owm OWMResponse
	err := json.Unmarshal(data, &owm)
	if err != nil {
		return Conditions{}, err
	}

	if len(owm.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid response %s: no weather available", data)
	}
	return Conditions{
		Summary:     owm.Weather[0].Main,
		Temperature: Temperature(owm.Main.Temp),
	}, nil
}

func (c *Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.key)
}

func MakeAPIRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected status", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}

	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}

	return conditions, nil
}

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}

	return conditions, nil
}

func RunCLI() {

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s LOCATION\n\nExample: %[1]s London,UK", os.Args[0])
	}
	location := os.Args[1]
	key := ""

	conditions, err := Get(location, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %.1f\n", conditions.Summary, conditions.Temperature.Celsius())
}
