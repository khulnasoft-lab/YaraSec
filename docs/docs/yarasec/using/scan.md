---
title: Scan with YaraSec
---


# Scanning with YaraSec

You can use YaraSec to scan running or at-rest container images, and local file systems.  YaraSec will match the assets it finds against the YARA rules it has been configured with.

## Scan a Container Image

Pull the image to your local repository, then scan it

```bash
docker pull node:latest

docker run -it --rm --name=yara-hunter \
    -e KHULNASOFT_PRODUCT=<ThreatMapper or ThreatStryker> \
    -e KHULNASOFT_LICENSE=<ThreatMapper or ThreatStryker license key> \
    -v /var/run/docker.sock:/var/run/docker.sock \
    docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2 \
# highlight-next-line
    --image-name node:latest

docker rmi node:latest
```

### Scan a Running Container

Mount the root directory into the YaraSec container at a location of your choosing (e.g. `/khulnasoft/mnt`) and specify the running container ID:

```bash
docker run -it --rm --name=yara-hunter \
    -e KHULNASOFT_PRODUCT=<ThreatMapper or ThreatStryker> \
    -e KHULNASOFT_LICENSE=<ThreatMapper or ThreatStryker license key> \
    -v /var/run/docker.sock:/var/run/docker.sock \
# highlight-next-line
    -v /:/khulnasoft/mnt \
    docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2 \
# highlight-next-line
    --host-mount-path /khulnasoft/mnt --container-id 69221b948a73
```

### Scan a filesystem

Mount the filesystem within the YaraSec container and scan it:

```bash
docker run -it --rm --name=yara-hunter \
    -e KHULNASOFT_PRODUCT=<ThreatMapper or ThreatStryker> \
    -e KHULNASOFT_LICENSE=<ThreatMapper or ThreatStryker license key> \
# highlight-next-line
    -v ~/src/YARA-RULES:/tmp/YARA-RULES \
    docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2 \
# highlight-next-line
    --local /tmp/YARA-RULES --host-mount-path /tmp/YARA-RULES
```

### Scan during CI/CD build

Refer to the detailed [documentation for CI/CD integration](https://github.com/khulnasoft-lab/YaraSec/tree/main/ci-cd-integration).
