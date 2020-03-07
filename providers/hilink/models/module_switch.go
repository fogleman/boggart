// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ModuleSwitch module switch
// swagger:model ModuleSwitch
type ModuleSwitch struct {

	// a p n retry enabled
	APNRetryEnabled int64 `json:"APNRetryEnabled,omitempty" xml:"apn_retry_enabled"`

	// b b o u enabled
	BBOUEnabled int64 `json:"BBOUEnabled,omitempty" xml:"bbou_enabled"`

	// c b s enabled
	CBSEnabled int64 `json:"CBSEnabled,omitempty" xml:"cbs_enabled"`

	// charger enabled
	ChargerEnabled int64 `json:"ChargerEnabled,omitempty" xml:"charger_enbaled"`

	// cradle enabled
	CradleEnabled int64 `json:"CradleEnabled,omitempty" xml:"cradle_enabled"`

	// d l n a enabled
	DLNAEnabled int64 `json:"DLNAEnabled,omitempty" xml:"dlna_enabled"`

	// data switch enabled
	DataSwitchEnabled int64 `json:"DataSwitchEnabled,omitempty" xml:"dataswitch_enabled"`

	// eco mode enabled
	EcoModeEnabled int64 `json:"EcoModeEnabled,omitempty" xml:"ecomode_enabled"`

	// encrypt enabled
	EncryptEnabled int64 `json:"EncryptEnabled,omitempty" xml:"encrypt_enabled"`

	// g d p eenabled
	GDPEenabled int64 `json:"GDPEenabled,omitempty" xml:"gdpr_enabled"`

	// help enabled
	HelpEnabled int64 `json:"HelpEnabled,omitempty" xml:"help_enabled"`

	// IPv6 enabled
	IPV6Enabled int64 `json:"IPv6Enabled,omitempty" xml:"ipv6_enabled"`

	// local update enabled
	LocalUpdateEnabled int64 `json:"LocalUpdateEnabled,omitempty" xml:"localupdate_enabled"`

	// monthly volume enabled
	MonthlyVolumeEnabled int64 `json:"MonthlyVolumeEnabled,omitempty" xml:"monthly_volume_enabled"`

	// mult s s ID enable
	MultSSIDEnable int64 `json:"MultSSIDEnable,omitempty" xml:"multssid_enable"`

	// o t a enabled
	OTAEnabled int64 `json:"OTAEnabled,omitempty" xml:"ota_enabled"`

	// p b enabled
	PBEnabled int64 `json:"PBEnabled,omitempty" xml:"pb_enabled"`

	// power off enabled
	PowerOffEnabled int64 `json:"PowerOffEnabled,omitempty" xml:"poweroff_enabled"`

	// powersave enabled
	PowersaveEnabled int64 `json:"PowersaveEnabled,omitempty" xml:"powersave_enabled"`

	// q r code enabled
	QRCodeEnabled int64 `json:"QRCodeEnabled,omitempty" xml:"qrcode_enabled"`

	// s d card enabled
	SDCardEnabled int64 `json:"SDCardEnabled,omitempty" xml:"sdcard_enabled"`

	// s m s enabled
	SMSEnabled int64 `json:"SMSEnabled,omitempty" xml:"sms_enabled"`

	// s n t p enabled
	SNTPEnabled int64 `json:"SNTPEnabled,omitempty" xml:"sntp_enabled"`

	// s t k enabled
	STKEnabled int64 `json:"STKEnabled,omitempty" xml:"stk_enabled"`

	// statistic enabled
	StatisticEnabled int64 `json:"StatisticEnabled,omitempty" xml:"statistic_enabled"`

	// u s s d enabled
	USSDEnabled int64 `json:"USSDEnabled,omitempty" xml:"ussd_enabled"`

	// wi fi enabled
	WiFiEnabled int64 `json:"WiFiEnabled,omitempty" xml:"wifi_enabled"`

	// zone time enabled
	ZoneTimeEnabled int64 `json:"ZoneTimeEnabled,omitempty" xml:"zonetime_enabled"`
}

// Validate validates this module switch
func (m *ModuleSwitch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModuleSwitch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleSwitch) UnmarshalBinary(b []byte) error {
	var res ModuleSwitch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
