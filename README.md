# Controlling SwitchBot Devices with Golang

This is a sample code that demonstrates how to control SwitchBot devices using Golang. It generates a signature based on your SwitchBot token and secret, and then sends a command to turn on a SwitchBot device.

## Prerequisites

* Go 1.16 or later
* A SwitchBot account with at least one device registered
* Your SwitchBot token and secret values
* The device ID of the SwitchBot device you want to control
* SwitchBot device (e.g., SwitchBot)
* SwitchBot Hub (e.g., Hub Mini)

Before running the application, you need to have a SwitchBot device and a SwitchBot Hub. The SwitchBot Hub is required to control the SwitchBot device remotely via the SwitchBot API.

## Installation

1. Create a `config.yml` file in the same directory as `main.go` with the following contents:

```
token: your_token
secret: your_secret
deviceID: your_device_id
```

2. Replace `your_token`, `your_secret`, and `your_device_id` with your actual values.

## Usage

1. Run the program:

```
go run main.go
```

2. The program will send a command to turn on the specified SwitchBot device. Check the output for the response.

## Notes
You can retrieve the deviceId for your SwitchBot device by sending a GET request to the SwitchBot API using your token

Here's how to do it:

Open your terminal and enter the following command, replacing `your_token` with your SwitchBot token

```
curl -s -X GET "https://api.switch-bot.com/v1.0/devices" -H Authorization:your_token" | jq
```
