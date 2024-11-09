# Snippy

Snippy is a simple one binary tool to manage your snippets.

## Screenshots

| Dark                              | Light                              |
| --------------------------------- | ---------------------------------- |
| ![alt text](screenshots/dark.png) | ![alt text](screenshots/light.png) |

## Installation

```bash
# Requires sqlite3 library
$ pnpm build
$ ./snippy

# or via docker (local)
$ docker build -t snippy .
$ docker run -it -p 8080:8080 -v ./db:/app snippy

# or via docker (remote)
$ docker run -it -p 8080:8080 -v ./db:/app ghcr.io/alindesign/snippy
```

## Features

- [x] Web based UI
- [x] Dark and Light mode
- [x] Snippets
  - [x] Create
  - [x] Read
  - [x] Update
  - [x] Delete
- [x] Server
- [x] Docker
- [ ] CLI Client
