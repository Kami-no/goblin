# Goblin

Hello-world web server in Golang running at port 8080.

## Features

### Docker

Docker image might be created with Dockerfile.

### GitLab CI

Note:

* replace 'git.local' with your GitLab address;
* replace 'registry.local:4567' with your registry address (for GitLab it's usually GitLab address and port 4567);
* replace 'cluster.local' with wildcard domain for Kubernetes Ingress.

#### .gitlab-ci.yml

Contains tasks to lint yaml, build binary, docker-image, launch app for review in Kubernetes and retag image in the end as 'latest'.

Requirements:

* tasks are tagged, you'll need at least one runner with tags docker, k8s and linux;
* outdated namespaces should be deleted manually through task 'dev_wipe';
* project var should be created `CM_KUBE_CONFIG` with base64 encoded Kubernetes config with cluster administrator permissions (CI will create separate namespaces for each branch of the project):

```bash
cat ~/.kube/config | base64 | pbcopy
```

Artifacts:

* binary will expire in 1 day;
* k8s-resources.yml with full Kubernetes deployment manifest will expire in 1 day;
* docker-image tagged latest for builds in master and tagged with commit hash always.

### Bazel

```bash
bazel build //...
bazel build //:goblin
bazel run //:goblin

bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //...
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:goblin
```
