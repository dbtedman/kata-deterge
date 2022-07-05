# [Deterge](https://github.com/dbtedman/kata-deterge)

> **⚠️ WARNING:** Not production ready code, instead a [Code Kata](https://github.com/dbtedman#code-kata) intended to
> hone my programming skills through practice and repetition.

[![ci workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-deterge/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-deterge/actions/workflows/ci.yml)
[![sast workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-deterge/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/kata-deterge/actions/workflows/sast.yml)
![language: go](https://img.shields.io/badge/language-go-blue.svg?style=for-the-badge)

Tool for sanitising data files of sensitive information through substitution with fake information.

- [Getting Started](#getting-started)
- [Verification](#verification)
- [Design](#design)
- [References](#references)
- [License](#license)

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

- [Prettier](https://prettier.io)
- [gofmt](https://pkg.go.dev/cmd/gofmt)

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

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

#### [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)

[Github Security](https://github.com/features/security) detects secrets incorrectly committed into the repository.

#### [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)

_Placeholder_

#### [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)

_Placeholder_

#### [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)

_Placeholder_

#### [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)

_Placeholder_

#### [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)

[Snyk](https://snyk.io) and [Github Security](https://github.com/features/security) scan Gradle and NPM dependencies for
know vulnerabilities and create pull requests to resolve the vulnerabilities when available.

#### [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/)

_Placeholder_

#### [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)

_Placeholder_

#### [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)

_Placeholder_

#### [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

_Placeholder_

## References

- [Cobra](https://cobra.dev) - A Commander for modern Go CLI interactions
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs

## License

The [MIT License](LICENSE.md) is used by this project.
