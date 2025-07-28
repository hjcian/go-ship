# Design doc
- ✅ basic functionalities
    - ✅ read config
        - name: "redis"
        - tag_pattern: "v*"
        - registry: "dockerhub"
    - ✅ get latest tag from remote
    - ✅ get local running container's tag
    - ✅ compare the tag
        - if different, pull the latest image and restart the container
        - if same, do nothing
    - ✅ restart the container
- 🚧 pulling based on the pattern, not just the latest
    - ✅ fetch latest is pattern not specified
    - 👉 fetch tags based on the patterns
- 🚧 decide the behavior if there are multiple containers with the same image
- 🚧 a simple static web page to show the status of the containers
- 🚧 support AWS, Azure, GCP container registry
- [ ] tests
- [ ] CI/CD

WatchedImage
- name
- tag_patterns
- registry



## Compatibility issues

docker version alignment

❎

```bash
failed to create new container: "specify mac-address per network" requires API version 1.44, but the Docker daemon API version is 1.41
```

```bash
$ docker version
Client:
 Cloud integration: v1.0.24
 Version:           20.10.14
 API version:       1.41
 Go version:        go1.16.15
 Git commit:        a224086
 Built:             Thu Mar 24 01:49:20 2022
 OS/Arch:           darwin/amd64
 Context:           default
 Experimental:      true

Server: Docker Desktop 4.8.2 (79419)
 Engine:
  Version:          20.10.14
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.15
  Git commit:       87a90dc
  Built:            Thu Mar 24 01:46:14 2022
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.5.11
  GitCommit:        3df54a852345ae127d1fa3092b95168e4a88e2f8
 runc:
  Version:          1.0.3
  GitCommit:        v1.0.3-0-gf46b6ba
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```
