## Arch Linux

This repository contains `x86_64`, `armv7h` and `aarch64` packages.

### Usage

Add the [gesundheit build key](gesundheit.pem) to your pacman keyring.

```
# curl -LO https://ushis.github.io/gesundheit/gesundheit.pem
# pacman-key --add gesundheit.pem
# pacman-key --lsign 3531C50E3C86C7955A02399317B4A45C3B3AA155
```

Add the repository to your pacman config.

```
# /etc/pacman.conf

[gesundheit]
Server = https://ushis.github.io/$repo/arch/$arch
```

Install the package.

```
# pacman -Sy gesundheit
```

## Debian

This repository contains `i386`, `amd64`, `armhf` and `arm64` packages.

### Usage

Add the [gesundheit build key](gesundheit.pem) to your apt keyring.

```
# curl -L https://ushis.github.io/gesundheit/gesundheit.pem | \
  gpg --dearmor > /etc/apt/trusted.gpg.d/gesundheit.gpg
```


Add the repository to your apt source list.

```
# /etc/apt/sources.list

deb https://ushis.github.io/gesundheit/deb /
```

Install the package.

```
# apt-get update
# apt-get install gesundheit
```

## Repository Index

{% include_relative _repo.md %}
