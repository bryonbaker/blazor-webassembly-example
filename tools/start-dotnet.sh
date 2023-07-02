#! /bin/bash

podman run --rm -it -v $(pwd):/app:Z -w /app -p 8080:8080 mcr.microsoft.com/dotnet/sdk:6.0