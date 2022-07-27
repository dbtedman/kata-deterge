# [Deterge](https://github.com/dbtedman/kata-deterge) - Security Policy

[![SAST GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/kata-deterge/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/kata-deterge/actions/workflows/sast.yml)

Outlines how security is considered during the development of [Deterge](https://github.com/dbtedman/kata-deterge).

-   [Security Disclosure Policy](#security-disclosure-policy)
-   [Security Update Policy](#security-update-policy)
-   [Security Related Configuration](#security-related-configuration)
-   [Known Security Gaps and Future Enhancements](#known-security-gaps-and-future-enhancements)

## Security Disclosure Policy

Email [dbtedman@gmail.com](mailto:dbtedman@gmail.com) with details about the security issue.

## Security Update Policy

Best efforts will be taken to apply code fixes or update vulnerable packages as soon as is possible, this will usually be within a couple of days.

## Security Related Configuration

> 💡 Organised around categories from the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/)
> and [OWASP API Security Top 10 2019](https://owasp.org/www-project-api-security/) projects.

-   [Broken Access Control](#broken-access-control)
-   [Cryptographic Failures](#cryptographic-failures)
-   [Injection](#injection)
-   [Insecure Design](#insecure-design)
-   [Security Misconfiguration](#security-misconfiguration)
-   [Vulnerable and Outdated Components](#vulnerable-and-outdated-components)
-   [Identification and Authentication Failures](#identification-and-authentication-failures)
-   [Software and Data Integrity Failures](#software-and-data-integrity-failures)
-   [Security Logging and Monitoring Failures](#security-logging-and-monitoring-failures)
-   [Server-Side Request Forgery](#server-side-request-forgery)
-   [Broken Object Level Authorization](#broken-object-level-authorization)
-   [Excessive Data Exposure](#excessive-data-exposure)
-   [Lack of Resources and Rate Limiting](#lack-of-resources-and-rate-limiting)
-   [Broken Function Level Authorization](#broken-function-level-authorization)
-   [Mass Assignment](#mass-assignment)
-   [Improper Assets Management](#improper-assets-management)

### Broken Access Control

[A01:2021](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)

-   [GitHub Secret scanning](https://github.com/features/security) detects secrets incorrectly committed into the
    repository.

### Cryptographic Failures

[A02:2021](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)

_No relevant mitigations are in place._

### Injection

[A03:2021](https://owasp.org/Top10/A03_2021-Injection/), [API8:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa8-injection.md)

_No relevant mitigations are in place._

### Insecure Design

[A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)

_No relevant mitigations are in place._

### Security Misconfiguration

[A05:2021](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/), [API7:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa7-security-misconfiguration.md)

_No relevant mitigations are in place._

### Vulnerable and Outdated Components

[A06:2021](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)

-   [Snyk](https://snyk.io) or [GitHub Dependabot](https://github.com/features/security) scans NPM dependencies for know vulnerabilities and creates pull requests to
    resolve the vulnerabilities when available.

### Identification and Authentication Failures

[A07:2021](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/), [API2:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa2-broken-user-authentication.md)

_No relevant mitigations are in place._

### Software and Data Integrity Failures

[A08:2021](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)

-   Third party dependencies are carefully selected and described with explicit version numbers in `package.json`.
-   [Snyk](https://snyk.io) or [GitHub Dependabot](https://github.com/features/security) will raise an alert if a
    suspicious package, or version of a package is being depended upon.
-   Dependencies required during development are segregated from those required during operation of this plugin.

### Security Logging and Monitoring Failures

[A09:2021](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/), [API10:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xaa-insufficient-logging-monitoring.md)

_No relevant mitigations are in place._

### Server-Side Request Forgery

[A10:2021](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

_No relevant mitigations are in place._

### Broken Object Level Authorization

[API1:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa1-broken-object-level-authorization.md)

_No relevant mitigations are in place._

### Excessive Data Exposure

[API3:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa3-excessive-data-exposure.md)

_No relevant mitigations are in place._

### Lack of Resources and Rate Limiting

[API4:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa4-lack-of-resources-and-rate-limiting.md)

_No relevant mitigations are in place._

### Broken Function Level Authorization

[API5:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa5-broken-function-level-authorization.md)

_No relevant mitigations are in place._

### Mass Assignment

[API6:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa6-mass-assignment.md)

_No relevant mitigations are in place._

### Improper Assets Management

[API9:2019](https://github.com/OWASP/API-Security/blob/master/2019/en/src/0xa9-improper-assets-management.md)

_No relevant mitigations are in place._

## Known Security Gaps and Future Enhancements

_None currently._
