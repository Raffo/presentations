<!-- $theme: gaia -->

# Introduction to Docker

---

# Agenda

- Docker basics
- Why Docker
- Creating our first Docker container
- A look behind the scenes
---

# This is a kind of short intro to docker. I've seen better tutorials... but they are about 700 slides. Go check it out if you want: https://container.training


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
# Dockerfile
- You can add things there if you really need them 
--- 

---
# What did just happen? 

- Our container is now in a stopped state.

- It still exists on disk, but all compute resources have been freed up.

- We will see later how to get back to that container.
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
# Tagging images
- Default tag `latest`
- Can be specified at build time
- Don't use `latest` ! 


---
# Links
- [container.training intro to docker](http://container.training/intro-selfpaced.yml.html#1)

- [What are container made from?](https://www.youtube.com/watch?v=sK5i-N34im8)




