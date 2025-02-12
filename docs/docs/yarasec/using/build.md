---
title: Build YaraSec
---

# Build YaraSec

YaraSec is a self-contained docker-based tool. Clone the [YaraSec repository](https://github.com/khulnasoft-lab/YaraSec), then build:

```bash
docker build --rm=true --tag=docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2 -f Dockerfile .
```

Alternatively, you can pull the official khulnasoft image at `docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2`.

```bash
docker pull docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2
```
