Sample Application with Go and Martini
======================================

Running Locally
---------------

First, you need to have a working go environment:

http://golang.org/doc/install

### Build
```sh
go build
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

[![Deploy to Scalingo](https://cdn.scalingo.com/deploy/button.svg)](https://my.scalingo.com/deploy)

Links
-----

https://golang.org
https://github.com/go-martini/martini
