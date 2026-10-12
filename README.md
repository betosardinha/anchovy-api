# Anchovy Api

[![go version](https://img.shields.io/github/go-mod/go-version/betosardinha/anchovy-api?logo=Go&logoColor=00ADD8&&color=00ADD8)](#)
[![release](https://img.shields.io/github/v/release/betosardinha/anchovy-api)](https://github.com/betosardinha/anchovy-api/releases)
[![check runs](https://img.shields.io/github/check-runs/betosardinha/anchovy-api/main?logo=Github)](#)

Backend project management API focused on task organization and activity tracking

## Install

The project is full docker friendly, and local go version and packages are not necessary. For image build the application you can run standard docker compose command:

```bash
docker compose -f docker-compose.dev.yml build
```

The project has a full makefile that can be used too, so building the project can also be:

```bash
make build
```

## Usage

Run API container: `make up`

Stop all API docker containers: `make down`

## Maintainer

[@betosardinha](https://github.com/betosardinha)

## License

[MIT](LICENSE) 2026 © Alberto Sardinha