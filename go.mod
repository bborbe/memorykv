module github.com/bborbe/memorykv

go 1.24.5

//replace github.com/bborbe/kv => ../kv

exclude (
	cloud.google.com/go v0.26.0
	sigs.k8s.io/structured-merge-diff/v6 v6.0.0
	sigs.k8s.io/structured-merge-diff/v6 v6.1.0
	sigs.k8s.io/structured-merge-diff/v6 v6.2.0
	sigs.k8s.io/structured-merge-diff/v6 v6.3.0
)

require (
	github.com/actgardner/gogen-avro/v9 v9.2.0
	github.com/bborbe/errors v1.3.0
	github.com/bborbe/kv v1.14.2
	github.com/golang/glog v1.2.5
	github.com/google/addlicense v1.1.1
	github.com/incu6us/goimports-reviser/v3 v3.9.1
	github.com/kisielk/errcheck v1.9.0
	github.com/maxbrunsfeld/counterfeiter/v6 v6.11.3
	github.com/onsi/ginkgo/v2 v2.23.4
	github.com/onsi/gomega v1.38.0
	golang.org/x/exp v0.0.0-20250718183923-645b1fa84792
	golang.org/x/lint v0.0.0-20241112194109-818c5a804067
	golang.org/x/vuln v1.1.4
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.9.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20250630185457-6e76a2b096b5 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/incu6us/goimports-reviser v0.1.6 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.65.0 // indirect
	github.com/prometheus/procfs v0.17.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	go.uber.org/automaxprocs v1.6.0 // indirect
	golang.org/x/mod v0.26.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/telemetry v0.0.0-20250721140356-96f361d9aaf7 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/tools v0.35.0 // indirect
	golang.org/x/tools/go/expect v0.1.1-deprecated // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
