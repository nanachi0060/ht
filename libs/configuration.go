package libs

type Configuration struct {
	ServicePort string
}

func GetConfiguration() (Configuration, error) {
	return Configuration{ServicePort: "8080"}, nil
}
