# seed
A cli for generating skeleton projects in go.

## Installation
Install on your `GOPATH` by running:

```bash
go install github.com/Phillip-England/seed@latest
```

If you'd rather install from source, clone the repo, install deps, build/rename binary, execute seed:

```bash
git clone https://github.com/Phillip-England/seed
cd seed
go mod tidy
go build
mv ./main seed
seed plant
```

