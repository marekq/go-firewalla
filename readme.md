go-firewalla
============

Download device details, flowlogs and alarms from the Firewalla MSP API. The results are stored to a local CSV file. 

## Installation

- Edit the `config.json` file with your Firewalla MSP API key and URL:
```json
{
    "token": "token",
    "url": "https://<url>.firewalla.net/v1/"
}
```
- Run `go build` to build the binary (`firewalla`)
- Run `chmod +x firewalla` to make the binary executable
- Now you can one of the following commands:
```bash
# Get all devices (devices.csv)
$ firewalla -mode devices

# Get all alarms (alarms.csv)
$ firewalla -mode alarms

# Get all flowlogs for last X hours (default 1) (flowlogs.csv)
$ firewalla -mode flowlogs -hours 2
```
