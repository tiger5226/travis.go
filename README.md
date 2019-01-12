# travis.go
Travis webhook library to make leveraging the payload information EASY with go

# Usage #

It makes handling the web hook really easy!

```golang
isPrivateRepo := true
err := travis.ValidateSignature(isPrivateRepo, r)
if err != nil {
    //Handle Invalid Sig
}

webhook, err := travis.NewFromRequest(r)
if err != nil {
    //Handle Invalid webhook parsing
}
if webhook.ShouldDeploy() {
    //Deploy the application
}

```

There are a few built in hooks. If you want your own tests for Deployment you can set a new function to be used during `ShouldDeploy`.


#### ShouldDeploy ####

This replaces the current logic. The default checks that it is on the master branch, it was successful, and that it is not a pull request.

```golang
travis.ShouldDeployFn = func (w travis.Webhook) bool {
    w.Branch = "develop" &&
    w.IsPullRequest = false
}
```

#### IsMatch ####

This is in addition to the current logic. The default only checks the branch, repo, and owner for match.

```golang
travis.IsMatchFn =  func (w travis.Webhook) bool {
    w.
}
```

The match can be checked like so:

```golang

webhook, err := travis.NewFromRequest(r)
if err != nil {
    //Handle Invalid webhook parsing
}
if !webhook.IsMatch(){
    //Handle lack of match
}

```