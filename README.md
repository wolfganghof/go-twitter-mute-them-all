# Twitterの全フォロユーザーをミュートしてタイムラインを静かにする

このツールでは、Twitterのフォローしているユーザを全員ミュートします。

# How to use

Fill in those keys in cred.go:

```go
// ConsumerKey is...
const ConsumerKey = ""

// ConsumerSecret is...
const ConsumerSecret = ""

// OAuthToken is...
const OAuthToken = ""

// OAuthSecret is...
const OAuthSecret = ""
```

Then,

```bash
make run
```
