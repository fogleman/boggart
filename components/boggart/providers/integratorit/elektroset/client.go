package elektroset

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/kihamo/boggart/components/boggart/providers/integratorit/internal"
	"go.uber.org/multierr"
)

type Client struct {
	base *internal.Client
}

func New(login, password string) *Client {
	c := &Client{}
	c.base = internal.New("https://lkk.oao-elektroset.ru/gate_lkk_myt", login, password, c.Auth)

	return c
}

func (c *Client) Auth(ctx context.Context) error {
	data := make([]struct {
		KDResult  uint64 `json:"kd_result"`
		NMResult  string `json:"nm_result"`
		IDProfile uint64 `json:"id_profile"`
		AuthCount uint64 `json:"cnt_auth"`
		Session   string `json:"session"`
	}, 0)

	hash := md5.New()
	hash.Write([]byte(c.base.Password()))

	err := c.base.DoRequest(ctx, "auth", "", map[string]string{
		"nm_email":         c.base.Login(),
		"nm_psw":           hex.EncodeToString(hash.Sum(nil)),
		"plugin":           "captchaChecker",
		"captcha_response": "",
	}, &data)

	if err != nil {
		return err
	}

	if data[0].Session == "" {
		return errors.New(data[0].NMResult)
	}

	c.base.SetSession(data[0].Session)

	return err
}

func (c *Client) Accounts(ctx context.Context) ([]Account, error) {
	data := make([]struct {
		Accounts []Account `json:"houses"`
	}, 0)

	err := c.base.DoRequest(ctx, "sql", "lka_get_houses", nil, &data)
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, 0)
	var e error

	for _, row := range data {
		for _, account := range row.Accounts {
			for i, service := range account.Services {
				account.Services[i].Balance, e = strconv.ParseFloat(service.BalanceRAW, 10)
				if e != nil {
					err = multierr.Append(err, e)
				}

				account.Services[i].NNAccount, e = strconv.ParseUint(service.NNAccountRAW, 10, 64)
				if e != nil {
					err = multierr.Append(err, e)
				}
			}

			accounts = append(accounts, account)
		}
	}

	return accounts, err
}