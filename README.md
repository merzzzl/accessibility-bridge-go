# Go Client for Accessibility-Bridge

A tiny **Makefile** that

1. clones the upstream project **merzzzl/accessibility-bridge** (if it isn’t present yet);
2. optionally pulls the latest changes;
3. generates Go code from every `.proto` file in `/app/src/main/proto`;
4. writes the result to root.

## 🛠  Prerequisites

| Tool / Version              | Check with                         |
| --------------------------- | ---------------------------------- |
| Go 1.22 or newer            | `go version`                       |
| `protoc` ≥ 3.21             | `protoc --version`                 |
| Protobuf plugins in `$PATH` | `protoc-gen-go`, `protoc-gen-go-grpc` |

Install the plugins once:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
````

## 🚀  Quick Start

```bash
# clone this repository (the one containing the Makefile)
git clone github.com/merzzzl/accessibility-bridge-go accessibility-bridge-go
cd bridge-proto

# generate Go code (the script will clone the upstream repo automatically)
make
```

Generated files will appear in `gen/`.

## 📚  Common Make Targets

| Target (`make …`)   | Description                                        |
| ------------------- | -------------------------------------------------- |
| *(default)* / `all` | Clone (if needed) **and** generate Go code         |
| `update`            | `git pull` in the upstream repo, then regenerate   |
| `proto`             | Generate code only (skip `git pull`)               |
| `clean`             | Remove the upstream clone and the `gen/` directory |

### Changing the Go import path

By default all generated files belong to

```
github.com/merzzzl/accessibility-bridge
```

## 🗂  Directory Layout

```
.
├── Makefile               # main build script
├── *.go                   # generated Go sources
└── accessibility-bridge/  # auto-cloned upstream project
```

## 📄  License

MIT
