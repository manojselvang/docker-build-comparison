# 🚀 Reducing Docker Image Size Using Multi-Stage Builds

## Overview

Container image size directly impacts deployment speed, storage consumption, network transfer time, and the overall attack surface.

This project demonstrates how Docker **Multi-Stage Builds** can significantly reduce the size of a production container image by separating the build environment from the runtime environment.

Using the same Go REST API, this repository compares:

- Single-Stage Docker Build
- Multi-Stage Docker Build

and measures the improvements achieved by adopting multi-stage builds.

---

# Problem Statement

A traditional Docker build packages everything into a single image, including:

- Go compiler
- Build cache
- Source code
- Dependencies
- Runtime binary

Although the application works, the final image is unnecessarily large and contains components that are never required in production.

This results in:

- Larger container images
- Slower image pulls
- Longer deployment times
- Increased storage usage
- Larger attack surface

---

# Solution

A Multi-Stage Docker build separates the application build process from the runtime environment.

Instead of shipping the complete development environment, only the compiled Go binary is copied into the final image.

This produces a lightweight production-ready container.

---

# Project Structure

```text
docker-build-comparison/
│
├── app/
│   ├── go.mod
│   ├── main.go
│   └── handlers.go
│
├── single-stage/
│   └── Dockerfile
│
└── multi-stage/
    └── Dockerfile
```

---

# Project Workflow

```
               Go REST API
                    │
        ┌───────────┴───────────┐
        │                       │
 Single-Stage Build      Multi-Stage Build
        │                       │
 Large Runtime Image     Minimal Runtime Image
        │                       │
     Comparison of Image Size, Layers,
 Deployment Speed and Security
```

---

# REST API

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | / | Welcome Message |
| GET | /health | Health Check |
| GET | /version | Application Version |
| GET | /time | Current Server Time |

---

# Outcome

After implementing Multi-Stage Builds, the production image:

- Removes the Go compiler
- Removes source code
- Removes build dependencies
- Produces a significantly smaller image
- Reduces deployment time
- Reduces storage requirements
- Improves container security

---

# Build Commands

## Single Stage

```bash
docker build -f single-stage/Dockerfile -t go-single-stage:v1 .
```

## Multi Stage

```bash
docker build -f multi-stage/Dockerfile -t go-multi-stage:v1 .
```

---

# Validation

Compare image sizes:

```bash
docker images
```

Inspect image layers:

```bash
docker history go-single-stage:v1

docker history go-multi-stage:v1
```

---

# Expected Results

| Metric | Single Stage | Multi Stage |
|----------|-------------|-------------|
| Image Size | Higher | Lower |
| Build Tools Included | Yes | No |
| Source Code Included | Yes | No |
| Runtime Image | Full SDK | Minimal Runtime |
| Deployment Speed | Slower | Faster |
| Attack Surface | Larger | Smaller |

---

# Key Takeaways

- Multi-stage builds create production-ready Docker images.
- Only the application binary is shipped to production.
- Smaller images improve portability and deployment efficiency.
- Removing unnecessary components reduces the attack surface.
- Multi-stage builds are a Docker best practice for production workloads.

---

# Author

**Manoj Selvan G**
linkedin.com/in/manojselvang/
manojselvang@gmail.com
