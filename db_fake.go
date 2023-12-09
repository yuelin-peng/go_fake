package go_fake

type DBFakeConfig struct {
	Resource string                   `yaml:"Resource"`
	Type     string                   `yaml:"Type"`
	LoadData []map[string]interface{} `yaml:"LoadData"`
	Mock     []*MockCase              `yaml:"Mock"`
}

type MockCase struct {
	Command   string `yaml:"Command"`
	Condition string `yaml:"Condition"`
	Resp      DBResp `yaml:"Resp"`
}

type DBResp struct {
	Error        error                    `yaml:"Error"`
	RowsAffected int                      `yaml:"RowsAffected"`
	Data         []map[string]interface{} `yaml:"Data"`
}

type DBFake struct {
	c *DBFakeConfig
}

func NewDBFake(configStr []byte) (*DBFake, error) {
	c := DBFakeConfig{}
	err := yaml.Unmarshal(configStr, &c)
	if err != nil {
		return nil, err
	}
	d := &DBFake{
		c: &c,
	}
	err = d.Fake()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (f *DBFake) Fake() error {

}

func (f *DBFake) LoadData() error {
	if len(f.c.LoadData) <= 0 {
		return nil
	}
	return nil
}
