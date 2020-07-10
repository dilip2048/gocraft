module github.com/conversenow/balrog

require (
	cloud.google.com/go v0.49.0
	cloud.google.com/go/storage v1.0.0
	github.com/IBM/go-sdk-core v1.1.0
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/Rhymond/go-money v1.0.1
	github.com/ajg/form v1.5.1 // indirect
	github.com/aybabtme/iocontrol v0.0.0-20150809002002-ad15bcfc95a0 // indirect
	github.com/benbjohnson/clock v1.0.0 // indirect
	github.com/buger/jsonparser v0.0.0-20200322175846-f7e751efca13
	github.com/centrifugal/centrifuge-go v0.3.0
	github.com/centrifugal/gocent v2.1.0+incompatible
	github.com/conversenow/elrond v1.1.18
	github.com/conversenow/lembas v0.1.10
	github.com/conversenow/sauron v0.0.10
	github.com/deckarep/golang-set v1.7.1
	github.com/dgraph-io/dgo v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.1
	github.com/friendsofgo/errors v0.9.2
	github.com/gavv/httpexpect v2.0.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-audio/audio v1.0.0
	github.com/go-audio/wav v1.0.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/render v1.0.1
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gocarina/gocsv v0.0.0-20191122093448-c6a9c812ac26
	github.com/golang-migrate/migrate/v4 v4.8.0
	github.com/gordonklaus/portaudio v0.0.0-20180817120803-00e7307ccd93
	github.com/gorilla/websocket v1.4.1
	github.com/huandu/xstrings v1.3.0 // indirect
	github.com/imdario/mergo v0.3.8
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/jesseward/azuretexttospeech v0.0.0-20190317132622-99c2e8afffa4
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/justinas/alice v1.2.0
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/lib/pq v1.0.0
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/nexmo-community/nexmo-go v0.8.1
	github.com/nlopes/slack v0.6.0
	github.com/oauth2-proxy/oauth2-proxy v1.1.2-0.20200514091635-d228d5a92817
	github.com/pkg/errors v0.8.1
	github.com/rs/zerolog v1.17.2
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.3
	github.com/valyala/fasthttp v1.14.0 // indirect
	github.com/volatiletech/authboss v2.3.0+incompatible
	github.com/volatiletech/sqlboiler v3.6.1+incompatible
	github.com/watson-developer-cloud/go-sdk v1.2.0
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	google.golang.org/api v0.20.0
	google.golang.org/genproto v0.0.0-20200331122359-1ee6d9798940
	google.golang.org/grpc v1.29.1
	gopkg.in/auth0.v1 v1.3.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	nhooyr.io/websocket v1.7.4
)

replace (
	//github.com/conversenow/elrond => ../elrond
	//github.com/conversenow/gimli => ../gimli
	//github.com/conversenow/lembas => ../lembas
	github.com/jesseward/azuretexttospeech => github.com/akshaylb/azuretexttospeech v0.0.0-20190904144142-42c3c7cee5d6
	github.com/nlopes/slack => github.com/akshaylb/slack v0.5.1-0.20190227121055-56a7add8c63f
)

go 1.14
