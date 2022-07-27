# [Deterge](https://github.com/dbtedman/kata-deterge)

> **⚠️ WARNING:** Not production ready code, instead a [Code Kata](https://github.com/dbtedman#code-kata) intended to
> hone my programming skills through practice and repetition.

[![ci workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-deterge/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-deterge/actions/workflows/ci.yml)
[![sast workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-deterge/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/kata-deterge/actions/workflows/sast.yml)

Tool for sanitising data files of sensitive information through substitution with fake information.

-   [Getting Started](#getting-started)
-   [Verification](#verification)
-   [Design](#design)
-   [Security](#security)
-   [References](#references)
-   [License](#license)

## Getting Started

### Prepare

Begin by [installing Go](https://go.dev/doc/install) if you have not done so already.

You can test your install by calling the following command:

```shell
go version
```

Your version must be greater than or equal to the version defined in `./go.mod` file.

### Install, Verify, and Build

Install, verify, and build `./deterge` binary.

```shell
nvm use && make
```

### Help

Learn about the available commands in the help menu.

```shell
./deterge --help
```

See [Commands](#commands) section for more information.

## Verification

### Linting

-   [Prettier](https://prettier.io)
-   [gofmt](https://pkg.go.dev/cmd/gofmt)

```shell
make lint
```

These rules can then be automatically applied:

```shell
make format
```

### Unit Testing

```shell
make test
```

## Design

### Usage Scenarios

#### MySQL Dump for Non-production Seeding of WordPress

> ⚠️ Not Yet Implemented

You have a mysql dump `sensitive.sql`, that contains sensitive data you wish to substitute with fake data before loading
into a non-production environment. The resulting `clean.sql` can then be loaded into a non-production environment.
The `wordpres` preset has an additional option to replace environment urls.

```shell
./deterge sql sensitive.sql clean.sql
```

### Repository Structure

_Placeholder_

### Commands

_Placeholder_

### Concepts

#### Internationalization

> 💡️ Perhaps using a package like [nicksnyder/go-i18n](https://github.com/nicksnyder/go-i18n). The currently selected
> language could be discovered from the environment, e.g. `LANG=en_AU.UTF-8` for command line
> and [Accept-Language](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language) header,
> e.g. `Accept-Language: en-AU,en;q=0.9` for the web ui.

## Security

See the [Security Policy](./SECURITY.md).

## References

-   [Cobra](https://cobra.dev) - A Commander for modern Go CLI interactions
-   [Viper](https://github.com/spf13/viper) - Go configuration with fangs

## License

The [MIT License](LICENSE.md) is used by this project.
