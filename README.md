# ArchCI

## Introduction

ArchCI is the continues integration service integrated with open source projects.

It's new and adopt lightweight container to run task. Thanks to docker ecosystem, we can use docker for resource isolation and kuberntes for scheduling tasks.

Worker: [simple-worker](https://github.com/ArchCI/simple-worker), [kubernetes-worker](https://github.com/ArchCI/kubernetes-worker)
Client: [aci](https://github.com/ArchCI/aci)
Document: [docs](https://github.com/ArchCI/docs)
Docker: [docker-distribution](https://github.com/ArchCI)

## Usage

Currently you need to setup docker, postgresql and redis before running archci.

```
cd ArchCI/archci
go get
go build
```

It's written in go so you can run it at most platforms. We're providing an all-in-one docker image, which might minimize the effot to setup the continues integration service.

## Development

Front-end

```
cd static
npm run watch
# Release
npm run lint
npm run build
```

