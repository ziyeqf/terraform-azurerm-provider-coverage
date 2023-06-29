# azure-rest-api-bridge-coverage

A tool to statistics the coverage status of `azure-rest-api-bridge`(https://github.com/magodo/azure-rest-api-bridge) of `terraform-provider-azurerm`.

## Usage

```shell
 bridge-coverage -bridge ./bridge.json -schema ./schema.json
```

Note: bridge json file could be generated by running `azure-rest-api-bridge ... > bridge.json`
Note: schema json file could be generated by using [`schema-api`](https://github.com/hashicorp/terraform-provider-azurerm/tree/main/internal/tools/schema-api) `schema-api -export schema.json`


