package structs

import ()

//Configuration struct
type Configuration struct {
	Title  string
	Server ServerConfiguration
}

//ServerConfiguration struct
type ServerConfiguration struct {
	Port string
}
