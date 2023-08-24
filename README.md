# ACScraper

ACScraper is a simple proxy server tool designed to intercept and extract the API token from the AC Infinity app. This API token is utilized to fetch status and sensor readings from the AC Infinity controller used in AC Infinity grow tents. Once the token is extracted, it can be further used in the TentStatter application to record values to a text file or integrated into custom code to gather metrics.

## Features

- Captures and displays the API token from the AC Infinity app's request headers.
- Displays the `User-Agent` header value.
- Lightweight and easy-to-run.
- Acts as a transparent proxy, forwarding all requests and responses without modifying them.

## Prerequisites

- Go (Golang) installed on your machine.

## Setup and Usage

1. **Clone the repository**:

   ```bash
   git clone https://github.com/YOUR_GITHUB_USERNAME/ACScraper.git
   cd ACScraper
   ```

2. **Run the proxy server**:

   ```bash
   go run main.go
   ```

3. Once the server is running, you'll see the IP address and port where it's listening. For example:

   ```
   Starting proxy server on IP: 192.168.1.10 and port: 8080
   On your phone, set your proxy server to: 192.168.1.10 with port: 8080
   ```

4. **Configure your phone**:

   Set your phone's proxy settings to the IP address and port displayed in the console output.

5. **Capture the token**:

   Open the AC Infinity app and perform an action that triggers a request to `http://www.acinfinityserver.com/api/`. The proxy server will capture and display the `token` and `User-Agent` headers from the request.

6. **Use the token**:

   Integrate the token into the TentStatter application or any custom code to start gathering data from the AC Infinity controller.
