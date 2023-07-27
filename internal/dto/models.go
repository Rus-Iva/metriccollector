package dto

import (
	"encoding/json"
	"github.com/Rus-Iva/metriccollector/internal/storage"
)

type Metrics struct {
	ID    string           `json:"id"`              // имя метрики
	MType string           `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *storage.Counter `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *storage.Gauge   `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

// MarshalJSON реализует интерфейс json.Marshaler.
func (m Metrics) MarshalJSON() ([]byte, error) {
	// чтобы избежать рекурсии при json.Marshal, объявляем новый тип
	type MetricsAlias Metrics

	aliasValue := struct {
		MetricsAlias
		// переопределяем поле внутри анонимной структуры
		Delta *int64   `json:"delta,omitempty"`
		Value *float64 `json:"value,omitempty"`
	}{
		// встраиваем значение всех полей изначального объекта (embedding)
		MetricsAlias: MetricsAlias(m),
		// задаём значение для переопределённого поля
		Delta: (*int64)(m.Delta),
		Value: (*float64)(m.Value),
	}

	return json.Marshal(aliasValue) // вызываем стандартный Marshal
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler.
func (m *Metrics) UnmarshalJSON(data []byte) (err error) {
	// чтобы избежать рекурсии при json.Unmarshal, объявляем новый тип
	type MetricsAlias Metrics

	aliasValue := &struct {
		*MetricsAlias
		// переопределяем поле внутри анонимной структуры
		Delta *int64   `json:"delta,omitempty"`
		Value *float64 `json:"value,omitempty"`
	}{
		MetricsAlias: (*MetricsAlias)(m),
	}
	// вызываем стандартный Unmarshal
	if err = json.Unmarshal(data, aliasValue); err != nil {
		return
	}
	m.Delta = (*storage.Counter)(aliasValue.Delta)
	m.Value = (*storage.Gauge)(aliasValue.Value)
	return
}
