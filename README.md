# Ubersnap Backend Engineer Test

> This project generate using [Go Bone Boilerplate](https://github.com/Zainal21/go-bone) [Go Framework]

## Getting started

This is built on top of [Go Fiber](https://docs.gofiber.io) Golang Framework.

## Dependencies

There is some dependencies that we used in this skeleton:

- [Go Fiber](https://docs.gofiber.io/) [Go Framework]
- [Viper](https://github.com/spf13/viper) [Go Configuration]
- [Cobra](https://github.com/spf13/cobra) [Go Modern CLI]
- [Logrus Logger](https://github.com/sirupsen/logrus) [Go Logger]
- [Goose Migration](https://github.com/pressly/goose) [Go Migration]
- [Gobreaker](https://github.com/sony/gobreaker) [Go Circuit Breaker]
- [GoCv](https://gocv.io/x/gocv) [Go OpenCV Wrapper]

## Requirement

- Golang version 1.21 or latest

## Usage

### Installation

clone this repository

```bash
git clone github.com/Zainal21/Ubersnap-backend-test.git
```

install required dependencies

```bash
make install
```

Setup env (Only Port, this project not connect to Database )

```bash
cp .env.example .env
```

### Run Service (Development Mode)

run current service after all dependencies installed

```bash
make start-http
```

### Build to Binary Executable (Production Build)

run current service after all dependencies installed

```bash
make build
```

### Run Service (Production Running)

run current service build binary project

```bash
make run-http
```

## Service & Postman Collection

- include Dockerfile Build if using container
- healthy check
- Convert image files from PNG to JPEG.
- Convert image files from PNG to JPEG.
- Resize images according to specified dimensions (width and height).
- Compress images to reduce file size while maintaining reasonable quality.

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/9050639-37a36603-d69a-4ad3-84aa-3a5500d8e4ee?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D9050639-37a36603-d69a-4ad3-84aa-3a5500d8e4ee%26entityType%3Dcollection%26workspaceId%3D57cb0c84-a57a-424e-a59e-2afd4c4a1a7d)
