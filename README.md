# Orchestrator CLI

The `orchestrator-cli` is a powerful command-line interface (CLI) tool designed to streamline the initial setup of production-ready CI/CD architectures for your software projects. It automates the process of generating essential configuration files tailored to your project's specific technology stack and deployment needs.

## Features

-   **Automated Project Detection:** Automatically identifies your project's archetype (e.g., PHP Laravel, Java Spring Boot, Python FastAPI, NodeJS NextJS) and its associated language version.
-   **Interactive Configuration:** Guides you through an interactive process to gather crucial details, such as your preferred database type (MySQL, PostgreSQL, MongoDB, or custom) and your target deployment environment (on-premise or cloud).
-   **Customizable Template Generation:** Generates highly customized architectural files, including `docker-compose.yml`, `Dockerfile`, Kubernetes deployment configurations, and GitHub Actions CI/CD pipelines, all based on the detected project type and your specific inputs.
-   **GitHub Integration:** Offers optional features like applying branch protection rules directly on your GitHub repository.

## Getting Started

### Installation

To get started with `orchestrator-cli`, you can build it from source:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Suprath/orchestrator-cli.git
    cd orchestrator-cli
    ```
2.  **Build the executable:**
    ```bash
    go build -o orchestrator-cli
    ```
    This will create an `orchestrator-cli` executable in your current directory.

### Usage

Navigate to the root directory of your project (the one you want to initialize with CI/CD) and run the `init` command:

```bash
./orchestrator-cli init
```

The CLI will then guide you through a series of prompts:

1.  **Project Type Detection:** It will first scan your current directory to detect the project's archetype and language version.
2.  **Application Name:** You'll be asked to provide a short, lowercase name for your application.
3.  **Database Type:** Choose from a list of common databases (MySQL, PostgreSQL, MongoDB) or select 'Custom' to specify your own.
4.  **Deployment Environment:** Select whether your application will be deployed 'On-Premise' or to the 'Cloud'.
5.  **GitHub Branch Protection (Optional):** You'll have the option to apply branch protection rules to your GitHub repository.

Upon completion, `orchestrator-cli` will generate the necessary architectural files in your project directory, ready for review and commitment to your version control system.

## Contributing

We welcome contributions to `orchestrator-cli`! If you'd like to contribute, please follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes and ensure tests pass.
4.  Submit a pull request with a clear description of your changes.
