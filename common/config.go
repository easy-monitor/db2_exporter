package common

type Conf struct {
	Modules []*Module `yaml:"module" json:"module"`
}

type Module struct {
	Name     string `yaml:"name" json:"name"`
	Uid string `yaml:"uid" json:"uid"`
	Pid string `yaml:"password" json:"pid"`
}