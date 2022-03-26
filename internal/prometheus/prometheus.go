package prometheus

import (
	"go-twitter-exporter/internal/config"
	"go-twitter-exporter/internal/twitter"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitExporter() {
	c := newMyCollector()
	r := prometheus.NewRegistry()
	r.MustRegister(c)
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)
	log.Fatal(http.ListenAndServe(":2020", nil))
}

// Metricsの定義

type myCollector struct {
	FollowersCount *prometheus.Desc
} // 今回働いてくれるインスタンス

func newMyCollector() *myCollector {
	return &myCollector{
		FollowersCount: prometheus.NewDesc(
			"GoTwitterExporter_Followers_Count",
			"A number of twitter follower count by id",
			[]string{"TwitterID"},
			nil,
		),
	}
}

// Describe と Collect

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.FollowersCount
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.GetConfig()
	for _, targetId := range conf.TARGET {
		followersCount := twitter.GetFollowersCount(targetId)

		ch <- prometheus.MustNewConstMetric(
			c.FollowersCount,
			prometheus.GaugeValue,
			float64(followersCount),
			targetId,
		)
	}
}
