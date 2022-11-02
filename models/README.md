# URCAPX API

This document describes the API contract between URCaps and URService. 

## How to read through swagger documentation
### Step 1: Enter docker container in interactive way
```bash
make interactive-docker
```
### Step 2: Install swagger
```bash
make install-swagger
```
### Step 3: Serve Swagger
```bash
make server-swagger
```

## How to use auto generation server stub
### Step 1: Generate mod file
```bash
go mod universal-robots.com/urservice
```
### Step 2: Generate server
```bash
make generate-server
```
### Step 3: Run server
```bash
make run
```

