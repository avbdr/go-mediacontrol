package device

type Device interface {
	// Methods
	Init()
	PowerOff() error
	PowerOn() error
	List() error
	Open(url string) error
	Key(key string) error
	VolUp() error
	VolDown() error
	Vol(vol int) (int, error)
	Test() error
	Text(text string) error
	Stream(url string) error
	Info() (string, error)
	Next() error
	Prev() error
	Pause() error
	Play() error
	GetToken() string
	Status() (interface{}, error)
}

type DeviceInfo struct {
	Name  string `json:"name"`
	Mac   string `json:"mac"`
	Ip    string `json:"ip"`
	Type  string `json:"type"`
	Token string `json:"token,omitempty"`
	Udn  string  `json:"udn"`
}

func Find (devices []DeviceInfo, dev DeviceInfo) int {
	for i, d := range devices {
		if d.Udn == dev.Udn {
			return i
		}
	}
	return -1
}
