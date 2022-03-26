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

type myCollector struct {
	FollowersCount *prometheus.Desc
	TweetCount     *prometheus.Desc
}

func newMyCollector() *myCollector {
	return &myCollector{
		FollowersCount: prometheus.NewDesc(
			"GoTwitterExporter_Followers_Count",
			"A number of twitter follower count by id",
			[]string{"TwitterID"},
			nil,
		),
		TweetCount: prometheus.NewDesc(
			"GoTwitterExporter_Tweets_Count",
			"A number of twitter tweet count by id",
			[]string{"TwitterID"},
			nil,
		),
	}
}

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.FollowersCount
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.GetConfig()
	userProfiles := twitter.GetMultipleUserProfiles(conf.TARGET)

	for _, profile := range userProfiles.Users {
		ch <- prometheus.MustNewConstMetric(
			c.FollowersCount,
			prometheus.GaugeValue,
			float64(profile.PublicMetrics.FollowersCount),
			profile.UserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.TweetCount,
			prometheus.GaugeValue,
			float64(profile.PublicMetrics.TweetCount),
			profile.UserName,
		)
	}
}
