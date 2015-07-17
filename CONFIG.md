The config file is a plain json file, with the structure defined in the
`datastructures.go` file, included in the `src` folder.

The root element is a `hash` object, with 3 child objects:
 + `interval` is a `int` parameter indicating the time (in seconds) to wait
 between checks.
 + `irc` is a `hash` with five childs indicating the details used to connect to
 irc. Support for multiple irc servers/channels is planned and a relative priority.
  1. `nick` is a `string` indicating the nickname Yggdrasil should have on irc
  2. `realname` is a `string` indicating the realname Yggrasil should have on irc
  3. `server` is a `string` indicating the hostname (or IP address) to connect to
  4. `port` is a `int` indicating the port to connect to
  5. `channel` is a `string` indicating the channel to join and to write in
 + `services` is a `array` of `hashes` indicating the services Yggdrasil is
 supposed to test. The hashes should be specified with the following structure:
  1. `host` is a `string` containing the hostname/IP address of the service to test
  2. `port` is a `int` indicating the port the service listens on
  3. `proto` is a `string` indicating the protocol of the service
  4. `type` is a `string` indicating the service type, for available values check
  the `ServiceType` type in `datastructures.go`
  5. `name` is a `string` indicating the name the service should have in logs and
  the various human readable messages in Yggdrasil
