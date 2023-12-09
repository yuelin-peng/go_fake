package go_fake

import (
	"os"
)

type FakeConfig interface {
	GetType() string
	GetResourceName() string
}

type Fake struct {
	err error
}

type BaseFakeConfig struct {
	Type     string `yaml:"Type"`
	Resource string `yaml:"Resource"`
}

func NewFake() *Fake {
	return &Fake{}
}

func (f *Fake) LoadFakeFile(file string) *Fake {
	c, err := os.ReadFile(file)
	if err != nil {
		util.Logs().Errorf("load fake file failed, err=%v", err)
		f.err = err
		return f
	}
	if len(c) == 0 {
		util.Logs().Warnf("load fake file is empty, file=%v", file)
		return f
	}
	b := BaseFakeConfig{}
	err = yaml.Unmarshal(c, &b)
	if err != nil {
		util.Logs().Errorf("load fake file failed, err=%v", err)
		f.err = err
		return f
	}
	return f
}

func (f *Fake) Error() error {
	return f.err
}
