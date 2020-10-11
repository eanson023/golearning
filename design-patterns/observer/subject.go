package observer

// Subject 主题 可观察者对象
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// WeatherData 天气数据
type WeatherData struct {
	data *Data
	// 利用组合的方式 观察者集合
	observers []Observer
}

// Data 具体的天气数据
type Data struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
}

// RegisterObserver 注册观察者
func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

// RemoveObserver 删除观察者
func (w *WeatherData) RemoveObserver(o Observer) {
	for idx, v := range w.observers {
		if v == o {
			// 删除元素
			w.observers = append(w.observers[:idx], w.observers[idx+1:]...)
		}
	}
}

// NotifyObservers 发送消息
func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		observer.update(w, w.data)
	}
}

// MeasureChanged 当气象数据改变 气象局会调用此方法
func (w *WeatherData) measureChanged() {
	w.NotifyObservers()
}

// SetMeasurements 利用这个方法测试布告板
func (w *WeatherData) SetMeasurements(data *Data) {
	w.data = data
	w.measureChanged()
}
