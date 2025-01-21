# Slack Bot Age Calculator

This is a simple Slack bot written in Go that listens for user commands and calculates the user's age based on the provided year of birth.

## Features

- Listens for the command `my yob is <year>`.
- Calculates the user's age by subtracting the provided year from 2025.
- Handles invalid input gracefully.
- Logs command events.

## Prerequisites

Make sure you have the following installed:

- Go (latest version recommended)
- A Slack workspace
- A Slack bot with permissions

## Setup

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

   ## Code Explanation

### Environment Variables
The bot reads the Slack bot token and app token from environment variables. This allows for secure management of sensitive information without hardcoding it into the source code.

### Command Definition
The bot registers a command to calculate age based on the provided year of birth. The command is defined in a way that it captures the year input from the user and processes it to compute the age.

### Command Event Logging
The bot logs incoming command events to the console. This feature is useful for monitoring and debugging, as it provides insights into the commands being issued and the corresponding timestamps.

### Error Handling
The bot handles invalid input gracefully by prompting the user to provide a valid year. If the input cannot be parsed as a number, the bot responds with an appropriate error message, ensuring a smooth user experience.