---
title: Configure Output
---


# Configure Output

YaraSec can writes output to `stdout`. It can be redirected to a file for further analysis.

```bash
docker run -i --rm --name=yara-hunter \
    -e KHULNASOFT_PRODUCT=<ThreatMapper or ThreatStryker> \
    -e KHULNASOFT_LICENSE=<ThreatMapper or ThreatStryker license key> \
    -v /var/run/docker.sock:/var/run/docker.sock \
    docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2 \
    --image-name node:latest \
# highlight-next-line
    --output=json > xmrig-scan.json
```
