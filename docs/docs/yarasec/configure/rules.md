---
title: Configure Rules
---


# Provide Custom Rules

YaraSec uses the YARA rules files (`*.yar` and `*.yara`) in the `/home/khulnasoft/rules` directory in the container.  You can provide an alternative set of rules, either at build-time, or by mounting a new rules directory into the container.

You can mount the rules directory over the existing one (using `-v $(pwd)/my-rules:/home/khulnasoft/rules`). Alternatively, you can mount the rules directory in a different location and specify it with `--rules-path`:

```bash
# Put your rules files (*.yar, *.yara) in the ./my-rules directory

mkdir ./my-rules

docker run -it --rm --name=yara-hunter \
    -e KHULNASOFT_PRODUCT=<ThreatMapper or ThreatStryker> \
    -e KHULNASOFT_LICENSE=<ThreatMapper or ThreatStryker license key> \
    -v /var/run/docker.sock:/var/run/docker.sock \
# highlight-next-line
    -v $(pwd)/my-rules:/tmp/my-rules \
    docker.io/khulnasoft/khulnasoft_malware_scanner_ce:2.5.2 --image-name node:latest \
# highlight-next-line
    --rules-path /tmp/my-rules
```


