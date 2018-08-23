Package config provides a simple way to configure an app with environment variables.

It supports:

- Int, String and Bool types
- Auto-documentation (just call the app with `--help`)
- Default values
- Checks if mandatory environment variables are set

For an example see: https://github.com/Blockdaemon/config/blob/master/example/main.go

The help for the example looks like this:

    $ ./example --help

    Set the following mandatory environment variables:

        EXAMPLE_PEERS        The amount of peers to connect to
        EXAMPLE_ID           A unique id for this app
        EXAMPLE_ACCEPT       Accept the terms and conditions

    Set the following optional environment variables to overwrite the default values

        EXAMPLE_PORT         The port to listen to (Default: 1234)
        EXAMPLE_HOST         The host to listen to (Default: 0.0.0.0)
        EXAMPLE_DEBUG        Start in debug mode (Default: false)