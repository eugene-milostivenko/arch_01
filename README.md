# warmhouse  project

# Задание 1. Анализ и планирование

## 1. Описание функциональности монолитного приложения

### Управление отоплением:

- Пользователи могут:
  - Включать отопление
  - Выключать отопление
  - Устанавливать температуру на термостате
- Система поддерживает:
  - Передачу заданных команд на систему отопления

### Мониторинг температуры:

- Пользователи могут:
  - Просматривать текущую температуру через веб-интерфейс
- Система поддерживает:
  - Запрос температуры с термостата и отображение её в веб-интерфейсе

### Функциональные ограничения:
- Пользователи не могут самостоятельно подключать новые устройства к системе
- Система не позволяет управлять другими умными устройствами, кроме термостатов

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go.
- База данных: PostgreSQL.
- Архитектура: Монолитная, все компоненты системы (обработка запросов, бизнес-логика, работа с данными) находятся в рамках одного приложения.
- Взаимодействие: Синхронное, запросы обрабатываются последовательно.
- Масштабируемость: Ограничена, так как монолит сложно масштабировать по частям.
- Развертывание: Требует остановки всего приложения.

### 3. Определение доменов и границы контекстов

- Управление умными устройствами
  - Регистрация и подключение устройств
  - Отправка команд на устройства
  - [IoT-инфраструктура] MQTT-брокер — приём телеметрии от устройств и маршрутизация событий к интеграционным сервисам
  - [ACL] Интеграция с вендорами
    - Интеграция с вендором 1
    - Интеграция с вендором 2

- Автоматизация
  - Управление сценариями автоматизации
  - Движок автоматизации: обработка триггеров, планирование и выполнение сценариев

- Телеметрия
  - Получение данных с устройств
  - Хранение исторических данных
  - Анализ данных
  - Видео стриминг

- Управление пользователями и контроль доступа
  - Регистрация пользователей
  - Аутентификация и авторизация
  - Управление профилем

- Поддержка пользователей [внешний домен, вне скоупа]
  - Регистр запросов пользователей
  - Чат с пользователем
  - Статистика обращений

### 4. Проблемы монолитного решения

- Сложность развёртывания: любое изменение требует остановки и перезапуска всего приложения, обновить отдельный модуль невозможно
- Сложность масштабирования: с ростом продукта потребуется найм новых сотрудников и выделение команд под отдельные домены, монолит не будет соответствовать структуре организации
- Снижение надёжности: сбой в одном модуле приводит к падению всей системы, с ростом функционала этот риск только увеличивается
- Отсутствие изоляции данных: все домены работают с единой базой данных, поэтому изменение схемы в одном месте может затронуть другие части системы
- Сложность тестирования: модули нельзя проверить в изоляции, для запуска тестов нужно поднимать всё приложение целиком
- Технологический lock-in: стек вынужденно однороден, нельзя выбрать подходящий инструмент для конкретного домена, например ClickHouse для телеметрии или специализированный медиасервер для видеостриминга


### 5. Визуализация контекста системы — диаграмма С4

![Диаграмма контекста](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/context.puml)

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**
![Диаграмма контейнеров](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/containers.puml)

**Диаграмма компонентов (Components)**

Device Management Service:
![Диаграмма компонентов — Device Management Service](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/components/device_management_service.puml)

Automation Executor Service:
![Диаграмма компонентов — Automation Executor Service](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/components/automation_executor_service.puml)

**Диаграмма кода (Code)**

Детализирован самый критичный участок To-Be системы — оркестрация выполнения шагов сценария автоматизации в Automation Executor Service.

Диаграмма классов — Automation Executor Service:
![Диаграмма классов — Automation Executor Service](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/code/automation_executor_class.puml)

Диаграмма последовательности — выполнение сценария автоматизации:
![Диаграмма последовательности — выполнение сценария автоматизации](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/code/scenario_execution_sequence.puml)

# Задание 3. Разработка ER-диаграммы

[diagrams/er.puml](diagrams/er.puml)

![ER-диаграмма](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/eugene-milostivenko/arch_01/refs/heads/warmhouse/diagrams/er.puml)

# Задание 4. Создание и документирование API

### 1. Тип API

Для взаимодействия микросервисов используются два типа API:

**REST API (OpenAPI 3.0)** — для синхронного взаимодействия между сервисами, когда вызывающей стороне необходим немедленный ответ. Используется для:
- Запросов от API Gateway к бэкенд-сервисам (Device Management, Telemetry, Automation Management)
- Межсервисных вызовов, где требуется подтверждение выполнения (например, Automation Engine → Device Management для отправки команды на устройство)

**AsyncAPI** — для асинхронного взаимодействия через Event Queue (Kafka). Используется для:
- Публикации телеметрических событий от Provider Integration Services — подписчикам не нужно получать ответ немедленно, а данные поступают с высокой частотой
- Потребления событий Telemetry Service (для сохранения) и Automation Engine Service (для оценки триггеров)

Асинхронный подход для телеметрии выбран потому, что:
1. Данные с устройств поступают непрерывным потоком, синхронная обработка создает риск bottleneck
2. Потребителей может быть несколько (Telemetry Service, Automation Engine)
3. Kafka обеспечивает буферизацию при пиковых нагрузках и гарантию доставки

### 2. Документация API

**REST API (OpenAPI 3.0):**
- [Device Management Service API](api/device-management-api.yaml) — получение информации об устройстве, отправка команд
- [Telemetry Service API](api/telemetry-api.yaml) — запрос исторических измерений телеметрии
- [Automation Management Service API](api/automation-management-api.yaml) — создание сценариев автоматизации

**AsyncAPI:**
- [Telemetry Events](api/telemetry-events-asyncapi.yaml) — событие получения новой телеметрии с устройства (Kafka topic `device.telemetry.received`)

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


