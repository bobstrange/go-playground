package main

type SensorData struct {
	SensorType string
	ModelID    string
	Value      float32
}

// anti pattern
func ReadValue(r SensorData) float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}

// 構造体のことは構造体が知っているべき
func (r SensorData) ReadValue() float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}
