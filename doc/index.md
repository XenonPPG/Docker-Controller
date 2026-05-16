# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/DockerController.proto](#proto_DockerController-proto)
    - [DockerInfo](#DockerController-DockerInfo)
    - [DockerVersion](#DockerController-DockerVersion)
    - [Event](#DockerController-Event)
    - [Event.AttributesEntry](#DockerController-Event-AttributesEntry)
    - [StreamEventsRequest](#DockerController-StreamEventsRequest)
  
    - [DockerService](#DockerController-DockerService)
  
- [proto/compose.proto](#proto_compose-proto)
    - [ContainerWithResourceUsage](#DockerController-ContainerWithResourceUsage)
    - [CreateProjectRequest](#DockerController-CreateProjectRequest)
    - [GetResourceUsageResponse](#DockerController-GetResourceUsageResponse)
    - [Project](#DockerController-Project)
    - [ProjectRequest](#DockerController-ProjectRequest)
  
    - [ComposeService](#DockerController-ComposeService)
  
- [proto/container.proto](#proto_container-proto)
    - [Container](#DockerController-Container)
    - [Container.LabelsEntry](#DockerController-Container-LabelsEntry)
    - [ContainerRequest](#DockerController-ContainerRequest)
    - [CreateContainerRequest](#DockerController-CreateContainerRequest)
  
    - [DockerController](#DockerController-DockerController)
  
- [proto/resources.messages.proto](#proto_resources-messages-proto)
    - [BlockIOStats](#DockerController-BlockIOStats)
    - [CpuStats](#DockerController-CpuStats)
    - [GetContainerResourceUsageRequest](#DockerController-GetContainerResourceUsageRequest)
    - [MemoryStats](#DockerController-MemoryStats)
    - [NetworkStats](#DockerController-NetworkStats)
    - [ResourceUsage](#DockerController-ResourceUsage)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_DockerController-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/DockerController.proto



<a name="DockerController-DockerInfo"></a>

### DockerInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| containers_running | [int32](#int32) |  |  |
| containers_stopped | [int32](#int32) |  |  |
| images_count | [int32](#int32) |  |  |
| os | [string](#string) |  |  |
| architecture | [string](#string) |  |  |
| total_memory | [uint64](#uint64) |  |  |
| cpus | [int32](#int32) |  |  |






<a name="DockerController-DockerVersion"></a>

### DockerVersion



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| api_version | [string](#string) |  |  |
| go_version | [string](#string) |  |  |






<a name="DockerController-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| action | [string](#string) |  |  |
| actor_id | [string](#string) |  |  |
| actor_name | [string](#string) |  |  |
| time | [int64](#int64) |  |  |
| attributes | [Event.AttributesEntry](#DockerController-Event-AttributesEntry) | repeated |  |






<a name="DockerController-Event-AttributesEntry"></a>

### Event.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="DockerController-StreamEventsRequest"></a>

### StreamEventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| types | [string](#string) | repeated |  |
| actions | [string](#string) | repeated |  |





 

 

 


<a name="DockerController-DockerService"></a>

### DockerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetInfo | [.google.protobuf.Empty](#google-protobuf-Empty) | [DockerInfo](#DockerController-DockerInfo) |  |
| GetVersion | [.google.protobuf.Empty](#google-protobuf-Empty) | [DockerVersion](#DockerController-DockerVersion) |  |
| GetResourceUsage | [.google.protobuf.Empty](#google-protobuf-Empty) | [ResourceUsage](#DockerController-ResourceUsage) |  |
| StreamEvents | [StreamEventsRequest](#DockerController-StreamEventsRequest) | [Event](#DockerController-Event) stream |  |
| Ping | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |

 



<a name="proto_compose-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/compose.proto



<a name="DockerController-ContainerWithResourceUsage"></a>

### ContainerWithResourceUsage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| container | [Container](#DockerController-Container) |  |  |
| resource_usage | [ResourceUsage](#DockerController-ResourceUsage) |  |  |






<a name="DockerController-CreateProjectRequest"></a>

### CreateProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| compose_file | [string](#string) |  |  |






<a name="DockerController-GetResourceUsageResponse"></a>

### GetResourceUsageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stats | [ContainerWithResourceUsage](#DockerController-ContainerWithResourceUsage) | repeated |  |






<a name="DockerController-Project"></a>

### Project



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| containers | [Container](#DockerController-Container) | repeated |  |






<a name="DockerController-ProjectRequest"></a>

### ProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |





 

 

 


<a name="DockerController-ComposeService"></a>

### ComposeService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateProjectRequest](#DockerController-CreateProjectRequest) | [Project](#DockerController-Project) |  |
| Get | [ProjectRequest](#DockerController-ProjectRequest) | [Project](#DockerController-Project) |  |
| Delete | [ProjectRequest](#DockerController-ProjectRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| Stop | [ProjectRequest](#DockerController-ProjectRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| Start | [ProjectRequest](#DockerController-ProjectRequest) | [Project](#DockerController-Project) |  |
| Restart | [ProjectRequest](#DockerController-ProjectRequest) | [Project](#DockerController-Project) |  |
| GetTotalResourceUsage | [ProjectRequest](#DockerController-ProjectRequest) | [ResourceUsage](#DockerController-ResourceUsage) |  |
| GetResourceUsage | [ProjectRequest](#DockerController-ProjectRequest) | [GetResourceUsageResponse](#DockerController-GetResourceUsageResponse) |  |

 



<a name="proto_container-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/container.proto



<a name="DockerController-Container"></a>

### Container



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| image | [string](#string) |  |  |
| status | [string](#string) |  |  |
| state | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |
| ports | [string](#string) | repeated |  |
| labels | [Container.LabelsEntry](#DockerController-Container-LabelsEntry) | repeated |  |






<a name="DockerController-Container-LabelsEntry"></a>

### Container.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="DockerController-ContainerRequest"></a>

### ContainerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="DockerController-CreateContainerRequest"></a>

### CreateContainerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image | [string](#string) |  |  |
| name | [string](#string) |  |  |





 

 

 


<a name="DockerController-DockerController"></a>

### DockerController


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateContainerRequest](#DockerController-CreateContainerRequest) | [Container](#DockerController-Container) |  |
| Get | [ContainerRequest](#DockerController-ContainerRequest) | [Container](#DockerController-Container) |  |
| Delete | [ContainerRequest](#DockerController-ContainerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| Start | [ContainerRequest](#DockerController-ContainerRequest) | [Container](#DockerController-Container) |  |
| Restart | [ContainerRequest](#DockerController-ContainerRequest) | [Container](#DockerController-Container) |  |
| Stop | [ContainerRequest](#DockerController-ContainerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetResourceUsage | [GetContainerResourceUsageRequest](#DockerController-GetContainerResourceUsageRequest) | [ResourceUsage](#DockerController-ResourceUsage) |  |

 



<a name="proto_resources-messages-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/resources.messages.proto



<a name="DockerController-BlockIOStats"></a>

### BlockIOStats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| read_bytes | [uint64](#uint64) |  |  |
| write_bytes | [uint64](#uint64) |  |  |






<a name="DockerController-CpuStats"></a>

### CpuStats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| percent | [double](#double) |  |  |
| online_cpus | [int32](#int32) |  |  |






<a name="DockerController-GetContainerResourceUsageRequest"></a>

### GetContainerResourceUsageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="DockerController-MemoryStats"></a>

### MemoryStats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| usage | [uint64](#uint64) |  |  |
| limit | [uint64](#uint64) |  |  |
| percent | [double](#double) |  |  |






<a name="DockerController-NetworkStats"></a>

### NetworkStats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface | [string](#string) |  |  |
| rx_bytes | [uint64](#uint64) |  |  |
| tx_bytes | [uint64](#uint64) |  |  |






<a name="DockerController-ResourceUsage"></a>

### ResourceUsage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cpu | [CpuStats](#DockerController-CpuStats) |  |  |
| memory | [MemoryStats](#DockerController-MemoryStats) |  |  |
| networks | [NetworkStats](#DockerController-NetworkStats) | repeated |  |
| block_io | [BlockIOStats](#DockerController-BlockIOStats) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

