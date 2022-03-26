# go-twitter-exporter

## Description

A prometheus exporter for twitter profiles.  
Current exported data.

- Followers count
- Tweet count

```
# HELP GoTwitterExporter_Followers_Count A number of twitter follower count by id
# TYPE GoTwitterExporter_Followers_Count gauge
GoTwitterExporter_Followers_Count{TwitterID="Twitter"} 6.1121053e+07
GoTwitterExporter_Followers_Count{TwitterID="TwitterDev"} 536663
# HELP GoTwitterExporter_Tweets_Count A number of twitter tweet count by id
# TYPE GoTwitterExporter_Tweets_Count gauge
GoTwitterExporter_Tweets_Count{TwitterID="Twitter"} 14995
GoTwitterExporter_Tweets_Count{TwitterID="TwitterDev"} 3880
```

## Configuration

`config.yaml`

```
TWITTER:
  TWITTER_BEARER_TOKEN: "TWITTER_BEARER_TOKEN"
TARGET:
  - "Twitter"
  - "TwitterDev"
```


