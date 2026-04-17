# startAutoscale

A simple utility to update an AWS Auto Scaling Group's capacity to 1 (both min, max, and desired capacity).

## Install

Make sure you have Go installed, then clone the repository and build:

```sh
make get
make build
```

Optionally, copy the built binary to a directory in your PATH:

```sh
mkdir -p ~/bin
cp startAutoscale ~/bin
```

## Usage

You can run the utility specifying the target Auto Scaling Group name using the `-group` flag:

```sh
startAutoscale -group "auto scale group name"
```

### Options

```
  -group string
        The name of the autoscale group you wish to set to capacity 1 (Required)
  -region string
        AWS Region (default "ap-southeast-2")
  -awscfg string
        Config file to load credentials from (default "~/.aws/credentials")
  -awsprofile string
        AWS Profile to use (default "default")
```
