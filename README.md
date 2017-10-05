# mass-wpscan
[![Build Status](https://travis-ci.org/l50/mass-wpscan.svg?branch=master)](https://travis-ci.org/l50/mass-wpscan)
[![Go Report Card](https://goreportcard.com/badge/github.com/l50/mass-wpscan)](https://goreportcard.com/report/github.com/l50/mass-wpscan)
[![License](http://img.shields.io/:license-mit-blue.svg)](https://github.com/l50/mass-wpscan/blob/master/LICENSE)


Used to scan multiple wordpress sites with wpscan. Specify an input file
with targets, the parameters you want to use with wpscan for each
target, and a file to output the results to.

### Build the binary:
```
make install
```

### Build the test lab:
```
make buildtest
```

### Destroy the test lab:
```
make destroytest
```

** Note that all test lab functionality requires that python 2.x and
docker be installed on the system you are on

### Usage:
```
./mass-wpscan [options]
Options:
  -i string [required]
    Input file with targets.
  -o stringng
    File to output information to.
  -p string [required]
    Arguments to run with wpscan.
```

Your input file should be formatted like this:
```
http://0.0.0.0:44400
http://0.0.0.0:44401
http://0.0.0.0:44402
```

### TODO:
- [] Decide if we want to use the default wpscan parameters, or force
  the user to input them
- [] Make the scans concurrent
- [] Add comments

## License
MIT

### Resources
- https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
- https://gist.github.com/petermbenjamin/8aeece9305bb44282799384365ab3a3c#file-user-go
- https://github.com/averagesecurityguy/searchscan/blob/master/README.md
