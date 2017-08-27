# censys-definitions
Censys internal Protobuf objects

## Regenerating
Only do this on OS X in a clean Python 2.7 virtualenv.

The build script requires that [sh](https://amoffat.github.io/sh/) is installed
in your virtualenv.

```console
$ # In a virtualenv
$ pip install sh
```

If you already have grpc or protobuf installed, remove it. Note that there's
several formula's for grpc and protobuf that may be installed in Homebrew.

```console
$ brew uninstall grpc/grpc/grpc
$ brew uninstall grpc
$ brew uninstall google-protobuf
$ brew uninstall protobuf
```

Install the versions from the [ZMap Homebrew
Tap](https://github.com/zmap/homebrew-formula).

```console
$ brew tap zmap/homebrew-formula
$ brew install zmap/formula/grpc@1.2
$ ./build.py
```

Install Protobuf for Golang
```console
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

Recompile all .proto files. This _does not_ regenerate any autogenerated .proto files:
```console
$ ./build.py
```

## Regenerating ZLint protobuf

[proto/zlint.proto](proto/zlint.proto) is partially machine generated. To regenerate `message Lints` when new lints are added to ZLint:

0. Ensure the protobuf Python package is installed in your virtualenv (`pip
   install protobuf`).
1. Save `--list-lints-json` output from ZLint to `zlint-lints.json`. See ZLint
   documentation for details.
2. Run `build.py --generate-zlint-state`, which will read `proto/zlint.proto` and generate `zlint-state.json`. *Warning*: This will overwrite any existing file.
3. Given a `zlint-state.json` and a `zlint-lints.json`, we can regenerate the proto file. Specify these files using `--zlint-state` and `--zlint-lints` arguments to `build.py`. The default values match the names used earlier. Run `build.py --generate-zlint-proto`.
4. Recompile the protobufs.

Putting it all together:

```console
$ zlint --list-lints-json > zlint-lints.json
$ ./build.py --generate-zlint-state
$ ./build.py --generate-zlint-proto --zlint-lints zlint-lints.json --zlint-state zlint-state.json
$ ./build.py all
```

## Uploading zsearch-definitions to PyPi

Bump the version number in `python/setup.py`. Be sure to commit & push this change.

Then, in the `python` directory of this repo:
```
python setup.py sdist upload
```

Verify the update appears [on PyPi](https://pypi.python.org/pypi/zsearch-definitions).
