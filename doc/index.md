# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/compose.proto](#proto_compose-proto)
    - [ContainerWithResourceUsage](#DockerController-ContainerWithResourceUsage)
    - [CreateProjectRequest](#DockerController-CreateProjectRequest)
    - [GetProjectResourceUsageResponse](#DockerController-GetProjectResourceUsageResponse)
    - [Project](#DockerController-Project)
    - [ProjectRequest](#DockerController-ProjectRequest)
  
    - [ComposeService](#DockerController-ComposeService)
  
- [proto/container.proto](#proto_container-proto)
    - [Container](#DockerController-Container)
    - [Container.LabelsEntry](#DockerController-Container-LabelsEntry)
    - [ContainerRequest](#DockerController-ContainerRequest)
    - [CreateContainerRequest](#DockerController-CreateContainerRequest)
  
    - [ContainerService](#DockerController-ContainerService)
  
- [proto/docker_controller.proto](#proto_docker_controller-proto)
    - [GetResourceUsageResponse](#DockerController-GetResourceUsageResponse)
    - [GetResourceUsageResponse.ContainerUsageEntry](#DockerController-GetResourceUsageResponse-ContainerUsageEntry)
    - [ListContainersResponse](#DockerController-ListContainersResponse)
    - [ListProjectsResponse](#DockerController-ListProjectsResponse)
    - [ResourceUsageMapValue](#DockerController-ResourceUsageMapValue)
  
    - [DockerService](#DockerController-DockerService)
  
- [proto/resources.messages.proto](#proto_resources-messages-proto)
    - [BlockIOStats](#DockerController-BlockIOStats)
    - [CpuStats](#DockerController-CpuStats)
    - [MemoryStats](#DockerController-MemoryStats)
    - [NetworkStats](#DockerController-NetworkStats)
    - [ResourceUsage](#DockerController-ResourceUsage)
  
- [Scalar Value Types](#scalar-value-types)



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






<a name="DockerController-GetProjectResourceUsageResponse"></a>

### GetProjectResourceUsageResponse



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
| CreateProject | [CreateProjectRequest](#DockerController-CreateProjectRequest) | [Project](#DockerController-Project) |  |
| GetProject | [ProjectRequest](#DockerController-ProjectRequest) | [Project](#DockerController-Project) |  |
| DeleteProject | [ProjectRequest](#DockerController-ProjectRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| StopProject | [ProjectRequest](#DockerController-ProjectRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| StartProject | [ProjectRequest](#DockerController-ProjectRequest) | [Project](#DockerController-Project) |  |
| RestartProject | [ProjectRequest](#DockerController-ProjectRequest) | [Project](#DockerController-Project) |  |
| GetProjectResourceUsage | [ProjectRequest](#DockerController-ProjectRequest) | [GetProjectResourceUsageResponse](#DockerController-GetProjectResourceUsageResponse) |  |
| GetProjectTotalResourceUsage | [ProjectRequest](#DockerController-ProjectRequest) | [ResourceUsage](#DockerController-ResourceUsage) |  |

 



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
| created_at | [string](#string) |  |  |
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





 

 

 


<a name="DockerController-ContainerService"></a>

### ContainerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateContainer | [CreateContainerRequest](#DockerController-CreateContainerRequest) | [Container](#DockerController-Container) |  |
| GetContainer | [ContainerRequest](#DockerController-ContainerRequest) | [Container](#DockerController-Container) |  |
| DeleteContainer | [ContainerRequest](#DockerController-ContainerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| PauseContainer | [ContainerRequest](#DockerController-ContainerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| StartContainer | [ContainerRequest](#DockerController-ContainerRequest) | [Container](#DockerController-Container) |  |
| RestartContainer | [ContainerRequest](#DockerController-ContainerRequest) | [Container](#DockerController-Container) |  |
| StopContainer | [ContainerRequest](#DockerController-ContainerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetContainerResourceUsage | [ContainerRequest](#DockerController-ContainerRequest) | [ResourceUsage](#DockerController-ResourceUsage) |  |

 



<a name="proto_docker_controller-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/docker_controller.proto



<a name="DockerController-GetResourceUsageResponse"></a>

### GetResourceUsageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| container_usage | [GetResourceUsageResponse.ContainerUsageEntry](#DockerController-GetResourceUsageResponse-ContainerUsageEntry) | repeated |  |






<a name="DockerController-GetResourceUsageResponse-ContainerUsageEntry"></a>

### GetResourceUsageResponse.ContainerUsageEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ResourceUsageMapValue](#DockerController-ResourceUsageMapValue) |  |  |






<a name="DockerController-ListContainersResponse"></a>

### ListContainersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| containers | [Container](#DockerController-Container) | repeated |  |






<a name="DockerController-ListProjectsResponse"></a>

### ListProjectsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [Project](#DockerController-Project) | repeated |  |






<a name="DockerController-ResourceUsageMapValue"></a>

### ResourceUsageMapValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image | [string](#string) |  |  |
| resource_usage | [ResourceUsage](#DockerController-ResourceUsage) |  |  |





 

 

 


<a name="DockerController-DockerService"></a>

### DockerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Ping | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetResourceUsage | [.google.protobuf.Empty](#google-protobuf-Empty) | [GetResourceUsageResponse](#DockerController-GetResourceUsageResponse) |  |
| GetTotalResourceUsage | [.google.protobuf.Empty](#google-protobuf-Empty) | [ResourceUsage](#DockerController-ResourceUsage) |  |
| ListContainers | [.google.protobuf.Empty](#google-protobuf-Empty) | [ListContainersResponse](#DockerController-ListContainersResponse) |  |
| ListProjects | [.google.protobuf.Empty](#google-protobuf-Empty) | [ListProjectsResponse](#DockerController-ListProjectsResponse) |  |

 



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
| name | [string](#string) |  |  |
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

