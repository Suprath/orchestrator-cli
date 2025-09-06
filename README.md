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

## Comprehensive Guide

# Orchestrator CLI: Docker Setup, CI/CD, and Deployment Guide

This guide provides comprehensive instructions on how to set up and use the `orchestrator-cli` within a Docker environment, trigger CI/CD pipelines, and understand the generated Terraform and Kubernetes configurations.

## 1. Running Orchestrator CLI with Docker

To run the `orchestrator-cli` Docker image and initialize your project, use the following command:

```bash
docker run --rm -it \
  -e GITHUB_TOKEN=<YOUR_GITHUB_TOKEN> \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v "$(pwd)":/app \
  suprathps/orchestrator:latest init
```

**Explanation of the command:**

*   `docker run`: Executes a Docker container.
*   `--rm`: Automatically removes the container when it exits.
*   `-it`: Runs the container in interactive mode with a pseudo-TTY, allowing you to interact with the CLI prompts.
*   `-e GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>`: Passes your GitHub Personal Access Token as an environment variable to the container. This is crucial for the CLI to interact with GitHub (e.g., for branch protection). **Remember to replace `<YOUR_GITHUB_TOKEN>` with your actual token.**
*   `-v /var/run/docker.sock:/var/run/docker.sock`: Mounts the Docker daemon's Unix socket into the container. This allows the `orchestrator-cli` (if it were to interact with Docker directly, e.g., for `docker compose build`) to communicate with the host's Docker daemon.
*   `-v "$(pwd)":/app`: Mounts your current working directory on the host machine into the `/app` directory inside the container. This is where the `orchestrator-cli` will detect your project and generate the architectural files.
*   `suprathps/orchestrator:latest`: Specifies the Docker image to use.
*   `init`: The command to execute within the container, which is the `orchestrator-cli`'s initialization command.

### How to Obtain a GitHub Personal Access Token

You can generate a GitHub Personal Access Token (PAT) using the `gh` (GitHub CLI) tool:

1.  **Install GitHub CLI:** Follow the instructions on the official GitHub CLI documentation: [https://cli.github.com/](https://cli.github.com/)
2.  **Authenticate with GitHub:**
    ```bash
    gh auth login
    ```
    Follow the prompts to authenticate your GitHub account.
3.  **Retrieve your token:**
    ```bash
    gh auth token
    ```
    This command will display your current GitHub token. **Copy this token carefully and keep it secure. Do not commit it to your repository.**

## 2. Triggering CI/CD Builds

Once you have tested your generated configurations locally (e.g., using `docker compose up -d --build`), you can trigger the CI/CD pipeline. The `orchestrator-cli` generates GitHub Actions workflows that are typically triggered by:

*   **Pushing to `main` or `develop` branches:** Any push to these branches will automatically start the CI/CD process.
*   **Creating a new Git tag:** Creating and pushing a new Git tag (e.g., `git tag v1.0.0` and `git push origin v1.0.0`) will also trigger a release build, often used for Docker image pushes to Docker Hub.

## 3. Setting Up Terraform and Kubernetes

The `orchestrator-cli` generates *template* files for Terraform and Kubernetes. These templates provide a starting point for your infrastructure and deployment, but require further configuration and application.

### Terraform (Infrastructure as Code)

Terraform is an open-source Infrastructure as Code (IaC) tool that allows you to define and provision infrastructure using a declarative configuration language.

**Generated File:** `terraform/main.tf` (or similar, depending on the template)

**Setup Steps:**

1.  **Install Terraform CLI:** Follow the official Terraform documentation: [https://learn.hashicorp.com/terraform/getting-started/install](https://learn.hashicorp.com/terraform/getting-started/install)
2.  **Configure your Cloud Provider:** Ensure you have the necessary credentials and configurations for your chosen cloud provider (e.g., AWS, Azure, GCP) set up in your environment.
3.  **Initialize Terraform:** Navigate to the directory containing your `main.tf` file and run:
    ```bash
    terraform init
    ```
    This command initializes a working directory containing Terraform configuration files.
4.  **Review the Plan:**
    ```bash
    terraform plan
    ```
    This command creates an execution plan, which shows you what Terraform will do to achieve the desired state. **Always review the plan carefully before applying.**
5.  **Apply the Configuration:**
    ```bash
    terraform apply
    ```
    This command executes the actions proposed in a Terraform plan to create, update, or destroy infrastructure.

### Kubernetes (Container Orchestration)

Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications.

**Generated File:** `kubernetes/deployment.yml` (or similar, depending on the template)

**Setup Steps:**

1.  **Install `kubectl`:** The Kubernetes command-line tool. Follow the official Kubernetes documentation: [https://kubernetes.io/docs/tasks/tools/install-kubectl/](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
2.  **Configure Kubernetes Context:** Ensure your `kubectl` is configured to connect to your Kubernetes cluster. This usually involves setting up your `kubeconfig` file.
3.  **Apply the Deployment:**
    ```bash
    kubectl apply -f kubernetes/deployment.yml
    ```
    This command applies the Kubernetes deployment configuration to your cluster, creating or updating your application's pods and other resources.
