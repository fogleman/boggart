package xmeye

type Response struct {
	Name      string
	Ret       uint64
	SessionID string
}

type LoginResponse struct {
	AliveInterval uint64
	ChannelNum    uint64
	DataUseAES    bool
	DeviceType    string `json:DeviceType `
	ExtraChannel  uint64
	Ret           uint64
	SessionID     Uint32
}

type LogSearch struct {
	Data     string
	Position uint32
	Time     Time
	Type     string
	User     string
}

type FileSearch struct {
	BeginTime  Time
	DiskNo     uint32
	EndTime    Time
	FileLength string
	FileName   string
	SerialNo   uint32
}

type SystemInfo struct {
	AlarmInChannel  uint64
	AlarmOutChannel uint64
	AudioInChannel  uint64
	BuildTime       Time
	CombineSwitch   uint64
	DeviceRunTime   string
	DigChannel      uint64
	EncryptVersion  string
	ExtraChannel    uint64
	HardWare        string
	HardWareVersion string
	SerialNo        string
	SoftWareVersion string
	TalkInChannel   uint64
	TalkOutChannel  uint64
	UpdataTime      Time
	UpdataType      string
	VideoInChannel  uint64
	VideoOutChannel uint64
}

type OEMInfo struct {
	Address   string
	Name      string
	OEMID     uint64
	Telephone string
}

type StorageInfo struct {
	PartNumber uint64
	PlysicalNo uint64
	Partition  []struct {
		DirverType    uint64
		IsCurrent     bool
		LogicSerialNo uint64
		NewEndTime    Time
		NewStartTime  Time
		OldEndTime    Time
		OldStartTime  Time
		Status        uint64
		RemainSpace   Uint32
		TotalSpace    Uint32
	}
}

type WorkState struct {
	AlarmState struct {
		AlarmIn     uint64
		AlarmOut    uint64
		VideoBlind  uint64
		VideoLoss   uint64
		VideoMotion uint64
	}
	ChannelState []struct {
		Bitrate uint64
		Record  bool
	}
}

type AlarmInfo struct {
	Channel   uint8
	Event     string
	StartTime Time
	Status    string
}
