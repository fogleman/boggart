package mikrotik

import (
	"context"
	"net"

	"github.com/kihamo/boggart/components/mqtt"
)

// подписываемся на собственные топики, это нужно для того чтобы при первом старте
// выключить активность соединений, которые удалены в самом микротике и соответственно
// они будут не доступны в последующий обновлениях, а в топике и них ретеншен сообщение с true
func (b *Bind) callbackMQTTInterfacesZombies(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
	if message.IsFalse() || !message.Retained() {
		return nil
	}

	topic := message.Topic()
	parts := topic.Split()

	if len(parts) < 3 {
		return nil
	}

	// проверяем наличие в списке
	item := &storeItem{
		version:        0,
		connectionName: parts[len(parts)-1],
		interfaceName:  parts[len(parts)-2],
		interfaceType:  parts[len(parts)-3],
	}

	if _, ok := b.connectionsActive.Load(item.String()); !ok {
		return b.MQTT().PublishAsyncRaw(ctx, topic, message.Qos(), true, false)
	}

	return nil
}

// Для ускорения проверки изменения в подключениях, чекаем syslog сообщения и парсим оттуда необходиму инфу
func (b *Bind) callbackMQTTSyslog(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
	var request map[string]interface{}

	if err := message.JSONUnmarshal(&request); err != nil {
		return err
	}

	tag, ok := request["tag"]
	if !ok {
		return nil
	}

	content, ok := request["content"]
	if !ok {
		return nil
	}

	switch tag {
	case b.config.SyslogTagWireless:
		check := wifiClientRegexp.FindStringSubmatch(content.(string))
		if len(check) < 4 {
			return nil
		}

		if _, err := net.ParseMAC(check[1]); err != nil {
			return err
		}

		if check[3] == "connected" || check[3] == "disconnected" {
			return b.Workers().TaskRunByName(ctx, TaskNameInterfaceConnection)
		}

		return b.Workers().TaskRunByName(ctx, TaskNameInterfaceConnection)

	case b.config.SyslogTagL2TP:
		check := vpnClientRegexp.FindStringSubmatch(content.(string))
		if len(check) < 2 {
			return nil
		}

		if check[2] == "in" || check[2] == "out" {
			return b.Workers().TaskRunByName(ctx, TaskNameInterfaceConnection)
		}
	}

	return nil
}
