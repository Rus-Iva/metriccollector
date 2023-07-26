package models

import (
	"encoding/json"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

// MarshalJSON реализует интерфейс json.Marshaler.
func (m Metrics) MarshalJSON() ([]byte, error) {
	// чтобы избежать рекурсии при json.Marshal, объявляем новый тип
	type MetricsAlias Metrics
	aliasValue := struct {
		MetricsAlias
		//Value string `json:"value,omitempty"`
	}{
		// встраиваем значение всех полей изначального объекта (embedding)
		MetricsAlias: MetricsAlias(m),
		//Value:        storage.GaugeString(*m.Value),
	}

	return json.Marshal(aliasValue) // вызываем стандартный Marshal
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler.
func (m *Metrics) UnmarshalJSON(data []byte) (err error) {
	// чтобы избежать рекурсии при json.Unmarshal, объявляем новый тип
	type MetricsAlias Metrics

	aliasValue := &struct {
		*MetricsAlias
		//Value string `json:"value,omitempty"`
	}{
		MetricsAlias: (*MetricsAlias)(m),
	}
	// вызываем стандартный Unmarshal
	if err = json.Unmarshal(data, aliasValue); err != nil {
		return
	}
	//val, err := strconv.ParseFloat(aliasValue.Value, 64)
	//if err != nil {
	//	return
	//}
	//m.Value = &val
	return
}
