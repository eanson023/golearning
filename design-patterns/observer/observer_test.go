package observer

import "testing"

func TestObserverPattern(t *testing.T) {
	var (
		// 主题
		wd *WeatherData = new(WeatherData)
		// 观察者(小小多态)
		observer Observer = new(GeneralDisplay)
	)
	// 向主题注册观察者
	wd.RegisterObserver(observer)

	data := &Data{
		Temperature: 23,
		Humidity:    65.3,
		Pressure:    45.4,
	}
	wd.SetMeasurements(data)
	data = &Data{
		Temperature: 43,
		Humidity:    25.3,
		Pressure:    49.4,
	}
	wd.SetMeasurements(data)
}
