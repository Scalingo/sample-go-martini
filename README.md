Sample Application with Go and Martini
======================================

Running Locally
---------------

First, you need to have a working go environment:

http://golang.org/doc/install

```sh
go get github.com/Appsdeck/sample-go-martini
go install github.com/Appsdeck/sample-go-martini
sample-go-martini
```

Deploying on Appsdeck
---------------------

Create an application on https://appsdeck.eu, then:

```
git remote add appsdeck git@appsdeck.eu:<name_of_your_app>.git
git push appsdeck master
```

And that's it!

The application is running at this url: https://sample-go-martini.appsdeck.eu/

Links
-----

https://golang.org
https://github.com/go-martini/martini
