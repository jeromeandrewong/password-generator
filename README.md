# Password Generator

A simple password generator written in Go that allows users to create customised passwords on the command line based on specific requirements such as length, uppercase/lowercase characters, numbers and special characters.

## Demo

![demo](/demo.gif)

## User Features

_Customizable Passwords_:
Generate passwords with specified lengths and character requirements.

_Flexibility_:
Control the number of uppercase, lowercase, numbers, and special characters in the password.

_Scramble Functionality_:
Shuffle the generated password for enhanced security.

## Project Structure

The project follows a structured organization for improved readability, maintainability, and scalability:

```javascript
├── README.md
├── go.mod
├── main.go
└── password
    ├── generator_test.go
    ├── generator.go // core logic for generating passwords based on user-defined requirements.
    └── requirements_test.go
    ├── requirements.go // functions related to handling user input for password criteria
    ├── utils_test.go
    ├── utils.go // helper functions
```

## Architecture Principles

**Dependency Inversion**

- dependency injection enables decoupling of high-level modules from their dependencies, allowing for easier testing and modification.
  - `InputGetter` Interface: Defined in `utils.go`, this interface provides a contract for obtaining user input.
  - `RealInputGetter`: Implements the `InputGetter` interface and serves as the real input retrieval mechanism.
  - `MockInputGetter`: Acts as a stub/mock for testing purposes, enabling controlled input simulation during testing.

**Separation of Concerns**

- The core logic for password generation, input handling and scrambling is encapsulated within distinct components.

## Testing

Execute the following command in the project directory to run tests

```bash
go test ./..
```
