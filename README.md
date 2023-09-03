[![Docker Image CI](https://github.com/marcusziade/MoodLint/actions/workflows/docker-image.yml/badge.svg)](https://github.com/marcusziade/MoodLint/actions/workflows/docker-image.yml)
[![Go](https://github.com/marcusziade/MoodLint/actions/workflows/go.yml/badge.svg)](https://github.com/marcusziade/MoodLint/actions/workflows/go.yml)
[![Imperative PR Title Check](https://github.com/marcusziade/MoodLint/actions/workflows/check-pr-title-imperative.yml/badge.svg)](https://github.com/marcusziade/MoodLint/actions/workflows/check-pr-title-imperative.yml)

---

# 🌟 MoodLint 🌟

Tired of messy PR titles? Want to enforce a consistent style across your team? MoodLint is here to rescue your git workflow! With MoodLint, say goodbye to irrelevant and unprofessional PR titles. Our tool validates the mood, prefix, and overall structure of your PR titles, ensuring they're up to the mark.

Let MoodLint bring peace and order to your PR chaos! Give it a star if you find it useful.

---

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Configuration](#configuration)
5. [Development](#development)
6. [Contributing](#contributing)
7. [License](#license)

## 🔥 Features:
- Validates PR title prefix for predefined types like `feat`, `fix`, `chore`, and more!
- Checks if your PR title and description are phrased in the **imperative mood**.
- Exits with a non-zero code for any validation failures, making it perfect for CI/CD pipelines.
- Comprehensive reporting and examples to guide you in case of an error.

## 🛠️ Installation

To install MoodLint as a GitHub Action, add the following YAML configuration to your repository under `.github/workflows/moodlint.yml`.

```yaml
name: MoodLint PR Title Check

on:
  pull_request:
    types: [opened, synchronize, reopened, edited, ready_for_review]

jobs:
  check-title:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v2
    - name: Run MoodLint
      uses: your-username/MoodLint@v1.0.0
```

Replace `your-username` with your GitHub username and adjust the version tag as needed.

## 👨‍💻 Usage

Once the GitHub Action is set up, it will automatically run whenever a Pull Request is opened or updated. If the title does not comply with the imperative mood or valid prefixes, the Action will fail and log an error message.

## ⚙️ Configuration

By default, MoodLint validates PR titles against the following allowed prefixes:

- `feat`
- `fix`
- `chore`
- `refactor`
- `regfix`

If you'd like to customize these prefixes, you can modify the `allowedPrefixes` variable inside the `src/validator/prefix_validator.go` file.

## 🐳 Running with Docker

You can also run MoodLint as a Docker container. Here's how to do it:

1. Build the Docker image:
    ```bash
    docker build -t moodlint .
    ```

2. Run the Docker container:
    ```bash
    docker run moodlint "Your PR Title Here"
    ```

This will execute MoodLint and check the given PR title.

## 📠 Development

To work on the MoodLint project:

1. Clone the repository
    ```
    git clone https://github.com/your-username/MoodLint.git
    ```
2. Navigate to the project directory
    ```
    cd MoodLint
    ```
3. Install dependencies
    ```
    go mod tidy
    ```

Run tests, if any:
```
go test ./...
```

Build the project:
```
go build
```

## 🤝 Contributing

If you'd like to contribute, please fork the repository and create a pull request, or simply open an issue for bugs and feature requests.

## 📝 License

MIT
