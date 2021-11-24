# RIO Tachograph Downloader

CLI for downloading DDD-Files from [rio.cloud](https://rio.cloud) to a local folder. This can be helpful to integrate third-party-software or for backup/archival tasks.

**Features:**
- Downloads all *.DDD files from the Tachograph Archive in the Compliant module of [rio.cloud](https://rio.cloud).
- Files that are already present are not downloaded again, only new files.
- Files are organized by *type*, *year* and *month* into sub-folders.
- Downloads are rate-limited to a given number of files per second in order to be nice to the API.

The downloader uses the [Rio Partner API](https://rio.cloud/de/rio-fuer/partner). It is required to get API-Keys from the RIOs Team in order to be able to use the downloader.

## Hot to install

**Build from source:**

    git clone https://github.com/t-muehlberger/rio-tachograph-downloader.git
    cd ./rio-tachograph-downloader
    go mod download
    go build -o rio-tachograph-downloader *.go

**Install via `go get`**

    go get -u github.com/t-muehlberger/rio-tachograph-downloader

## How to run

The following environment variables should be defined:

    API_INTEGRATION_ID=...
    API_CLIENT_ID=...
    API_CLIENT_SECRET=...
    TARGET_DIR="." (optional)

Run the binary to download the files.

    ./rio-tachograph-downloader

