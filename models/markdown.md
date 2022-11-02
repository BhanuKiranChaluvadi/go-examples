


# urcap API
API description for URCaps.
  

## Informations

### Version

1.0.0

### License

[Apache '2.0'](http://www.apache.org/licenses/LICENSE)

### Contact

Universal Robots swpl@universal-robots.com https://www.universal-robots.com/

## Tags

  ### <span id="tag-urcap"></span>urcap

Operations available on urcap

  ### <span id="tag-artifact"></span>artifact

Operations available on artifact

  ### <span id="tag-container"></span>container

Operations available on container

  ### <span id="tag-web-archive"></span>webArchive

Operations available on web archive

  ### <span id="tag-polyscope-bundle"></span>polyscopeBundle

Operations available on polyscope bundle

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json
  * multipart/form-data

### Produces
  * application/json

## All endpoints

###  artifact

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/artifacts/{artifactType}/name | [get artifact actual name by path](#get-artifact-actual-name-by-path) | Gets the actual name of an artifact from the resource path |
| GET | /api/v1/artifacts/{artifactType}/info | [get artifact metadata by path](#get-artifact-metadata-by-path) | Gets the urcap metadata from the resource path |
  


###  container

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/containers/{vendorID}/{urcapID}/{artifactName} | [get container by Id](#get-container-by-id) | Gets a container info by urcap ID's and artifactName |
| GET | /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/portmapping | [get container port mapping by Id](#get-container-port-mapping-by-id) | Port mapping of a container |
| PUT | /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/start | [start container by Id](#start-container-by-id) | Start a container |
| PUT | /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/status | [status container by Id](#status-container-by-id) | Status of a container |
| PUT | /api/v1/containers/stop | [stop all containers](#stop-all-containers) | Stops all containers |
| PUT | /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/stop | [stop container by Id](#stop-container-by-id) | Stop a container |
| GET | /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/validate | [validate container by Id](#validate-container-by-id) | Validates container artifact presence |
  


###  urcap

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/urcaps | [add urcap](#add-urcap) | Adds a new urcap |
| DELETE | /api/v1/urcaps/{vendorID}/{urcapID} | [delete urcap by Id](#delete-urcap-by-id) | Deletes a urcaps info by ID's |
| GET | /api/v1/urcaps/{vendorID}/{urcapID} | [get urcap by Id](#get-urcap-by-id) | Gets a urcap info by urcap ID's |
| GET | /api/v1/urcaps | [get urcaps](#get-urcaps) | Gets all urcaps. |
| PUT | /api/v1/urcaps/{vendorID}/{urcapID} | [update urcap by Id](#update-urcap-by-id) | Updates a urcap info by ID's |
  


## Paths

### <span id="add-urcap"></span> Adds a new urcap (*addUrcap*)

```
POST /api/v1/urcaps
```

This operation install urcaps from a zipped urcap file. This operation 
expects zipped .urcapx file to be uploaded as form data.


#### Consumes
  * multipart/form-data

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| urcapxFile | `formData` | file | `io.ReadCloser` |  | ✓ |  | The zipped urcapx file. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-urcap-201) | Created | metadata |  | [schema](#add-urcap-201-schema) |
| [default](#add-urcap-default) | | error response |  | [schema](#add-urcap-default-schema) |

#### Responses


##### <span id="add-urcap-201"></span> 201 - metadata
Status: Created

###### <span id="add-urcap-201-schema"></span> Schema
   
  

[Manifest](#manifest)

##### <span id="add-urcap-default"></span> Default Response
error response

###### <span id="add-urcap-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-urcap-by-id"></span> Deletes a urcaps info by ID's (*deleteUrcapById*)

```
DELETE /api/v1/urcaps/{vendorID}/{urcapID}
```

This operation uninstall a urcap.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-urcap-by-id-200) | OK | Updated urcap. |  | [schema](#delete-urcap-by-id-200-schema) |
| [default](#delete-urcap-by-id-default) | | error |  | [schema](#delete-urcap-by-id-default-schema) |

#### Responses


##### <span id="delete-urcap-by-id-200"></span> 200 - Updated urcap.
Status: OK

###### <span id="delete-urcap-by-id-200-schema"></span> Schema
   
  

[Manifest](#manifest)

##### <span id="delete-urcap-by-id-default"></span> Default Response
error

###### <span id="delete-urcap-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-artifact-actual-name-by-path"></span> Gets the actual name of an artifact from the resource path (*getArtifactActualNameByPath*)

```
GET /api/v1/artifacts/{artifactType}/name
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactType | `path` | string | `string` |  | ✓ |  | The type of the artifact |
| artifactPath | `query` | string | `string` |  | ✓ |  | Path to the resource/artifact on host machine |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-artifact-actual-name-by-path-200) | OK | artifact name |  | [schema](#get-artifact-actual-name-by-path-200-schema) |
| [default](#get-artifact-actual-name-by-path-default) | | error |  | [schema](#get-artifact-actual-name-by-path-default-schema) |

#### Responses


##### <span id="get-artifact-actual-name-by-path-200"></span> 200 - artifact name
Status: OK

###### <span id="get-artifact-actual-name-by-path-200-schema"></span> Schema
   
  



##### <span id="get-artifact-actual-name-by-path-default"></span> Default Response
error

###### <span id="get-artifact-actual-name-by-path-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-artifact-metadata-by-path"></span> Gets the urcap metadata from the resource path (*getArtifactMetadataByPath*)

```
GET /api/v1/artifacts/{artifactType}/info
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactType | `path` | string | `string` |  | ✓ |  | The type of the artifact |
| artifactPath | `query` | string | `string` |  | ✓ |  | Path to the resource/artifact on host machine |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-artifact-metadata-by-path-200) | OK | URCap metadata |  | [schema](#get-artifact-metadata-by-path-200-schema) |
| [default](#get-artifact-metadata-by-path-default) | | error |  | [schema](#get-artifact-metadata-by-path-default-schema) |

#### Responses


##### <span id="get-artifact-metadata-by-path-200"></span> 200 - URCap metadata
Status: OK

###### <span id="get-artifact-metadata-by-path-200-schema"></span> Schema
   
  

[Metadata](#metadata)

##### <span id="get-artifact-metadata-by-path-default"></span> Default Response
error

###### <span id="get-artifact-metadata-by-path-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-container-by-id"></span> Gets a container info by urcap ID's and artifactName (*getContainerById*)

```
GET /api/v1/containers/{vendorID}/{urcapID}/{artifactName}
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactName | `path` | string | `string` |  | ✓ |  | artifact name of container defined in manifest |
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-container-by-id-200) | OK | A single container. |  | [schema](#get-container-by-id-200-schema) |
| [default](#get-container-by-id-default) | | error |  | [schema](#get-container-by-id-default-schema) |

#### Responses


##### <span id="get-container-by-id-200"></span> 200 - A single container.
Status: OK

###### <span id="get-container-by-id-200-schema"></span> Schema
   
  

[Container](#container)

##### <span id="get-container-by-id-default"></span> Default Response
error

###### <span id="get-container-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-container-port-mapping-by-id"></span> Port mapping of a container (*getContainerPortMappingById*)

```
GET /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/portmapping
```

Gets host port mapped to the container port

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactName | `path` | string | `string` |  | ✓ |  | artifact name of container defined in manifest |
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-container-port-mapping-by-id-200) | OK | Container Port Mapping. |  | [schema](#get-container-port-mapping-by-id-200-schema) |
| [default](#get-container-port-mapping-by-id-default) | | error |  | [schema](#get-container-port-mapping-by-id-default-schema) |

#### Responses


##### <span id="get-container-port-mapping-by-id-200"></span> 200 - Container Port Mapping.
Status: OK

###### <span id="get-container-port-mapping-by-id-200-schema"></span> Schema

##### <span id="get-container-port-mapping-by-id-default"></span> Default Response
error

###### <span id="get-container-port-mapping-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-urcap-by-id"></span> Gets a urcap info by urcap ID's (*getUrcapById*)

```
GET /api/v1/urcaps/{vendorID}/{urcapID}
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-urcap-by-id-200) | OK | A single urcap. |  | [schema](#get-urcap-by-id-200-schema) |
| [default](#get-urcap-by-id-default) | | error |  | [schema](#get-urcap-by-id-default-schema) |

#### Responses


##### <span id="get-urcap-by-id-200"></span> 200 - A single urcap.
Status: OK

###### <span id="get-urcap-by-id-200-schema"></span> Schema
   
  

[Manifest](#manifest)

##### <span id="get-urcap-by-id-default"></span> Default Response
error

###### <span id="get-urcap-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-urcaps"></span> Gets all urcaps. (*getUrcaps*)

```
GET /api/v1/urcaps
```

This operation provides a view of installed urcaps's data in JSON.


#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-urcaps-200) | OK | OK |  | [schema](#get-urcaps-200-schema) |

#### Responses


##### <span id="get-urcaps-200"></span> 200 - OK
Status: OK

###### <span id="get-urcaps-200-schema"></span> Schema
   
  

[][Manifest](#manifest)

### <span id="start-container-by-id"></span> Start a container (*startContainerById*)

```
PUT /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/start
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactName | `path` | string | `string` |  | ✓ |  | artifact name of container defined in manifest |
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#start-container-by-id-200) | OK | OK |  | [schema](#start-container-by-id-200-schema) |
| [default](#start-container-by-id-default) | | error |  | [schema](#start-container-by-id-default-schema) |

#### Responses


##### <span id="start-container-by-id-200"></span> 200 - OK
Status: OK

###### <span id="start-container-by-id-200-schema"></span> Schema

##### <span id="start-container-by-id-default"></span> Default Response
error

###### <span id="start-container-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="status-container-by-id"></span> Status of a container (*statusContainerById*)

```
PUT /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/status
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactName | `path` | string | `string` |  | ✓ |  | artifact name of container defined in manifest |
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#status-container-by-id-200) | OK | OK |  | [schema](#status-container-by-id-200-schema) |
| [default](#status-container-by-id-default) | | error |  | [schema](#status-container-by-id-default-schema) |

#### Responses


##### <span id="status-container-by-id-200"></span> 200 - OK
Status: OK

###### <span id="status-container-by-id-200-schema"></span> Schema

##### <span id="status-container-by-id-default"></span> Default Response
error

###### <span id="status-container-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="stop-all-containers"></span> Stops all containers (*stopAllContainers*)

```
PUT /api/v1/containers/stop
```

TODO

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#stop-all-containers-200) | OK | Container Port Mapping. |  | [schema](#stop-all-containers-200-schema) |
| [default](#stop-all-containers-default) | | error |  | [schema](#stop-all-containers-default-schema) |

#### Responses


##### <span id="stop-all-containers-200"></span> 200 - Container Port Mapping.
Status: OK

###### <span id="stop-all-containers-200-schema"></span> Schema

##### <span id="stop-all-containers-default"></span> Default Response
error

###### <span id="stop-all-containers-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="stop-container-by-id"></span> Stop a container (*stopContainerById*)

```
PUT /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/stop
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactName | `path` | string | `string` |  | ✓ |  | artifact name of container defined in manifest |
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#stop-container-by-id-200) | OK | OK |  | [schema](#stop-container-by-id-200-schema) |
| [default](#stop-container-by-id-default) | | error |  | [schema](#stop-container-by-id-default-schema) |

#### Responses


##### <span id="stop-container-by-id-200"></span> 200 - OK
Status: OK

###### <span id="stop-container-by-id-200-schema"></span> Schema

##### <span id="stop-container-by-id-default"></span> Default Response
error

###### <span id="stop-container-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-urcap-by-id"></span> Updates a urcap info by ID's (*updateUrcapById*)

```
PUT /api/v1/urcaps/{vendorID}/{urcapID}
```

This operation updates an exsisting urcap.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-urcap-by-id-200) | OK | Updated urcap. |  | [schema](#update-urcap-by-id-200-schema) |
| [default](#update-urcap-by-id-default) | | error |  | [schema](#update-urcap-by-id-default-schema) |

#### Responses


##### <span id="update-urcap-by-id-200"></span> 200 - Updated urcap.
Status: OK

###### <span id="update-urcap-by-id-200-schema"></span> Schema
   
  

[Manifest](#manifest)

##### <span id="update-urcap-by-id-default"></span> Default Response
error

###### <span id="update-urcap-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="validate-container-by-id"></span> Validates container artifact presence (*validateContainerById*)

```
GET /api/v1/containers/{vendorID}/{urcapID}/{artifactName}/validate
```

TODO

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| artifactName | `path` | string | `string` |  | ✓ |  | artifact name of container defined in manifest |
| urcapID | `path` | string | `string` |  | ✓ |  | urcapID of urcap defined in manifest |
| vendorID | `path` | string | `string` |  | ✓ |  | vendorID of urcap defined in manifest |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#validate-container-by-id-200) | OK | Container Port Mapping. |  | [schema](#validate-container-by-id-200-schema) |
| [default](#validate-container-by-id-default) | | error |  | [schema](#validate-container-by-id-default-schema) |

#### Responses


##### <span id="validate-container-by-id-200"></span> 200 - Container Port Mapping.
Status: OK

###### <span id="validate-container-by-id-200-schema"></span> Schema

##### <span id="validate-container-by-id-default"></span> Default Response
error

###### <span id="validate-container-by-id-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

## Models

### <span id="container"></span> Container


> All the properties related to container run.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| devices | [][Device](#device)| `[]*Device` |  | | List of device to be mounted in container. |  |
| environment | [][EnvironmentVariable](#environment-variable)| `[]*EnvironmentVariable` |  | | List of environment variables inside the container. |  |
| hostaccess | string| `string` |  | | Container can communicate with host process through hosthost.internal dns name. This setting requires special priviledge |  |
| image | string| `string` | ✓ | | The image name:tag to be loaded. The loaded image should be found at path name/name:tag |  |
| ingress | [Ingress](#ingress)| `Ingress` |  | |  |  |
| name | string| `string` | ✓ | | Name of the container artifact | `daemon-cpp` |
| ports | [][Port](#port)| `[]*Port` |  | | list of open ports on the container |  |
| volumeMounts | [][VolumeMount](#volume-mount)| `[]*VolumeMount` |  | | List of volumeMounts to persist data. |  |



### <span id="environment-variable"></span> EnvironmentVariable


> Adds environment variables to the container
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` | ✓ | | The container environment variable name. Convention is to use capital letters with underscore delimiters. | `LOG_LEVEL` |
| value | string| `string` | ✓ | | The container environment variable value. Any configuration information required at runtime by urcap container. | `info` |



### <span id="error"></span> Error


> error model
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | string| `string` |  | | Error codes.This field contains a string succinctly identifying     the problem. | `missing_field` |
| message | string| `string` |  | | This field contain a plainly-written, developer-oriented explanation of the solution to the problem in complete, well-formed sentences. | `The `first_name` field is required.` |
| more_info | url (formatted string)| `string` |  | | This field SHOULD contain a publicly-accessible URL where information about the error can be read in a web browser. | `https://docs.api.example.com/v2/users/create_user#first_name` |
| target | [ErrorTarget](#error-target)| `ErrorTarget` |  | |  |  |



### <span id="error-response"></span> ErrorResponse


> An error response (any response with 400- or 500-series status code) MUST return an error container model. This model MUST contain an errors field, SHOULD contain a trace field, and MAY contain a status_code field, as outlined
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| errors | [][Error](#error)| `[]*Error` |  | | List of errors |  |
| status_code | uint32 (formatted integer)| `uint32` |  | | This field MAY contain the HTTP status code used for the response. Otherwise, it MUST be omitted. | `402` |
| trace | string| `string` |  | | This field SHOULD contain a lowercase UUID uniquely identifying the request. | `9daee671-916a-4678-850b-10b911f0236d` |



### <span id="error-target"></span> Error_target


> Error Target Model. This field MAY contain an error target model. Otherwise, it MUST be omitted.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | This field MUST contain the name of the problematic field (with dot-syntax if necessary), query parameter, or header. | `username` |
| type | string| `string` |  | | This field MUST contain field, parameter, or header | `field` |



### <span id="ingress"></span> Ingress


> Container ingress configuration. Automatically configure ingress path routing for ex. rest api.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| masterIndex | string| `string` |  | | The container should handle the master index. | `/index` |
| name | string| `string` | ✓ | | The ingress rule name. Logical name of ingress rule. | `rest-api` |
| port | uint32 (formatted integer)| `uint32` | ✓ | | The container port number of http server. This is the port number on the container of http server | `80` |
| protocol | string| `string` |  | `"http"`| The ingress protocol. Add this setting to select between http and websocket - the later allowing connection upgrades | `websocket` |
| proxyurl | string| `string` |  | `"/original"`| The proxy url forwarded to container. Select between short urls and long urls arriving at the container backend. When short urls are selected the matching prefix is removed from the url before forwarding. | `/shorturl` |



### <span id="manifest"></span> Manifest


> URCap Manifest Specification. The Manifest file is a YAML file defining a multi-artifacts based urcap application. The Manifest.yaml file is located on the top level directory/folder.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| apiVersion | string| `string` | ✓ | | Version (major.minor.patch) of the Manifest specification used. Tools not implementing required version MUST reject the configuration file. | `1.0.0` |
| artifacts | [ManifestArtifacts](#manifest-artifacts)| `ManifestArtifacts` |  | |  |  |
| metadata | [Metadata](#metadata)| `Metadata` | ✓ | |  |  |



### <span id="manifest-artifacts"></span> Manifest_artifacts


> All artifacts like containers, polyscopeBundle, webArchive that makes up an urcap.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| containers | [][Container](#container)| `[]*Container` |  | | List of container artifacts |  |
| polyscopeBundles | [][PolyscopeBundle](#polyscope-bundle)| `[]*PolyscopeBundle` |  | | List of osgi polyscope bundles |  |
| webArchives | [][WebArchive](#web-archive)| `[]*WebArchive` |  | | List of web archives |  |



### <span id="metadata"></span> Metadata


> Metadata of the urcap
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| copyright | string| `string` |  | | Specifies copy rights for the urcap | `Copyright (c) 2009-2021 Universal Robots. All rights reserved.` |
| description | string| `string` |  | | Short description of urcap | `Sample gripper URCap` |
| license | string| `string` |  | | License of various software/hardware used while developing the urcap | `TODO` |
| urcapID | string| `string` | ✓ | | Urcap id for this application. The ID must be unique for each urcap application developed by a company (vendorID). | `dockerdaemon` |
| urcapName | string| `string` | ✓ | | Urcap name of this application. This Will be displayed on user interface. | `Docker Daemon` |
| vendorID | string| `string` | ✓ | | Company urcap developer ID. The ID is shared among all urcap applications developed by company. | `universal-robots` |
| vendorName | string| `string` | ✓ | | Urcap name of this application. This Will be displayed on user interface. | `Universal Robots` |
| version | string| `string` | ✓ | | Urcap version (major.minor.patch) | `1.0.0` |



### <span id="polyscope-bundle"></span> PolyscopeBundle


> Configuration for PolyscopeBundle artifact
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bundle | string| `string` | ✓ | | The polyscope bundle path. Path to the jar file relative to the zipped urcap top folder. | `polyscope-ui-1.0-SNAPSHOT.jar` |
| name | string| `string` | ✓ | | The name of the polyscope bundle. | `polyscope-ui` |



### <span id="port"></span> Port


> These details are used to establish the connection between the docker host and container

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| containerPort | int32 (formatted integer)| `int32` | ✓ | | This is the port number on the container. The container port number is mapped to port name on container host | `40405` |
| name | string| `string` | ✓ | | This is logical port name on docker host. Port name is resolved to one of the available port number on the docker host and queried using rest api. The logical_port on docker host is mapped to containerPort on container | `xmlrpc` |
| protocol | string| `string` |  | `"tcp"`| Type of port mapping between the container host and container. TCP(Default)/UDP |  |



### <span id="web-archive"></span> WebArchive


> An web archive contains a folder with all web related pages.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| folder | string| `string` | ✓ | | Path of the folder with all web related pages. This path is relative to the zipped urcap top folder. | `mywebarchive` |
| masterIndex | string| `string` |  | | Path to the main index pages for the front to load. |  |
| name | string| `string` | ✓ | | The name of the web archive | `mywebarchive` |



### <span id="device"></span> device


> container device. Adds host device to container
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | | TODO | `Select the serial port with dongle atthed` |
| device | string| `string` | ✓ | | The host device path : the containert device path. 
The physical device mapping between host and container. The host part can contain regular expression if specific device is not known. | `/dev/ttyS[0-3]:/dev/serial:rw` |
| name | string| `string` |  | | The name of mapped device. Logical name of device | `Dongle interface` |



### <span id="volume-mount"></span> volumeMount


> Use to mount volumes into container. Data can be persistent or temporary according to selected type

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| mountPath | string| `string` | ✓ | | Absolute path for the data folders to be mounted on container. This folder is meant to have persistent data | `/gripper/data` |
| type | string| `string` |  | | Type of volume to mount. For tmpfs data will be lost following restart where as persistent will keep data until urcap is uninstalled. | `tmpfs` |


