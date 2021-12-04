module github.com/real-mielofon/omp-bot

go 1.17

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/joho/godotenv v1.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozonmp/rtg-service-api/pkg/rtg-service-api v0.0.0-00010101000000-000000000000
	github.com/ozonmp/rtg-service-facade/pkg/rtg-service-facade v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/snovichkov/zap-gelf v1.0.1
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go.uber.org/zap v1.19.1
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozonmp/rtg-service-api/pkg/rtg-service-api => ./pkg/rtg-service-api

replace github.com/ozonmp/rtg-service-facade/pkg/rtg-service-facade => ./pkg/rtg-service-facade

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20211029224645-99673261e6eb // indirect
	golang.org/x/sys v0.0.0-20211031064116-611d5d643895 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
