# DockerController
### Простой микросервис для управления docker'ом

---

## Стек:
- Go
- gRPC

## Краткое описание:
### [Ссылка на полную документацию protobuf](https://github.com/XenonPPG/Docker-Controller/blob/master/doc/index.md)

### 1. ContainerService

Предназначен для точечного управления отдельными Docker-контейнерами. Обеспечивает полный контроль над их жизненным циклом (запуск, остановка, пауза) и мониторинг их состояния.

**Предоставляемые методы:**

* `CreateContainer`
* `GetContainer`
* `DeleteContainer`
* `PauseContainer`
* `StartContainer`
* `RestartContainer`
* `StopContainer`
* `GetContainerResourceUsage`

### 2. ComposeService

Отвечает за управление группами контейнеров через сущность "Проект" (аналог Docker Compose). Позволяет контролировать жизненный цикл проектов и собирать метрики использования ресурсов для всей группы контейнеров.

**Предоставляемые методы:**

* `CreateProject`
* `GetProject`
* `DeleteProject`
* `StopProject`
* `StartProject`
* `RestartProject`
* `GetProjectResourceUsage`
* `GetProjectTotalResourceUsage`

### 3. DockerService

Служит для выполнения глобальных операций на уровне всего Docker-демона. Позволяет получать агрегированную статистику по ресурсам, списки запущенных сущностей (контейнеров и проектов), а также проводить очистку системы.

**Предоставляемые методы:**

* `Ping`
* `GetResourceUsage`
* `GetTotalResourceUsage`
* `ListContainers`
* `ListProjects`
* `CleanUp`