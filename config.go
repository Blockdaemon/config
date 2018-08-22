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
}

// DescribeString describes a string parameter
func (c *Config) DescribeString(name string, description string, mandatory bool, defaultValue string) {
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

// DescribeInt describes an integer parameter
func (c *Config) DescribeInt(name string, description string, mandatory bool, defaultValue int) {
	c.DescribeString(name, description, mandatory, strconv.Itoa(defaultValue))
}

// DescribeBool describes a boolean parameter
func (c *Config) DescribeBool(name string, description string, mandatory bool, defaultValue bool) {
	stringDefaultValue := "false"
	if defaultValue {
		stringDefaultValue = "true"
	}
	c.DescribeString(name, description, mandatory, stringDefaultValue)
}

// PrintUsage prints the usage information
func (c *Config) PrintUsage() {
	fmt.Println("\nUse the following environment variables:")

	for _, parameter := range c.Parameters {
		details := ""
		if parameter.Mandatory {
			details += "Mandatory"
		} else {
			details += "Optional"
		}
		if parameter.DefaultValue != "" {
			details += ", Default: " + parameter.DefaultValue
		}

		fmt.Println("  " + parameter.Name + " - " + parameter.Description + " (" + details + ")")
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
			result := os.Getenv(parameter.Name)
			if result == "" {
				fmt.Printf("Error: Mandatory environment variable %s not set!\n", parameter.Name)
				failed = true
			}
		}
	}

	if failed {
		c.PrintUsage()
		os.Exit(-1)
	}
}

// GetString returns a string parameter
func (c *Config) GetString(name string) string {
	parameter := c.Parameters[name]
	result := os.Getenv(name)
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
