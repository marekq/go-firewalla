go-firewalla
============

Download device details and alarms from the Firewalla MSP API. The results are stored to a local CSV file. 

## Installation

- Edit the `config.json` file with your Firewalla MSP API key and URL:
```json
{
    "token": "token",
    "url": "https://<url>.firewalla.net/v1/"
}
```
- Run `go build` to build the binary.
- Now you can one of the following commands:
```bash
# Get all devices
./go-firewalla devices

# Get all alarms
./go-firewalla alarms

# Get all devices and alarms
./go-firewalla all
```
