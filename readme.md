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
Usage of ./firewalla:
  -debug
        Debug mode
  -help
        Display help menu
  -hours int
        (Optional) Flowlog hours to retrieve (default 24)
  -limit int
        (Optional) Results limit (default 200)
  -mode string
        Mode to run the script (devices, alarms, flowlogs) (default "flowlogs")
```


## CLI Examples
```bash
# Get all devices (write to devices.csv)
$ firewalla -mode devices

# Get all alarms (write to alarms.csv)
$ firewalla -mode alarms

# Get all flowlogs for last X hours (default 24) (write to flowlogs/ folder)
$ firewalla -mode flowlogs
$ firewalla -mode flowlogs -hours 3
$ firewalla -mode flowlogs -hours 48 -debug
```
