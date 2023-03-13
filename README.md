# Warp cli

`warp` is a cli tool for the [Warp-charger](https://www.warp-charger.com).
The cli tool brings the http api to the terminal.

## Examples

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


## Warp Charger information

[Product Page](https://www.warp-charger.com)

[API Documentation](https://www.warp-charger.com/api.html)