# `default` Example for {{ cookiecutter.provider }}/{{ cookiecutter.name }}

An example which shows the _default_ usage of the
`terraform-{{ cookiecutter.provider }}-{{ cookiecutter.name }}` module.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=1.4.0,<2.0.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_example"></a> [example](#module\_example) | ../../ | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_tag"></a> [tag](#input\_tag) | The tag to use for deployment of resources during testing | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tag"></a> [tag](#output\_tag) | The tag to use for deployment of resources during testing |
<!-- END_TF_DOCS -->
