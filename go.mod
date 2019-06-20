module github.com/hortbot/hortbot

go 1.12

require (
	cloud.google.com/go v0.40.0 // indirect
	github.com/Microsoft/go-winio v0.4.12 // indirect
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/aws/aws-sdk-go v1.20.4 // indirect
	github.com/bkaradzic/go-lz4 v1.0.0 // indirect
	github.com/bmatcuk/doublestar v1.1.1
	github.com/cenkalti/backoff v2.2.0+incompatible // indirect
	github.com/cloudflare/golz4 v0.0.0-20150217214814-ef862a3cdc58 // indirect
	github.com/containerd/continuity v0.0.0-20190426062206-aaeac12a7ffc // indirect
	github.com/cznic/b v0.0.0-20181122101859-a26611c4d92d // indirect
	github.com/cznic/fileutil v0.0.0-20181122101858-4d67cfea8c87 // indirect
	github.com/cznic/golex v0.0.0-20181122101858-9c343928389c // indirect
	github.com/cznic/internal v0.0.0-20181122101858-3279554c546e // indirect
	github.com/cznic/sortutil v0.0.0-20181122101858-f5f958428db8 // indirect
	github.com/cznic/strutil v0.0.0-20181122101858-275e90344537 // indirect
	github.com/cznic/zappy v0.0.0-20181122101859-ca47d358d4b1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ericlagergren/decimal v0.0.0-20190507143103-cc8cbe209b64 // indirect
	github.com/fortytw2/leaktest v1.3.0
	github.com/fsouza/fake-gcs-server v1.8.0 // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/gobuffalo/flect v0.1.5
	github.com/gocql/gocql v0.0.0-20190610222256-e00e8c6226e8 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.4.0
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.3.0
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/goware/urlx v0.2.0
	github.com/hako/durafmt v0.0.0-20190612201238-650ed9f29a84
	github.com/jackc/pgx v3.4.0+incompatible // indirect
	github.com/jakebailey/irc v0.0.0-20190407213833-8d2a5d226230
	github.com/jessevdk/go-flags v1.4.1-0.20181221193153-c0795c8afcf4
	github.com/joho/godotenv v1.3.0
	github.com/kshvakov/clickhouse v1.3.7 // indirect
	github.com/leononame/clock v0.1.5
	github.com/lib/pq v1.1.1
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/maxbrunsfeld/counterfeiter/v6 v6.2.0
	github.com/mjibson/esc v0.2.0
	github.com/mongodb/mongo-go-driver v1.0.3 // indirect
	github.com/nakagami/firebirdsql v0.0.0-20190609025626-90ca2b3395f5 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/ory/dockertest v3.3.4+incompatible
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pierrec/lz4 v2.2.3+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/null v8.0.0+incompatible
	github.com/volatiletech/sqlboiler v3.4.0+incompatible
	github.com/xanzy/go-gitlab v0.18.0 // indirect
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036 // indirect
	go.mongodb.org/mongo-driver v1.0.3 // indirect
	go.opencensus.io v0.22.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190618222545-ea8f1a30c443 // indirect
	golang.org/x/net v0.0.0-20190619014844-b5b0513f8c1b // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190619223125-e40ef342dc56 // indirect
	golang.org/x/tools v0.0.0-20190619215442-4adf7a708c2d
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3 // indirect
	google.golang.org/grpc v1.21.1 // indirect
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.3 // indirect
	gotest.tools v2.3.0+incompatible
	mvdan.cc/xurls/v2 v2.0.0
)

replace gopkg.in/DATA-DOG/go-sqlmock.v1 => github.com/DATA-DOG/go-sqlmock v1.3.3

replace github.com/maxbrunsfeld/counterfeiter => github.com/maxbrunsfeld/counterfeiter/v6 v6.2.0
