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
