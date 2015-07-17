# ArchCI

## Introduction

ArchCI is the continues integration service integrated with open source projects.

It's new and adopt lightweight container to run task. Thanks to docker ecosystem, we can use docker for resource isolation and kuberntes for scheduling tasks. For more information, please refer to [slides](http://slides.com/tobychan/archci).

Worker: [simple-worker](https://github.com/ArchCI/simple-worker), [kubernetes-worker](https://github.com/ArchCI/kubernetes-worker)

Client: [aci](https://github.com/ArchCI/aci)

Document: [docs](https://github.com/ArchCI/docs)

Docker: [docker-distribution](https://github.com/ArchCI)

## Usage

```
docker run -d -p 80:80 archci/archci
```
