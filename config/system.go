package config

type Harbor struct {
	Enable   bool   `yaml:"enable" json:"enable"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Scheme   string `json:"scheme" yaml:"scheme"`
}
type System struct {
	Addr        string `json:"addr" yaml:"addr"`
	Provisioner string `json:"provisioner" yaml:"provisioner"`
	Harbor      Harbor `json:"harbor" yaml:"harbor"`
}
