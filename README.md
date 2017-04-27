# Drone - Sentry Integration
Generates a small dockerized tool that notifies Sentry whenever a new Release / Deploy happens.

Not that in the context of Sentry there is a difference between a `deploy` and a `release`. A
`release` spans multiple services and is meant to imply a large roll out. A deploy is for a single
codebase and implies the change to a single system. A release could happen over the course of a
long time, possibly even multiple days, whereas a deploy is pretty much a point in time. In effect,
a release can be made up of many different deploys. All deploys happen in the context of a release.

To keep this in line with the Sentry way of doing things, we perform two steps:

1. Ensure that a Release exists with the same version identifier. For example, if we had deployed
   something to staging before but are now deploying it to production, then the Release would
   already exist, but no production deploys would exist yet. If it's the first time deploying it
   anywhere, then we need to create that overall release.
2. Create a deploy for the same version under the existing release and tag it with the environment
   it is being pushed out to. E.g. the first time it might be to staging and the second time to
   production.

When running the program or the docker instance, the following environment variables need to be set:

* sentry_auth_token = a write-allowed API key created in https://sentry.io/api
* sentry_organization = the name of your Sentry organization (shown in the URL when you're logged in)
* sentry_project = the name of the project that is being released / deployed (e.g. "api")
* sentry_release_version = the version # or name of what is being released / deployed
* sentry_release_environment = the name of the environment (e.g. "stage" or "prod")

For example:

```yaml
pipeline:
  sentry-staging:
    image: urbint/drone-sentry
    sentry_auth_token: 1234567890abcdefg1234567890abcde1234567890abcdefg1234567890abcde
    sentry_organization: my-org-short-name
    sentry_project: my-project-name
    sentry_release_version: 1.2.3
    sentry_release_environment: staging
    when:
      branch: master
```