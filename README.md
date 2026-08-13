## Overview

WhatsApp integration service for the ru-menu ecosystem. Consumes restaurant menu updates from the scraper and delivers them to WhatsApp users via automated messages and media.

## Architecture

Part of the ru-menu pipeline: scraper -> whatsapp

- Receives menu events containing restaurant data, meals and meal changes
- Authenticates with WhatsApp using credentials stored in SQLite
- Sends formatted menu messages and images to registered phone numbers
- Runs as AWS Lambda function with S3 database persistence

## Features

- Event-driven menu delivery to WhatsApp
- Multi-meal messaging with formatting
- Image delivery for menu photos
- Persistent authentication storage (SQLite with S3 backup)
- Local development mode with event testing

## Tech Stack

- Go 1.25.5
- AWS Lambda
- AWS S3
- Whatsmeow (WhatsApp library)
- SQLite (state management)

## Project Structure

- `handler.go` - Lambda event handler
- `whatsapp.go` - WhatsApp client operations
- `s3.go` - S3 integration
- `models.go` - Data structures
- `config.go` - Configuration
- `formatter.go` - Message formatting
- `event.json` - Local test event

## Automated Updates

WhatsApp web client libraries require frequent maintenance. The Whatsmeow dependency is automatically updated weekly via GitHub Actions to stay compatible with WhatsApp protocol changes.

- Runs every Monday at 3 AM UTC
- Updates to latest Whatsmeow version
- Auto-commits and pushes changes if updates found

Note: WhatsApp web clients frequently break due to protocol updates and reverse-engineering changes. Regular dependency updates are critical to maintain functionality.

## Repository

Connected to https://github.com/iugx/ru-menu
