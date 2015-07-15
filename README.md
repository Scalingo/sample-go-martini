Sample Application with Go and Martini
======================================

Running Locally
---------------

First, you need to have a working go environment:

http://golang.org/doc/install

### Godep

This go application is using github.com/kr/godeps to manage
its dependencies, it is required to install it in order to
get the right version of each third-party library.

`go get github.com/kr/godeps`

### Build
```sh
godep build
```

### Execute
```sh
./sample-go-martini
```

Deploying on Scalingo
---------------------

Create an application on https://scalingo.com, then:

```
git remote add scalingo git@scalingo.com:<name_of_your_app>.git
git push scalingo master
```

And that's it!

The application is running at this url: https://sample-go-martini.scalingo.io/

Deploy in one click
-------------------

[![Deploy to Scalingo](https://cdn.scalingo.com/deploy/button.png)](https://my.scalingo.com/deploy)
[![Deploy to Scalingo](https://cdn.scalingo.com/deploy/button-white.png)](https://d.scalingo.dev/deploy)

Links
-----

https://golang.org
https://github.com/go-martini/martini
