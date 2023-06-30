# {{ cookiecutter.name }} Terraform Module

This is the repository for the `{{ cookiecutter.name }}` Terraform module, which
the [`n3tuk`][n3tuk] [Terraform Cloud][tfc] [Private Module Registry][pmr] hosts
as [`{{ cookiecutter.name }}/{{ cookiecutter.provider }}`][module].

[n3tuk]: https://github.com/n3tuk
[tfc]: https://app.terraform.io/app
[pmr]: https://app.terraform.io/app/n3tuk/registry/private/modules
[module]: https://app.terraform.io/app/n3tuk/registry/modules/private/n3tuk/{{ cookiecutter.name }}/{{ cookiecutter.provider }}

> **Warning**
> This Terraform module is custom-written for the [`n3tuk`][n3tuk] Organisation,
> and so designed to be simple and highly opinionated. It's not created for
> general usage, nor will it ever be as such. Anyone is free to take it and use
> it as they see fit. As such, please note that this module may disappear or
> change at any time, and may not be suitable for any other purpose.

## Description

TODO: Update with the details of the purpose of this module.

## Usage

To use this module, call a `module {}` block with the `source` argument as
`app.terraform.io/n3tuk/{{ cookiecutter.name }}/{{ cookiecutter.provider }}`,
with the `version` you require:

```hcl
module "example" {
  source  = "app.terraform.io/n3tuk/{{ cookiecutter.name }}/{{ cookiecutter.provider }}"
  version = "= 0.1.0"
  # insert required variables here
}
```

Please ensure that you have logged in with [Terraform Cloud][tfc] using your
local `terraform` application as so to retrieve the credentials you need to
download Terraform modules from the [Private Module Registry][pmr]:

```sh
$ terraform login
Terraform will request an API token for app.terraform.io using your browser.

If login is successful, Terraform will store the token in plain text in
the following file for use by subsequent commands:
    /Users/jonathanwright/.terraform.d/credentials.tfrc.json

Do you want to proceed?
```

## Examples

For more detailed examples, see the [`examples/`][examples] directory in this
repository, which includes full example configurations and information about
each configuration too.

[examples]: https://github.com/n3tuk/{{ cookiecutter.repository.name }}/tree/master/examples

## Testing and Validation

This Terraform module uses a set of tools such as [Taskfile (or
`task`)][taskfile] and [`pre-commit`][pre-commit] to manage this repository
through development, testing and committing stages, as well as the utility
[`terraform-docs`][terraform-docs] to automatically manage the documentation,
and [`tfsec`][tfsec] for static security analysis of the configuration.

[taskfile]: https://taskfile.dev/
[pre-commit]: https://pre-commit.com/
[terraform-docs]: https://terraform-docs.io/
[tfsec]: https://github.com/aquasecurity/tfsec

For full details on how to work with this repository and these tools, see the
[`CONTRIBUTING.md`][contributing-md] document in this repository.

[contributing-md]: https://github.com/n3tuk/{{ cookiecutter.repository.name }}/blob/master/.github/CONTRIBUTING.md

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=1.4.0,<2.0.0 |

## Providers

No providers.

## Modules

No modules.

## Resources

No resources.

## Inputs

No inputs.

## Outputs

No outputs.
<!-- END_TF_DOCS -->
