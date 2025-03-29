<img src="https://raw.githubusercontent.com/cloudfoundry/logos/master/CF_Icon_4-colour.png" alt="CF logo" height="100" align="left"/>

# Cloud Foundry CLI

The official command line interface for [Cloud Foundry](https://cloudfoundry.org).

View the latest help documentation for [**The v8 CLI**](https://cli.cloudfoundry.org/en-US/v8) or [**The v7 CLI**](https://cli.cloudfoundry.org/en-US/v7), or run `cf help -a` to view all commands available in your currently installed version.

[![GitHub version](https://badge.fury.io/gh/cloudfoundry%2Fcli.svg)](https://github.com/cloudfoundry/cli/releases/latest)
[![Documentation](https://img.shields.io/badge/docs-online-ff69b4.svg)](https://docs.cloudfoundry.org/cf-cli)
[![Command help pages](https://img.shields.io/badge/command-help-lightgrey.svg)](https://cli.cloudfoundry.org)
[![Slack](https://slack.cloudfoundry.org/badge.svg)](https://slack.cloudfoundry.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/cloudfoundry/cli/blob/main/LICENSE)

CF CLI Binary Download Status:

[![Downloads Status](https://uptime.com/devices/services/widget/689896/c6d4bb7ddd16186d/service?light)](https://uptime.com/devices/services/689896/01026e1a663caab4)

---

## Quick Navigation

<p align="center">
<a href="#getting-started"><b>Getting Started</b></a> •
<a href="#downloads"><b>Download</b></a> •
<a href="#known-issues"><b>Known Issues</b></a> •
<a href="#filing-issues--feature-requests"><b>Bugs/Feature Requests</b></a> •
<a href="#plugin-development"><b>Plugin Development</b></a> •
<a href="#contributing--build-instructions"><b>Contributing</b></a>
</p>

---
## Getting Started

Download and install the cf CLI from the [Downloads Section](#downloads) for either the [v8 cf CLI](https://github.com/cloudfoundry/cli/wiki/V8-CLI-Installation-Guide) or the [v7 cf CLI](https://github.com/cloudfoundry/cli/wiki/V7-CLI-Installation-Guide).

Once installed, you can log in and push an app with a few simple commands:

```bash
# Log in to your Cloud Foundry instance
cf login -a <API_ENDPOINT>

# Push your application
cf push my-app
```

### Supported Versions

The cf CLI currently has two actively supported versions:

1. **v8 CLI** - Latest version backed by the [v3 CC API](http://v3-apidocs.cloudfoundry.org/). See [v8 documentation](https://docs.cloudfoundry.org/cf-cli/v8.html) for details.
2. **v7 CLI** - Previous version also backed by the [v3 CC API](http://v3-apidocs.cloudfoundry.org/). See [v7 documentation](https://docs.cloudfoundry.org/cf-cli/v7.html) for details.

For information about version support timelines, see our [CLI v7 & v8 Versioning and Support Policy](https://github.com/cloudfoundry/cli/wiki/Versioning-and-Support-Policy).

### Getting Help

If you have questions or need assistance:
- Join the #cli channel in [our Slack community](https://slack.cloudfoundry.org/)
- Post to the [cf-dev mailing list](https://lists.cloudfoundry.org/archives/list/cf-dev@lists.cloudfoundry.org/)
- [Open a GitHub issue](https://github.com/cloudfoundry/cli/issues/new)

## Contributing & Build Instructions

We welcome contributions from the community! Please read our [contributors' guide](.github/CONTRIBUTING.md) to get started.

### Translations

If you'd like to submit updated translations, please see the [i18n README](https://github.com/cloudfoundry/cli/blob/main/cf/i18n/README-i18n.md) for instructions on how to submit an update.

### CLI in Action

![Example](.github/cf_example.gif)

### Plugins

Extend the functionality of the CF CLI with [community contributed plugins](https://plugins.cloudfoundry.org).

## Downloads

### Installation Instructions

- **[Install V8 CLI](https://github.com/cloudfoundry/cli/wiki/V8-CLI-Installation-Guide)** (Recommended)
- **[Install V7 CLI](https://github.com/cloudfoundry/cli/wiki/V7-CLI-Installation-Guide)**
- **[Switching Between Multiple Versions](https://github.com/cloudfoundry/cli/wiki/Version-Switching-Guide)**

### Package Managers

```bash
# Homebrew (macOS)
brew install cloudfoundry/tap/cf-cli@8

# APT (Debian/Ubuntu)
sudo apt update
sudo apt install cf8-cli

# YUM (RHEL/CentOS)
sudo yum install cf8-cli

# Chocolatey (Windows)
choco install cloudfoundry-cli
```

For more installation options, see the [installation guides](https://github.com/cloudfoundry/cli/wiki/V8-CLI-Installation-Guide).

## Known Issues

> **Note:** For the most up-to-date information on issues and workarounds, please review [the open and closed GitHub issues](https://github.com/cloudfoundry/cli/issues).

### Windows Issues

* **Password Prompts**: On Windows in Cygwin and Git Bash, interactive password prompts (in `cf login`) do not hide the password properly ([issue #1835](https://github.com/cloudfoundry/cli/issues/1835)). Use `cf auth` instead or switch to Windows `cmd`.

* **SSH Display**: `cf ssh` may not display correctly if the `TERM` environment variable is not set. Setting `TERM=msys` often resolves these issues.

* **SSH Hanging**: `cf ssh` will hang when run from MINGW32 or MINGW64 shells. Use PowerShell instead.

### Certificate Issues

* **Self-Signed Certificates**: CF CLI doesn't use OpenSSL. Custom/Self-Signed Certificates must be [installed in specific locations](https://docs.cloudfoundry.org/cf-cli/self-signed.html) to use `login`/`auth` without `--skip-ssl-validation`.

### Other Issues

* **API Tracing**: Tracing to terminal (`CF_TRACE=true`, `-v` option, or `cf config --trace`) doesn't work well with some plugins. Use file tracing instead. On Linux, `CF_TRACE=/dev/stdout` works as an alternative.

* **.cfignore Encoding**: Files must use UTF-8 encoding for the CLI to interpret correctly ([issue #281](https://github.com/cloudfoundry/cli/issues/281#issuecomment-65315518)).

* **Linux Architecture**: If you see "bash: .cf: No such file or directory", ensure you're using the [correct binary for your architecture](https://askubuntu.com/questions/133389/no-such-file-or-directory-but-the-file-exists).

* **Warning Output**: X-Cf-Warnings are printed to `stdout`. Set `CF_RAISE_ERROR_ON_WARNINGS` to redirect warnings to `stderr` ([issue #2164](https://github.com/cloudfoundry/cli/issues/2164)).

* **Org Creation**: Fixed in v7.2 - Earlier versions had false negative messages when non-admin users with the user-org-creation flag enabled ran `cf create-org` ([issue #1879](https://github.com/cloudfoundry/cloud_controller_ng/issues/1879)).

## Filing Issues & Feature Requests

Before filing an issue:

1. Update to the [latest CLI version](https://github.com/cloudfoundry/cli/releases) and try the command again
2. Check if the issue is already reported in the [open issues](https://github.com/cloudfoundry/cli/issues)

If the problem persists and no existing issue covers it, [file a new issue](https://github.com/cloudfoundry/cli/issues/new/choose) with:

- CLI version (`cf version`)
- Cloud Foundry API endpoint version (`cf api`)
- Steps to reproduce
- Expected vs. actual behavior
- Any relevant logs or error messages

## Plugin Development

The CF CLI supports extending functionality through plugins. Build your own commands and workflows!

### Resources

* [Plugin Development Guide](https://github.com/cloudfoundry/cli/tree/main/plugin/plugin_examples)
* [Official Plugins Repository](https://plugins.cloudfoundry.org/)
* [Example Plugins](https://github.com/cloudfoundry/cli/tree/main/plugin/plugin_examples)

### Import Path

When developing plugins, use:
```go
import "code.cloudfoundry.org/cli/plugin"
```

Legacy plugins using `github.com/cloudfoundry/cli/plugin` will still work if they vendor the plugins directory.

### Installing Plugins

```bash
# Install from repository
cf install-plugin -r CF-Community "plugin-name"

# Install from local file
cf install-plugin path/to/plugin.exe

# List installed plugins
cf plugins
```
