# Beat entry level IT exam of Lithuania

Have you ever wondered if you could still pass `Information Technology` National-level Maturity Exam of Lithuania?
How would you do it in reality? Are requirements could still be relevant for today's IT world practitioners?

Let's have fun and celebrate Programmer's day with this challenge.

## How

 * Read the [requirements](requirements/requirements.md)
 * Write code in you favorite language (see [`go`](go/main.go), [`js`](js/main.js), [`php`](php/main.php) folders)
 * Create a Pull request and see automated tests check your results 

## Running tests locally

Assuming you have [`docker`](https://docs.docker.com/engine/install/) installed.

### PHP 8.2

```shell
docker run -v ${PWD}/php:/opt/php aurelijusbanelis/callenge:php 
```
Where your code is in `php` folder with `php/main.php` file.

### JavaScript (Node 20)

```shell
docker run -v ${PWD}/js:/opt/js aurelijusbanelis/callenge:js 
```
Where your code is in `js` folder with `js/main.js` file.

### Go 1.21

```shell
docker run -v ${PWD}/go:/opt/go aurelijusbanelis/callenge:go 
```
Where your code is in `go` folder with `go/main.go` file.

### Not trusting docker?

Build it yourself:
```shell
docker build --target go . -t aurelijusbanelis/callenge:go
docker build --target php . -t aurelijusbanelis/callenge:php
docker build --target js . -t aurelijusbanelis/callenge:js
```

See [Dockerfile](Dockerfile) and [GitHub Actions](.github/workflows/infrastructure.yml) for details.

## References

* https://www.nsa.smm.lt/wp-content/uploads/2023/06/2023_IT_VBE_pg-web.pdf - original exam example
