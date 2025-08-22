# 📡 IoT Device Simulator

## 🎯 Problem Statement
Developers working on **end-to-end IoT solutions** often face a common challenge:  
They need to build and test backend/frontend software *before* real devices are available.  
This project solves that problem by allowing developers to **simulate IoT devices** and their behaviors, producing realistic data streams that can be consumed by software under development.

---

## 📂 Project Overview
The **IoT Device Simulator** is a CLI tool that simulates IoT devices and publishes messages to an MQTT broker or a webhook api endpoint.  
It provides configurable simulation options for message frequency, device behavior, and MQTT publishing.

---

## ⚙️ Configuration

The project uses two YAML files:

### `config.yaml`
This file defines simulator-level configuration.

```yaml
send-mqtt: true
mqtt:
  mqtt-host: localhost
  mqtt-port: 1883
  mqtt-user: user
  mqtt-password: user

# if send-mqtt is false provide a webhook endpoint
api:
  endpoint: http://localhost:3000/send-data/
  api-key-header-name: x-api-key
  api-key: sk_test_8f93b2a7c4d14f2a9e8d1c5b7a9e3f12

```

**Fields:**
- `send-mqtt` → Boolean. If `true`, messages are sent to the configured MQTT broker otherwise it will send to the webhook api.
- `mqtt-host` → MQTT broker host.
- `mqtt-port` → MQTT broker port.
- `mqtt-username` → Username for MQTT authentication (optional).
- `mqtt-password` → Password for MQTT authentication (optional).
- `endpoint` → webhook api endpoint (optional).
- `api-key-header-name` → webhook api endpoint header (optional).
- `api-key` → webhook api key (optional).

---

### `specs.yaml`
This file defines the devices and their message simulation behaviors.

```yaml
messages:
  - title: msg_1
    device: device_a
    type: string
    options:
      - first_msg
      - second_msg
      - third_msg
    rate: once
  - title: msg_2
    device: device_b
    type: int
    rate: 2pm
  - title: msg_3
    device: device_b
    type: int
    rate: 2pm
    max: 100
    min: 0
  - title: msg_4
    device: device_c
    type: object
    body:
    - title: msg_4_1
      type: int
      max: 100
      min: 0
    - title: msg_4_2
      type: decimal
      max: 200
      min: 1
    rate: 10pm
```

**Fields:**
- `topic` → MQTT topic the message will be published to.
- `title` → msg title in the json.
- `rate` → Frequency of message publishing. Supports:
  - `once` → Sends the message one time only.
  - `Xps` → Sends message `X` times per second.
  - `Xpm` → Sends message `X` times per minute.
  - `Xph` → Sends message `X` times per hour.
  - `Xpd` → Sends message `X` times per day.
- `type` → string, int, decimal, boolean, object.
- `options` → if type is string you need to provide a list of options to send from.
- `max` → if type is a number you need to provide range.
- `min` → if type is a number you need to provide range.
- `body` → if type is an object you need to provide the body of the messages to include inside this object.
---

## 📊 Where does iot sim fit in your stack

```
             ┌──────────────────┐
             │  specs.yaml      │
             │ (device behavior)│
             └─────────┬────────┘
                       │
                       ▼
             ┌──────────────────┐
             │   Simulator CLI  │
             │      (Go)        │
             └─────────┬────────┘
                       │
                       ▼
             ┌──────────────────┐
             │ MQTT Broker      │
             │  or webhook api  │
             └─────────┬────────┘
                       │
                       ▼
             ┌──────────────────┐
             │  Your Software   │
             │(Backend/UI/etc.) │
             └──────────────────┘
```

---

## 🚀 Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/dashify-it/iot-simulator.git
   cd iot-sim
   ```

2. Build the project:
   ```bash
   go build -o simulator
   ```

3. Run the simulator with configs:
   ```bash
   ./simulator --config=config.yaml --specs=specs.yaml
   ```

4. Show help:
   ```bash
   ./simulator --help
   ```

---


## ✅ Who Is This For?
- IoT developers building **backend APIs** that need to process device data.
- Frontend engineers working on **dashboards** that visualize IoT data.
- QA engineers who need to **stress test systems** with simulated device traffic.

---

## 📌 Roadmap
- Support for protocols beyond MQTT (HTTP, CoAP, etc.).
- Add ui in place of the config and specs files.
- Provide a mocking api for the testing frontend only.
