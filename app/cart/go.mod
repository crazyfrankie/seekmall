module github.com/crazyfrankie/seekmall/app/cart

go 1.23.5

replace github.com/crazyfrankie/seekmall/rpc_gen => ../../rpc_gen

require (
	github.com/crazyfrankie/gem v0.0.9
	github.com/crazyfrankie/seekmall/rpc_gen v0.0.0-00010101000000-000000000000
	github.com/google/wire v0.6.0
	github.com/joho/godotenv v1.5.1
	github.com/spf13/viper v1.19.0
	go.etcd.io/etcd/client/v3 v3.5.12
	google.golang.org/grpc v1.70.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.12 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250204164813-702378808489 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250124145028-65684f501c47 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
