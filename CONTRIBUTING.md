# Contribution Guidelines for Pecklin

Welcome to the Pecklin project! We're excited to have you contribute to our touch typing app. Whether you're fixing bugs, improving documentation, or adding new features, your contributions are greatly appreciated.

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [Getting Started](#getting-started)
3. [Reporting Issues](#reporting-issues)
4. [Submitting Contributions](#submitting-contributions)
5. [Coding Standards](#coding-standards)
6. [Testing Your Changes](#testing-your-changes)
7. [Commit Messages](#commit-messages)
8. [Review Process](#review-process)
9. [Documentation](#documentation)
10. [License](#license)

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). Be respectful and constructive in all communications.

## Getting Started

1. **Fork the Repository**  
   Fork the repository to your GitHub account and clone it locally:
   ```bash
   git clone https://github.com/your-username/pecklin.git
   cd pecklin
   ```

2. **Set Up Your Development Environment**  
   Ensure you have Go installed on your system. You can verify this by running:
   ```bash
   go version
   ```
   Install any dependencies:
   ```bash
   go mod tidy
   ```

3. **Understand the Codebase**  
   Review the project structure and read the documentation in the `docs/` folder. Look at the open issues and project roadmap to understand where help is needed.

## Reporting Issues

If you find a bug or have a suggestion, create an issue in the [Issues](https://github.com/your-username/pecklin/issues) tab. Provide as much detail as possible:
- Steps to reproduce
- Expected behavior
- Actual behavior
- Environment (OS, Go version, etc.)

## Submitting Contributions

1. **Create a Branch**  
   Use a descriptive name for your branch:
   ```bash
   git checkout -b feature/short-description
   ```

2. **Make Your Changes**  
   Follow the [Coding Standards](#coding-standards) below. Ensure your changes are modular and documented.

3. **Run Tests**  
   Write and run tests to verify your changes:
   ```bash
   go test ./...
   ```

4. **Push Your Branch**  
   Push your changes to your fork:
   ```bash
   git push origin feature/short-description
   ```

5. **Submit a Pull Request**  
   Open a pull request (PR) on the main repository. Provide a clear description of your changes and link to any related issues.

## Coding Standards

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Use `go fmt` to format your code:
  ```bash
  go fmt ./...
  ```
- Write clear, concise, and reusable code.
- Add comments for complex logic and exported functions.

## Testing Your Changes

- Add unit tests for new functions or features in the `*_test.go` files.
- Use the standard Go testing framework.
- Ensure all tests pass before submitting:
  ```bash
  go test ./...
  ```

## Commit Messages

Use descriptive commit messages following this format:
```
<type>: <short description>

<optional longer description>
```

Types include:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation update
- `test`: Testing changes
- `refactor`: Code restructuring
- `chore`: Maintenance tasks

Example:
```
feat: add support for saving typing progress

Added functionality to save typing session progress locally using JSON files.
```

## Review Process

1. Your PR will be reviewed by the maintainers or other contributors.
2. Be open to feedback and update your PR if requested.
3. Once approved, your PR will be merged into the main branch.

## Documentation

- Update documentation in the `docs/` folder for any new features.
- If your changes impact the README or other user-facing files, update them accordingly.

### Project structure

```

├── data/                              # Data access layer (local and remote).
│   ├── local/                         # Local data management.
│   │   └── database/                  # Database initialization and access logic.
│   │       ├── dao.go                 # Interfaces for database operations.
│   │       └── initializeDatabase.go  # Logic for initializing the database.
│   └── remote/                        # Remote API interaction.
│       └── remoteRepository.go        # Logic for accessing remote data sources.
├── domain/                            # Core business logic and data models.
│   └── models/                        # Domain models representing core entities.
├── main.go                            # Application startup logic.
├── pkg/                               # Public reusable packages for controllers and utilities.
│   ├── controllers/                   # Controllers for managing application logic.
│   └── utils/                         # Shared utilities for the project.
├── presentation/                      # Presentation layer for UI or CLI interaction.

```

## License

By contributing to Pecklin, you agree that your contributions will be licensed under the same [license](LICENSE) as the project.

Thank you for contributing to Pecklin! We look forward to your input. 🎉
