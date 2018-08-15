<!-- $theme: gaia -->

# Introduction to Docker

---

# This is a kind of short intro to docker. I've seen better tutorials... but they are about 700 slides. Go check it out if you want: https://container.training

--- 
# Docker basics

---
# What is a container

- Not a new idea... 
- Zones, Jails
- "Containers" are linux only
- Containers are not a real thing!
	- Nothing in the kernel speaks about "containers"
	- Combination of Linux's namespaces and CGroups
	
---
# Architecture
![VMs and containers](http://cdn.rancher.com/wp-content/uploads/2017/02/16175231/VMs-and-Containers.jpg)

---

# Why Docker
---

# Why Docker
- A simple way to package applications in a self contained fashion
- Run in the same way on laptop and server...
- ... and have reproduceable deployments
---
# Let's get our hands dirty
---
# Docker quickstart
```
docker run -it ubuntu /bin/bash
curl
apt-get update
apt-get install -y curl
curl
exit
```
---
# Terminology
- Dockerfile: a file with a description on how to build a docker image
- Docker image: a built container
- Image tag: version of an image
---

# Dockerfile commands
- `RUN` command (non interactive) executed during the build
- `CMD` default command to run when no default is given
- `ENTRYPOINT` base command to run

---

# More directives
- `COPY`: copies a file to the container at build time 
```
COPY . . 
```
---
# Let's create a first Dockerfile
---

# Dockerfile 

```
FROM ubuntu

RUN apt-get update && apt-get install -y vim curl
```
- Let's add this to a file named `Dockerfile`
---

# Let's build the image

```
mkdir testfolder && cd testfolder
touch Dockerfile
# Add content from the previous slide
docker build -t first-image .
docker run -it first-image /bin/bash
figlet hello
```
---
# Let's install the package
```
apt-get update
apt-get install -y figlet
figlet hello
```
---
# Again..
```
docker run -it first-image /bin/bash
figlet hello
```
---
# What did just happen? 

- `figlet` is not available as it is not in the docker image.

---

# Layers
- Each command is a layer. 
- `RUN rm ...` will remove things... which will be still available in the previous layer!
- `RUN apt-get install XX && ap-get remove XX` would do (single layer)
---

# Let's edit our local files

```
docker run -v $(pwd):/testfolder -it first-image /bin/bash
echo "hello world" > /testfolder/foo
exit
cat foo
```

---
# A closer look to Docker
- There are many commands and functionalities, i.e.:
```
docker images
docker ps -a
docker pull golang
docker inspect first-image
docker system prune
```
---
# Tagging images
- Very important to know what we are running
- Default tag `latest`
- Can be specified at build time
- Don't use `latest` ! 
```
docker tag <newImageId> figlet
```
---

# Layers, layers, layers, ...
```
FROM ubuntu
RUN apt-get install curl
RUN apt-get install vim 
RUN apt-get install ...
...
```
- Each instructions creates a layer
- Less layers, easier caching
- Layers used wrong can expose secrets
---

# Improving our dockerfiles: multistage
- Multistage builds allows to have multiple stage
- One stage for building, one for the final image
```
FROM golang AS builder
RUN ...
FROM alpine
COPY --from=builder /go/bin/mylittlebinary /usr/local/bin/
```
---
# Where do we store the images? 
- Docker registry: where the images are stored
- Can be self hosted or public
- Docker hub ("default")
- ECR, GCR, ... 
---
# Curiosities
- How does Docker work on a mac? 
---
# Things we haven't covered
- Docker networking
- Compose, volumes
---
# Remarks
- Don't store configuration in an image
- **Never** store secrets inside an image
---
# Bonus
- Find a way to download these slides...
- ... start from `docker pull x0rg/docker-slides`


# Links
- [container.training intro to docker](http://container.training/intro-selfpaced.yml.html#1)

- [What are container made from?](https://www.youtube.com/watch?v=sK5i-N34im8)




