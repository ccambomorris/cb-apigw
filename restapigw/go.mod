module github.com/cloud-barista/cb-apigw/restapigw

go 1.15

// CB-STORE 관련
replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	github.com/xujiajun/nutsdb v0.5.0 => github.com/xujiajun/nutsdb v0.5.1-0.20200320023740-0cc84000d103
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	github.com/cloud-barista/cb-log v0.4.0
	github.com/cloud-barista/cb-store v0.4.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.7.0
	github.com/google/uuid v1.3.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.2 // indirect
	github.com/influxdata/influxdb v1.8.3
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/ugorji/go v1.1.13 // indirect
	github.com/unrolled/secure v1.0.8
	go.etcd.io/bbolt v1.3.5 // indirect
	go.opencensus.io v0.22.5
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/zap v1.18.1 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/api v0.33.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210719143636-1d5a45f8e492 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	sigs.k8s.io/yaml v1.2.0 // indirect
)
