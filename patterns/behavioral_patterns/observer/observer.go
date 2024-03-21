package main

import "fmt"

type Observer interface {
	Update(temperature float32)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type WeatherStation struct {
	temperature float32
	observers   []Observer
}

func (w *WeatherStation) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherStation) RemoveObserver(o Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherStation) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature)
	}
}

func (w *WeatherStation) SetTemperature(temp float32) {
	fmt.Println("WeatherStation: New Temperature Measurement:", temp)
	w.temperature = temp
	w.NotifyObservers()
}

type TemperatureDisplay struct {
	subject Subject
}

func (t *TemperatureDisplay) Update(temperature float32) {
	fmt.Println("TemperatureDisplay: New Temperature Received:", temperature)
}

type Fan struct {
	subject Subject
}

func (f *Fan) Update(temperature float32) {
	if temperature > 25 {
		fmt.Println("Fan: It's hot here, turning myself on.")
	} else {
		fmt.Println("Fan: It's nice and cool, turning myself off.")
	}
}

func main() {
	weatherStation := &WeatherStation{}

	tempDisplay := &TemperatureDisplay{}
	fan := &Fan{}

	weatherStation.RegisterObserver(tempDisplay)
	weatherStation.RegisterObserver(fan)

	weatherStation.SetTemperature(30)
	weatherStation.SetTemperature(20)
}
