package config

import (
	"github.com/fspcons/ports-service/src/utils"
	"os"
)

// Data is a config wrapper struct
type Data struct {
	RestAPIAddress string
	PortsFilePath  string
	//... additional config for the service could be added here
}

// IsValid checks if config is valid
func (ref *Data) IsValid() (bool, string) {
	if utils.IsEmpty(ref.PortsFilePath) {
		return false, "the env var 'PORTS_FILE_PATH' is invalid"
	}
	if utils.IsEmpty(ref.RestAPIAddress) {
		ref.RestAPIAddress = ":8080" //default value
	}

	return true, ""
}

// ReadFromEnv returns structs holding app config data that was set in
// environment variables. It panics on unset environment variables.
func ReadFromEnv() Data {
	data := Data{
		RestAPIAddress: os.Getenv("REST_ADDRESS"),
		PortsFilePath:  os.Getenv("PORTS_FILE_PATH"),
	}
	if ok, msg := data.IsValid(); !ok {
		panic(msg)
	}

	return data
}
