# Testing

## Provider Tests
In order to test the provider, you can simply run `make test`.

```bash
$ make test
```

## Acceptance Tests

You can run the complete suite of sendgrid acceptance tests by doing the following:

```bash
$ make testacc TEST="./sendgrid/" 2>&1 | tee test.log
```

To run a single acceptance test in isolation replace the last line above with:

```bash
$ make testacc TEST="./sendgrid/" TESTARGS='-run=TestAccSendgridApiKey_BasicWithScopes'
```

A set of tests can be selected by passing `TESTARGS` a substring. For example, to run all sendgrid tests:

```bash
$ make testacc TEST="./sendgrid/" TESTARGS='-run=TestAccSendgridApiKey_BasicWithScopes'
```

### Test Parameters

The following parameters are available for running the test. The absence of some non-required parameters
will cause certain tests to be skipped.

* **TF_ACC** (`integer`) **Required** - must be set to `1`.
* **SENDGRID_API_KEY** (`string`) **Required**  - A valid Sendgrid API key.

**For example:**
```bash
export TF_ACC=1
export SENDGRID_API_KEY=...
$ make testacc TEST="./TestAccSendgridApiKey_BasicWithScopes/" 2>&1 | tee test.log
```
