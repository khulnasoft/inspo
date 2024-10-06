# inspo
[![GitHub release](https://img.shields.io/github/release/khulnasoft/inspo.svg)](https://github.com/khulnasoft/inspo/releases/latest)
[![Validations](https://github.com/khulnasoft/inspo/actions/workflows/validations.yaml/badge.svg)](https://github.com/khulnasoft/inspo/actions/workflows/validations.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft/inspo)](https://goreportcard.com/report/github.com/khulnasoft/inspo)
[![License: MIT](https://img.shields.io/badge/License-MIT%202.0-blue.svg)](https://github.com/khulnasoft/inspo/blob/main/LICENSE)

**A tool for exploring a docker image, layer contents, and discovering ways to shrink the size of your Docker/OCI image.**

To analyze a Docker image simply run inspo with an image tag/id/digest:
```bash
inspo <your-image-tag>
```

or you can inspo with docker command directly
```
alias inspo="docker run -ti --rm  -v /var/run/docker.sock:/var/run/docker.sock khulnasoft/inspo"
inspo <your-image-tag>

# for example
inspo nginx:latest
```

or if you want to build your image then jump straight into analyzing it:
```bash
inspo build -t <some-tag> .
```

Building on Macbook (supporting only the Docker container engine)

```bash
docker run --rm -it \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v  "$(pwd)":"$(pwd)" \
      -w "$(pwd)" \
      -v "$HOME/.inspo.yaml":"$HOME/.inspo.yaml" \
      khulnasoft/inspo:latest build -t <some-tag> .
```

Additionally you can run this in your CI pipeline to ensure you're keeping wasted space to a minimum (this skips the UI):
```
CI=true inspo <your-image>
```
