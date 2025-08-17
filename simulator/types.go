package simulator

type Config struct {
	SendMqtt bool `default:"true" yaml:"send-mqtt"`
	Mqtt     struct {
		Host     string `yaml:"mqtt-host"`
		Port     int    `yaml:"mqtt-port"`
		User     string `yaml:"mqtt-user"`
		Password string `yaml:"mqtt-password"`
	} `yaml:"mqtt"`
	API struct {
		Endpoint         string `yaml:"endpoint"`
		ApiKeyHeaderName string `yaml:"api-key-header-name"`
		Key              string `yaml:"api-key"`
	} `yaml:"api"`
}

type Message struct {
	Title      string `yaml:"title"`
	Device     string `yaml:"device"`
	Type       string `yaml:"type"`
	Rate       string `yaml:"rate"`
	RateType   MessageRate
	RateNumber int
	Max        float64   `yaml:"max"`
	Min        float64   `yaml:"min"`
	Options    []string  `yaml:"options"`
	Body       []Message `yaml:"body"`
}

type Specs struct {
	// Devices  []string  `yaml:"devices"`
	Messages []Message `yaml:"messages"`
}

func (m *Message) SetDefaults() {
	if m.Device == "" {
		m.Device = "default_device"
	}
	if m.Type == "" {
		m.Type = "int"
	}
	if m.Rate == "" {
		m.Rate = "once"
	}
	if m.Max <= m.Min {
		m.Max = 100
		m.Min = 0
	}
	for i := range m.Body {
		m.Body[i].SetDefaults()
	}
}
