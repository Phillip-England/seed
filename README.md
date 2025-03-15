# seed
Generate go skeleton projects quickly.

## Installation
Install on your `GOPATH` using:
```bash
go install github.com/Phillip-England/seed@latest
```

If you prefer to build from source, with go installed on your system, clone the repo, install deps, then build:
```bash
git clone https://github.com/Phillip-England/seed
cd seed
go mod tidy
go build -o seed
```

## Usage
If you run `seed plant` without any further args, seed will assume you want the project generated in your current directory.

To specify, a target directory, just pass the name of the directory you'd like to generate. This will generate an app directory containing our project (assuming ./app doesn't already exist).
```bash
seed plant app
```
