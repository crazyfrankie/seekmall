module github.com/crazyfrankie/seekmall/app/payment

go 1.23.5

replace github.com/crazyfrankie/seekmall/rpc_gen => ../../rpc_gen

require (
	github.com/crazyfrankie/seekmall/rpc_gen v0.0.0-00010101000000-000000000000
	github.com/wechatpay-apiv3/wechatpay-go v0.2.20
	go.etcd.io/etcd/client/v3 v3.5.18
	google.golang.org/grpc v1.70.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.etcd.io/etcd/api/v3 v3.5.18 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.18 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250204164813-702378808489 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250124145028-65684f501c47 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
