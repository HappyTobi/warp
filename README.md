# Warp cli

`warp` is a cli tool to port the http api to the command line.

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

### charge tracker
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

- Download the source from github https://github.com/HappyTobi/warp/archive/refs/heads/main.zip,
and extract it.

```bash
make release
cd build/
```


## Warp Charger information

[Product Page](https://www.warp-charger.com)

[API Documentation](https://www.warp-charger.com/api.html)