package boggart // import "github.com/kihamo/boggart"

import _ "github.com/mailru/easyjson/gen"

//go:generate /bin/bash -c "echo 'Run install tools'"
////go:generate /bin/bash -c "brew install golangci-lint"
////go:generate /bin/bash -c "brew install gettext"
////go:generate /bin/bash -c "go get -u github.com/elazarl/go-bindata-assetfs/..."
////go:generate /bin/bash -c "go get -u golang.org/x/tools/..."
////go:generate /bin/bash -c "go get -u github.com/mailru/easyjson/..."
////go:generate /bin/bash -c "go get -u github.com/alvaroloes/enumer"
//go:generate /bin/bash -c "echo 'Run gettext'"
//go:generate /bin/bash -c "find components/boggart/internal/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/alsa/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/astro/sun/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/broadlink/sp3s/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/chromecast/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/ds18b20/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/esphome/native_api/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/fcm/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/gpio/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/herospeed/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/hikvision/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/homie/esp/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/mercury/v1/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/led_wifi/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/lg_webos/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/nut/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/pulsar/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/xiaomi/roborock/miio/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/boggart/bind/xmeye/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/mqtt/internal/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "find components/storage/internal/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "echo 'Run go bindata'"
//go:generate /bin/bash -c "cd components/boggart/internal && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=internal -nometadata -nomemcopy templates/... assets/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/alsa && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=alsa -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/astro/sun && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=sun -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/boggart && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=boggart -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/broadlink/sp3s && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=sp3s -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/chromecast && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=chromecast -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/dom24 && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=dom24 -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/ds18b20 && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=ds18b20 -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/esphome/mqtt && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=mqtt -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/esphome/native_api && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=nativeapi -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/fcm && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=fcm -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/gpio && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=gpio -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/herospeed && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=herospeed -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/hikvision && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=hikvision -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/homie/esp && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=esp -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/integratorit/elektroset && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=elektroset -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/led_wifi && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=ledwifi -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/lg_webos && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=webos -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/mercury/v1 && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=v1 -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/mercury/v3 && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=v3 -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/nut && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=nut -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/openhab && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=openhab -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/openweathermap && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=openweathermap -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/pass24online && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=pass24online -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/pulsar && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=pulsar -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/timelapse && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=timelapse -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/wol && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=wol -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/xiaomi/roborock/miio && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=miio -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/xiaomi/scale && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=scale -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/xmeye && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=xmeye -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/zigbee/z_stack && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=zstack -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart/bind/zigbee/zigbee2mqtt && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=zigbee2mqtt -nometadata -nomemcopy templates/..."
//go:generate /bin/bash -c "cd components/boggart && enumer -type=BindStatus -trimprefix=BindStatus -output=bind_status_enumer.go"
//go:generate /bin/bash -c "cd components/mqtt/internal && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=internal -nometadata -nomemcopy templates/... locales/..."
//go:generate /bin/bash -c "cd components/storage/internal && go-bindata-assetfs -ignore='(.*?[.]po$)' -o ./bindata_assetfs.go -pkg=internal -nometadata -nomemcopy locales/..."
//go:generate /bin/bash -c "cd components/boggart/bind/alsa && enumer -type=Status -trimprefix=Status -output=status_enumer.go -transform=snake"
//go:generate /bin/bash -c "cd providers/pantum && enumer -type=ProductID -trimprefix=ProductID -output=product_id_enumer.go"
//go:generate /bin/bash -c "cd providers/pantum && enumer -type=CartridgeStatus -trimprefix=CartridgeStatus -output=cartridge_status_enumer.go"
//go:generate /bin/bash -c "cd providers/pantum && enumer -type=DrumStatus -trimprefix=DrumStatus -output=drum_status_enumer.go"
//go:generate /bin/bash -c "cd components/boggart/tasks && enumer -type=Status -trimprefix=Status -output=status_enumer.go"
//go:generate /bin/bash -c "echo 'Run go easyjson'"
//go:generate /bin/bash -c "easyjson components/boggart/internal/handlers/manager.go"
//go:generate /bin/bash -c "easyjson components/boggart/internal/handlers/workers.go"
//go:generate /bin/bash -c "echo 'Run go tools'"
//go:generate /bin/bash -c "goimports -e -w -format-only `find . -type f -name '*.go' -not -path './vendor/*'`"
//go:generate /bin/bash -c "echo 'Run linters'"
//go:generate /bin/bash -c "golangci-lint run -c .golangci.yml"
