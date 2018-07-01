/*
Package resourcetypes provides ability to manage resource types through the Gnocchi API.

Example of Listing resource types

  allPages, err := resourcetypes.List(client).AllPages()
  if err != nil {
    panic(err)
	}

  allResourceTypes, err := resourcetypes.ExtractResourceTypes(allPages)
  if err != nil {
    panic(err)
	}

  for _, resourceType := range allResourceTypes {
    fmt.Printf("%+v\n", resourceType)
  }

Example of Getting a resource type

  resourceTypeName := "compute_instance"
  resourceType, err := resourcetypes.Get(gnocchiClient, resourceTypeName).Extract()
  if err != nil {
    panic(err)
  }

Example of Creating a resource type

  resourceTypeOpts := resourcetypes.CreateOpts{
    Name: "compute_instance_network",
    Attributes: map[string]resourcetypes.AttributeOpts{
      "port_name": resourcetypes.AttributeOpts{
        Type: "string",
        Details: map[string]interface{}{
          "max_length": 128,
          "required":   false,
        },
      },
      "port_id": resourcetypes.AttributeOpts{
        Type: "uuid",
        Details: map[string]interface{}{
          "required": true,
        },
      },
    },
  }
  resourceType, err := resourcetypes.Create(gnocchiClient, resourceTypeOpts).Extract()
  if err != nil {
    panic(err)
  }
*/
package resourcetypes
