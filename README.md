# Go Embed Single-Page Application

Although there are many alternatives to deploying a single-page application, you might find a situation where you need to deploy it in an isolated environment or just portability concern.

The example is using SvelteKit to generate a single-page application and embed it with Golang with Docker multi-stage build for smaller image.

## Tech

- [SvelteKit](https://kit.svelte.dev/ "SvelteKit")
- [Go](https://go.dev/ "Golang")
- [Echo](https://echo.labstack.com/ "Echo Framework")
- [Fiber](https://gofiber.io/ "Fiber Framework")
- [Docker](https://www.docker.com/ "Docker")

## Build the image

Go HTTP standard library

```bash
make build
```

Echo framework

```bash
make build APP_NAME=echo
```

Fiber framework

```bash
make build APP_NAME=fiber
```

## Run the application

The default port is `5050`, configure a runnable app port with `APP_PORT=xxxx`.

Go HTTP standard library

```bash
make run
```

Echo framework

```bash
make run APP_NAME=echo APP_PORT=5051
```

Fiber framework

```bash
make run APP_NAME=fiber APP_PORT=5052
```

## Update
Reduce the binary size around 50%-70% by using [UPX](https://upx.github.io/).

Before:
```
REPOSITORY   TAG       IMAGE ID       CREATED              SIZE
spa          fiber     ddbdf411532d   7 seconds ago        8.98MB
spa          echo      f0f144aecd27   53 seconds ago       10.9MB
spa          http      1b4f9bd0632b   About a minute ago   7.29MB
```
After:
```
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
spa          fiber     6a2d174a176b   6 minutes ago   5.23MB
spa          echo      a82aa5113481   7 minutes ago   5.79MB
spa          http      4248e779a14f   8 minutes ago   4.49MB
```