<!-- $theme: gaia -->

# Introduction to Docker

---

# Agenda

- Docker basics
- Why Docker
- Creating our first Docker container
- A look behind the scenes

--- 
# Docker basics

---
# What is a container

- Not a new idea... 
- Zones, Jails
- Containers are not a real thing!
- Combination of Linux's namespaces and CGroups
---
# Architecture
![](/Users/raffo/Downloads/dockerjm1.webp)

---

# Why Docker
---

# Why Docker
- A simple way to package applications in a self contained fashion
- Run in the same way on laptop and server
- Reproduceable cases
---

# Terminology
- Dockerfile: a file with a description on how to build a docker image
- Docker image: a built container
- Image tag: version of an image

---
# Your first docker container

- Create a dockerfile 
- Build a docker image

---

# Dockerfile 

```
FROM ubuntu

RUN apt-get update && apt-get install -y vim curl
```
---
# Let's build the image

```
mkdir testfolder && cd testfolder
docker build -t first-image .
docker run -it first-image /bin/bash
```
---

# Let's edit our files
```
docker run -v $(pwd):/testfolder -it first-image /bin/bash
cat "hello world" > /testfolder/foo
exit
cat foo
```

---
# A closer look
```
docker images
docker ps -a
docker pull golang
docker inspect first-image
docker system prune
```
---
# Videos to watch

[What are container made from?](https://www.youtube.com/watch?v=sK5i-N34im8)




