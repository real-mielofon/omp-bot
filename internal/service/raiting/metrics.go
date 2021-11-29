package raiting

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	notFoundCountTotal = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: "rtg_service_api",
		Name:      "not_found_count_total",
		Help:      "Total count of Not found events",
	})

	cudEventsCountTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "rtg_service_api",
		Name:      "cud_count_total",
		Help:      "Total count of CUD events",
	}, []string{"cudEvent"})

	countEventsInRetranslator = promauto.NewGauge(prometheus.GaugeOpts{
		Subsystem: "rtg_service_api",
		Name:      "events_retranslator_total",
		Help:      "Total count events in retranslator",
	})
)

//go:generate stringer -linecomment -type=cudEvent
type cudEvent uint

const (
	_ = cudEvent(iota)
	// CUDEventCreate create
	CUDEventCreate // create
	// CUDEventUpdate update
	CUDEventUpdate // update
	// CUDEventDelete delete
	CUDEventDelete // delete
)

// AddNotFound add metric NotFound
func AddNotFound() {
	notFoundCountTotal.Add(1)
}

// AddCUDEvent add metric CUD event
func AddCUDEvent(e cudEvent) {
	cudEventsCountTotal.WithLabelValues(e.String()).Add(1)
}

// IncEventsInRetranslator inc metric EventsInRetranslator
func IncEventsInRetranslator(c uint) {
	countEventsInRetranslator.Add(float64(c))
}

// DecEventsInRetranslator dec metric EventsInRetranslator
func DecEventsInRetranslator(c uint) {
	countEventsInRetranslator.Add(-float64(int64(c)))
}

// InitMetrics init metrics
func InitMetrics() {
	countEventsInRetranslator.Set(0)
}
