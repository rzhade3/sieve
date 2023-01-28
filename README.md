# Sieve

Welcome to Sieve, a simple command line utility written in Golang that allows you to check if any files on a given domain are publicly exposed.

### Installation

To install Sieve, you need to have Golang installed in your system.

1. First, clone this repository using git clone `https://github.com/rzhade3/sieve.git`
2. Navigate to the project directory using `cd sieve`
3. Build the application using `go build .`
4. The binary file will be generated in the project directory, you can use this file to run the application.

### Usage

```bash
./sieve -domains <path to domain list> -files <path to files config list>
```

Sieve will check for the given files on the given domains and will tell you if any of the files are publicly exposed.

An example domain list and files config list are provided in the `config` directory.

**Note**

Make sure the files you are checking for exist on the specified domains and are publicly accessible before running the application.
