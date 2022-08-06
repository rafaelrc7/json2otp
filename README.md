# json2otp

I made this simple script so that I could convert andOTP backup json files into
otpauth urls that I could use with pass.

## Usage

json2otp reads from stdin and outputs to stdout, thus, you can easily pipe the
json into it or from a file

eg.:
```sh
json2otp < otp.json > otps

gpg --decrypt < otps.json.gpg | json2otp > otps
```

## Building
### Dependencies
- Go

### Building
clone repo and cd into int and just run

```sh
go build
```

## Licence
Licenced under the [MIT Licence](/LICENCE).

