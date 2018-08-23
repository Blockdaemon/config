Package config provides a simple way to configure an app with environment variables.

It supports:

- Int, String and Bool types
- Auto-documentation (just call the app with `--help`)
- Default values
- Checks if mandatory environment variables are set

Example:

    config := new(Config)
	config.SetPrefix("EXAMPLE_")
    config.DescribeInt("PORT", "The port to listen to", false, 1234)
    config.DescribeString("HOST", "The host to listen to", false, "0.0.0.0")
    config.DescribeBool("DEBUG", "Start in debug mode", false, true)
    config.Parse()

    host := config.GetString("HOST")
    port := config.GetInt("PORT")
    debug := config.GetBool("DEBUG")

	fmt.Printf("Host: %s, Port: %d, Debug: %t\n", host, port, debug)

The help for the example looks like this:

    $ ./example --help

    Use the following environment variables:
    EXAMPLE_PORT - The port to listen to (Optional, Default: 1234)
    EXAMPLE_HOST - The host to listen to (Optional, Default: 0.0.0.0)
    EXAMPLE_DEBUG - Start in debug mode (Optional, Default: true)