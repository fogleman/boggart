package dom24

import (
	"context"
	"fmt"

	"github.com/kihamo/boggart/components/boggart/installer"
	"github.com/kihamo/boggart/components/boggart/installer/openhab"
	"github.com/kihamo/boggart/providers/dom24/client/accounting"
)

func (b *Bind) InstallersSupport() []installer.System {
	return []installer.System{
		installer.SystemOpenHab,
	}
}

func (b *Bind) InstallerSteps(context.Context, installer.System) ([]installer.Step, error) {
	accountingResponse, err := b.provider.Accounting.AccountingInfo(accounting.NewAccountingInfoParams())
	if err != nil {
		return nil, fmt.Errorf("get accounting info failed: %w", err)
	}

	itemPrefix := openhab.ItemPrefixFromBindMeta(b.Meta())
	cfg := b.config()

	const (
		idBalance = "Balance"
		idBill    = "Bill"
	)

	channels := make([]*openhab.Channel, 0, len(accountingResponse.Payload.Data)*2)
	var id string

	for _, account := range accountingResponse.Payload.Data {
		id = "Account" + openhab.IDNormalizeCamelCase(account.Ident) + "_"

		channels = append(channels,
			openhab.NewChannel(id+idBalance, openhab.ChannelTypeNumber).
				WithStateTopic(cfg.TopicAccountBalance.Format(account.Ident)).
				AddItems(
					openhab.NewItem(itemPrefix+id+idBalance, openhab.ItemTypeNumber).
						WithLabel(account.AccountType+" #"+account.Ident+" [%.2f ₽]").
						WithIcon("price"),
				),
			openhab.NewChannel(id+idBill, openhab.ChannelTypeString).
				WithStateTopic(cfg.TopicAccountBill.Format(account.Ident)).
				AddItems(
					openhab.NewItem(itemPrefix+id+idBill, openhab.ItemTypeString).
						WithLabel("Bill #"+account.Ident+" [%s]").
						WithIcon("returnpipe"),
				),
		)
	}

	return openhab.StepsByBind(b, nil, channels...)
}
