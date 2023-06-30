# terraform-azurerm-provider-coverage

A tool to statistics the coverage status of [`azure-rest-api-bridge`](https://github.com/magodo/azure-rest-api-bridge) of `terraform-provider-azurerm`.

## Usage

```shell
 terraform-azurerm-provider-coverage -input ./coverage.json -schema ./schema.json -ignore-schema name,resource_group_name -map-identity KEY
```

- Note: schema json file could be generated by using [`schema-api`](https://github.com/hashicorp/terraform-provider-azurerm/tree/main/internal/tools/schema-api) `schema-api -export schema.json`

## Input Sample
```json
{
  "azurerm_resource_group": {
      "/location": {},
      "/name": {},
      "/tags/KEY": {}
  }
}
```

## Parameters

- `input`: the Coverage JSON file, in the format of `Input Sample`.
- `schema`: the Schema JSON file.
- `ignore-schema`: the schema path to ignore, separated by `,`.
- `map-identity`: the key of elements in TypeMap, defaults to `0`.
- `ignore-empty-resources`: Whether to ignore schema of uncovered and empty resources, defaults to `false`.