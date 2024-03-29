// Package config provides a simple way to configure and auto-document an app with environment variables.
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parameter describes a simple configuration parameter
type Parameter struct {
	Name         string
	Mandatory    bool
	Description  string
	DefaultValue string
}

// Config describes multiple configuration parameters
type Config struct {
	Parameters map[string]Parameter
	Prefix     string
}

// SetPrefix sets a prefix for all environment variables.
//
// Example that for the environment variable 'EXAMPLE_PORT':
//     config.SetPrefix("EXAMPLE_")
//     config.DescribeInt("PORT", "The port to listen to", false, 1234)
//     port := config.GetInt("PORT")
//
// It makes sense to set the prefix to the application name to prevent
// collisions with environment variables for other applications.
func (c *Config) SetPrefix(prefix string) {
	c.Prefix = prefix
}

func (c *Config) describe(name string, description string, mandatory bool, defaultValue string) {
	if c.Parameters == nil {
		c.Parameters = make(map[string]Parameter)
	}

	c.Parameters[name] = Parameter{
		Name:         name,
		Description:  description,
		Mandatory:    mandatory,
		DefaultValue: defaultValue,
	}
}

// DescribeMandatoryString describes a mandatory string parameter
func (c *Config) DescribeMandatoryString(name string, description string) {
	c.describe(name, description, true, "")
}

// DescribeOptionalString describes an optional string parameter with a default value
func (c *Config) DescribeOptionalString(name string, description string, defaultValue string) {
	c.describe(name, description, false, defaultValue)
}

// DescribeMandatoryInt describes a mandatory integer parameter
func (c *Config) DescribeMandatoryInt(name string, description string) {
	c.describe(name, description, true, "")
}

// DescribeOptionalInt describes an optional integer parameter with a default value
func (c *Config) DescribeOptionalInt(name string, description string, defaultValue int) {
	c.describe(name, description, false, strconv.Itoa(defaultValue))
}

// DescribeMandatoryBool describes a mandatory boolean parameter
func (c *Config) DescribeMandatoryBool(name string, description string) {
	c.describe(name, description, true, "")
}

// DescribeOptionalBool describes an optional boolean parameter with a default value
func (c *Config) DescribeOptionalBool(name string, description string, defaultValue bool) {
	stringDefaultValue := "false"
	if defaultValue {
		stringDefaultValue = "true"
	}
	c.describe(name, description, false, stringDefaultValue)
}

// PrintUsage prints the usage information
func (c *Config) PrintUsage() {
	fmt.Println("Set the following mandatory environment variables:")
	fmt.Println("")

	for _, parameter := range c.Parameters {
		if parameter.Mandatory {
			fmt.Printf("   %-20s %s\n", c.Prefix+parameter.Name, parameter.Description)
		}
	}

	fmt.Println("")

	fmt.Println("Set the following optional environment variables to overwrite the default values")
	fmt.Println("")
	for _, parameter := range c.Parameters {
		if !parameter.Mandatory {
			fmt.Printf("   %-20s %s (Default: %s)\n", c.Prefix+parameter.Name, parameter.Description, parameter.DefaultValue)
		}
	}
}

// Parse parses the existing configuration and fails if mandatory parameters are missing
func (c *Config) Parse() {
	if len(os.Args) >= 2 {
		switch strings.ToLower(os.Args[1]) {
		case "help", "--help", "-help", "-h":
			c.PrintUsage()
			os.Exit(-1)
		}
	}

	failed := false
	for _, parameter := range c.Parameters {
		if parameter.Mandatory {
			result := os.Getenv(c.Prefix + parameter.Name)
			if result == "" {
				fmt.Printf("Error: Mandatory environment variable %s not set!\n", c.Prefix+parameter.Name)
				failed = true
			}
		}
	}

	if failed {
		fmt.Println("")
		c.PrintUsage()
		os.Exit(-1)
	}
}

// GetString returns a string parameter
func (c *Config) GetString(name string) string {
	parameter := c.Parameters[name]
	result := os.Getenv(c.Prefix + name)
	if result == "" {
		return parameter.DefaultValue
	}
	return result
}

// GetInt returns an integer parameter
func (c *Config) GetInt(name string) int {
	stringResult := c.GetString(name)
	intResult, err := strconv.Atoi(stringResult)

	if err != nil {
		fmt.Printf("Error: Value '%s' of environment variable %s is not an integer!\n", stringResult, name)
		os.Exit(-1)
	}

	return intResult
}

// GetBool returns a boolean parameter
func (c *Config) GetBool(name string) bool {
	stringResult := c.GetString(name)
	switch strings.ToLower(stringResult) {
	case "true", "1", "yes", "on":
		return true
	default:
		return false
	}
}
