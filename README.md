# alfred-boxpn

A https://boxpn.com/ alfred workflow to start and stop vpn connections.

## Installation

1. Download the latest release and install it in Alfred.
2. Update `PATH` variable with the location of `openvpn` and `osascript`, for example
  ```
  echo "$(dirname $(which openvpn)):$(dirname $(which osascript))"
  ```

## Usage

The keyword is `vpn` and that will list the available profiles.

## Stopping the vpn

Use the subcommand `close`

