# ACScraper

ACScraper is a simple proxy server tool designed to intercept and display the API token from requests made by the AC Infinity app. This token can be used in conjunction with apps such as [TentStatter](https://github.com/dwot/TentStatter) to collect and log data from AC Infinity grow tents.

## Features

- Intercepts requests to `http://www.acinfinityserver.com/api/` to capture the API token and user-agent string.
- Simple to use with clear console instructions for proxy settings.
- Supports multiple platforms: Windows, Linux, and macOS.

## Getting Started

### 1. **Download and Install**

Choose the appropriate build for your operating system from the [Releases](https://github.com/dwot/ACScraper/releases) section:

Also, download the corresponding `.md5` file to verify the integrity of your download.

### 2. **Verify Download (Optional)**

Before unpacking, you can verify the integrity of your download by comparing the MD5 checksum of your downloaded file with the one provided:

```bash
# For Linux/macOS
md5sum acscraper-v0.0.1-YOUR_OS_HERE.tar.gz

# For Windows (in PowerShell)
Get-FileHash -Algorithm MD5 .\acscraper-v0.0.1-windows-amd64.zip
```

The output should match the contents of the corresponding `.md5` file.

### 3. **Unpack and Run**

Uncompress the downloaded file and run the ACScraper executable.

On first launch, the tool will display your local IP address and prompt you to set your phone's proxy server to the displayed IP with port 8080.

### 4. **Capture the Token**

With the proxy settings adjusted on your phone, initiate a request from the AC Infinity app. ACScraper will display the captured token and user-agent string in the console.

