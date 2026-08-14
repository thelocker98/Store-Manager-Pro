# Store-Manager-Pro

## Building
Here is the information on how to build this with docker and with go cli.

### Go CLI
To build this into a deployable binary use this command. Make sure to update the build version

```bash
go build -ldflags "-X gitea.locker98.com/locker98/Store-Manager-Pro/utils.Version=v0.8.1 -X gitea.locker98.com/locker98/Store-Manager-Pro/utils.Commit=$(git rev-parse --short HEAD)"
```

### Docker
Here are the commands to build the storemanagerpro-server docker image.

```bash
# Build for multi platform deployment
docker buildx build --platform linux/arm64,linux/amd64 --build-arg VERSION=v0.8.1 --build-arg COMMIT=$(git rev-parse --short HEAD) -t gitea.locker98.com/locker98/store-manager-pro:latest -t gitea.locker98.com/locker98/store-manager-pro:v0.8.1 . --push
```
