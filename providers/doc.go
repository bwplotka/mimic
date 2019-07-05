/*
Providers are a way of defining a configuration specification in Go structs for projects that either do not currently
expose their configuration structs or are difficult to work with (due to deps, complexity etc.).

Each Provider should be as simple as possible and map as closely to the specification of the configuration that the
project documents.

Do not add abstractions, opinions or validation into the providers this should be done by an abstraction or the end user
that is wanting to use the provider. This avoids magic like the transforming of a property on a struct from one value to
another without the knowledge of the user.

Providers will be created by a caller wanting to generate specific config and the caller should `Add` this struct explicitly
to the Generator and not be done in the provider code.

As Providers should be basic structs you should be able to do away with the vast majority of conditionals or funcs within
these packages

Use the folder structure to help highlight what config the code will produce:
```
providers
  - vault       // Configuration passed to the vault product itself.
  - terraform
    - vault     // Terraform provider/resource configuration used by terraform.
```
 */
package providers
