# Warp cli

`warp` is a cli tool for the [Warp-charger](https://www.warp-charger.com).
The cli tool brings the http api to the terminal.


## Command overview
| Command | Description |
| --- | --- |
| `warp info` | Get information about the warp charger |
| `warp info version` | Get the version of the warp charger |
| `warp info update` | Check if an warp charger update is available |
| `warp info name` | Get the name of the warp charger |
| `warp infor display-name` | Get the display name of the warp charger |
| `warp info modules` | Get the modules of the warp charger |
| `warp info features` | Get the features of the warp charger |
| `warp meter values` | Get the meter values of the warp charger |
| `warp users list` | Get the users of the warp charger |
| `warp charge-tracker` | Get information about the charge tracker |
| `warp charge-tracker log` | Get the charge tracker log (csv) |
| `warp version` | Get the version of the warp cli |

Each command has a help page, which can be accessed with the `-h` or `--help` flag.
The help page prints the usage of the command and the available flags.

## Configuration
The configuration file is located at `~/.warp.yaml`.

#### Default configuration
```yaml
csv:
    comma: ; # separator for csv
    header: true # add header to csv
date_time:
    time_format: 15:04:05 02-01-2006 # date time format
    time_zone: Europe/Berlin # time zone to print the date time
power:
    price: "0.35" # price per kWh
```

#### Build warp cli

To build a executable for your system, your have to do the following steps:
- Install go 1.19 or higher
- Download the source from github https://github.com/HappyTobi/warp/archive/refs/heads/main.zip,
and extract it.

Run the following commands in the extracted folder:
```bash
make release
cd build/
```

## Example usage

### Info version
```console
$ warp info name -c "http://mywarp.ip -u "username" -p "password"

{
 "display_type": "WARP Charger Pro 22kW +NFC",
 "name": "warp-UTD",
 "type": "warp",
 "uid": "UTD"
}
```

### Charge tracker
```console
$ warp charge-tracker log -c "http://mywarp.ip -u "username" -p "password"

[
 {
  "Time": "2023-03-06T18:14:00Z",
  "User": "happyTobi",
  "PowerMeterStart": 111.929,
  "PowerMeterEnd": 139.224,
  "Duration": "13:06:58"
 }
]
```

## Warp Charger information

[Product Page](https://www.warp-charger.com)

[API Documentation](https://www.warp-charger.com/api.html)