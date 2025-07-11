package envconst

const (
	// UseRootCmd App name and default name command
	UseRootCmd = `ssh+`

	// LongRootCmd Long description root command
	LongRootCmd = `
This application allows you to create alias connections 
and use them when you choose to connect to an ssh connection

To start, create a connection via:
-- ssh+ create

Once created, select your alias connection:
-- ssh+ connect

`
)
