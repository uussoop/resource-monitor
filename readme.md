# Simple System Monitoring with Telegram Notifications

This is a simple Go application for monitoring system resources (CPU, memory, and disk) and sending notifications via Telegram when any of these resources reach specified thresholds.

## Prerequisites

Before running this application, you need to set two environment variables:

- `TELEGRAM_ADMINID`: Your Telegram admin user or group chat ID.
- `TELEGRAM_APITOKEN`: Your Telegram Bot API token.

These variables are required for the application to send notifications via Telegram. Make sure to set them in your environment before running the application.

## Usage

### Start Monitoring

You can start the system monitoring by calling the `SimpleMonitoring` function. It takes three parameters:

- `cpu`: The CPU utilization threshold (in percentage) to trigger a notification.
- `memory`: The memory usage threshold (in percentage) to trigger a notification.
- `disk`: The disk usage threshold (in percentage) to trigger a notification.

Here's an example of how to start monitoring with default values of 80% for each resource:

```go
api.SimpleMonitoring(80, 80, 80)
```

### Customize Thresholds

If you want to customize the resource thresholds, simply adjust the values passed to `SimpleMonitoring`. For example, to monitor CPU at 90%, memory at 70%, and disk at 60%, you can do the following:

```go
api.SimpleMonitoring(90, 70, 60)
```

### Notifications

The application continuously monitors the specified resources and sends Telegram notifications when any of them exceed the specified thresholds. Notifications will be sent to the Telegram user or group chat specified by `TELEGRAM_ADMINID`.

## Monitoring Logic

The `Monitor` function contains the logic for monitoring the system resources. It checks the CPU utilization, memory usage, and disk usage at regular intervals and sends notifications when any of these resources exceed the thresholds.

## Dependencies

This application relies on the following dependencies:

- `resourceapi`: The resource monitoring functions are expected to be provided by an external package or module. Ensure that these functions are implemented and available.

## License

This application is provided under an open-source license. You are free to modify and use it according to your requirements.
