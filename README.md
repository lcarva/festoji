# festoji

Seasonal festive emoji for the day 🎉

## Usage

Simply run the golang program:

```text
go run main.go
```

Or, compile it into a binary and run it:

```text
go build -o bin/festoji main.go
./bin/festoji
```

The current festive emoji will be printed to stdout.

You can also use it as a container:

```text
podman run --rm quay.io/lucarval/festoji:latest
```

### Verifying Image Signature

The festoji container image is signed and attested. [cosign](https://github.com/sigstore/cosign)
version 2 is required.

To verify the image signature:

```text
cosign verify quay.io/lucarval/festoji:latest \
  --certificate-github-workflow-repository lcarva/festoji \
  --certificate-identity 'https://github.com/lcarva/festoji/.github/workflows/package.yaml@refs/heads/master' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com'
```

To verify the image SBOM attestation:

```text
cosign verify-attestation quay.io/lucarval/festoji:latest \
  --type spdx \
  --certificate-github-workflow-repository lcarva/festoji \
  --certificate-identity 'https://github.com/lcarva/festoji/.github/workflows/package.yaml@refs/heads/master' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com'
```

To verify the image SLSA Provenance attestation:

```text
cosign verify-attestation quay.io/lucarval/festoji:latest \
  --type slsaprovenance \
  --certificate-github-workflow-repository lcarva/festoji \
  --certificate-identity 'https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@refs/tags/v1.7.0' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com'
```

## Customizing

Festoji comes with preset default rules. These can be overwritten by creating the file
`~/.festoji.yaml`. For example:

```yaml
---
# The default character to be used when no rules match.
default: 🐚

# If set to true, this configuration extends the default one. Rules are inserted after the
# default ones. If set to false, the default configuration is completely ignored.
extend: true

rules:
- name: Xmas
  emoji: 🎄
  # This rule will start matching 14 days prior to December 25th
  span: 14
  month: 12
  day: 25
- name: Thanksgiving
  emoji: 🦃
  # This rule will start matching 7 days prior to the fourth Thursday in November
  span: 7
  month: 11
  week: 4
  weekday: 4
```
