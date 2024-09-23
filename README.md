# inspo
[![GitHub release](https://img.shields.io/github/release/khulnasoft/inspo.svg)](https://github.com/khulnasoft/inspo/releases/latest)
[![Validations](https://github.com/khulnasoft/inspo/actions/workflows/validations.yaml/badge.svg)](https://github.com/khulnasoft/inspo/actions/workflows/validations.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft/inspo)](https://goreportcard.com/report/github.com/khulnasoft/inspo)
[![License: MIT](https://img.shields.io/badge/License-MIT%202.0-blue.svg)](https://github.com/khulnasoft/inspo/blob/main/LICENSE)

**A tool for exploring a docker image, layer contents, and discovering ways to shrink the size of your Docker/OCI image.**

To analyze a Docker image simply run develop with an image tag/id/digest:
```bash
develop <your-image-tag>
```

or you can develop with docker command directly
```
alias develop="docker run -ti --rm  -v /var/run/docker.sock:/var/run/docker.sock khulnasoft/develop"
develop <your-image-tag>

# for example
develop nginx:latest
```

or if you want to build your image then jump straight into analyzing it:
```bash
develop build -t <some-tag> .
```

Building on Macbook (supporting only the Docker container engine)

```bash
docker run --rm -it \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v  "$(pwd)":"$(pwd)" \
      -w "$(pwd)" \
      -v "$HOME/.develop.yaml":"$HOME/.develop.yaml" \
      khulnasoft/develop:latest build -t <some-tag> .
```

Additionally you can run this in your CI pipeline to ensure you're keeping wasted space to a minimum (this skips the UI):
```
CI=true develop <your-image>
```
