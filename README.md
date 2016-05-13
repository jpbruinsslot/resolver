Resolver
========

Overview
--------

This project served as a service for assets resolving for static files that are
being served by a CDN and whose filenames are hashed.

For instance when you reference static files from a template in your favourite
framework, you'll probably have a tag like this:

```html
<link type="text/css" media="screen" href="assets.css" rel="stylesheet">
```

Now, when those files are served by a CDN, and you hashed the filenames, you
want to reference those correct, up-to-date files. So the filename `assets.css`
need to be resolved to the correct hashed filename that is served by the CDN.

One way to do this is by using *Resolver*. It creates a small service which
you can query to resolve the filenames ot its hashed equivalent available on
the CDN. It does so by storing the key, value combinations of filenames, and 
hashes in a json file.  It exposes several endpoints which can be used to
query or update this file.

Hopefully you'll be able to create an equivalent of a `template_tag` in your
chosen framework that enables you to references a function that can query
the resolver service. From our example above, this will look something like
this, if you were using Django.

```html
<link type="text/css" media="screen" href="{% resolve_assets assets.css %}" rel="stylesheet">
```

Endpoints
---------

Resolver exposes the following endpoint:

### GET /assets/

Response:
```javascript
// STATUS 200
{
    "asset1": "hash1",
    "asset2": "hash2"
}
```

### POST /assets/

Input:
```javascript
{
    "asset1": "new_hash1",
    "asset2": "new_hash2"
}
```

Response:
```javascript
// STATUS 201
{
    "asset1": "new_hash1",
    "asset2": "new_hash2"
}
```

Running Resolver
----------------

The easiest of running this service is using docker. A binary is also provided
in `bin/` if you want to run it that way.

```bash
$ docker build -t resolver .
$ docker run \
    -p 4000:4000 \
    -e "RESOLVER_USER=[user-name]" \
    -e "RESOLVER_KEY=[key]" \
    -v /dir-on-host:/data \
    -it resolver

$ curl -U [user-name]:[password] http://localhost:4000/assets/
```

Resolver uses http basic authentication, you can create a key at the following
site: [Htpasswd Generator](http://www.htaccesstools.com/htpasswd-generator/)
