# CoverCraft

CoverCraft is a project that generates cover letters from a job description and a resume using an AI model powered by ChatGPT-3. This project is built using Go (Golang), PostgreSQL, and the Echo web framework for creating a web-based application with simple HTML forms.

## Table of Contents

- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before you can run the CoverCraft project, you need to have the following prerequisites installed:

- Go (Golang): Install it from [the official Go website](https://golang.org/dl/).

- PostgreSQL: Install PostgreSQL and create a database for the project. Update the database connection configuration in `config/db.go`.

### Installation

1. Clone the CoverCraft repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/CoverCraft.git
   cd CoverCraft
   ```
Install project dependencies using the provided Makefile:

```bash
Copy code
make install
```
### Usage
Start the application using the Makefile:

```bash
Copy code
make run
```
Alternatively, you can start the application using the following command:

```bash
Copy code
go run app/main.go
```
Access the CoverCraft application by navigating to http://localhost:your_port in your web browser.

Use the simple HTML forms to input a job description and resume.

Submit the form to generate a cover letter using the AI model powered by ChatGPT-3.

### Project Structure
The project is organized into the following directories:

- ```app/main.go```: Contains the entry point of the application.

- ```config/db.go```: Manages the database connection and configuration.

- ```domain/repositories```: Contains the data access layer.

- ```domain/aggregates```: Houses business models with access roots.

- ```domain/interfaces/http```: Contains HTTP interfaces and handlers for various components, such as job handling and cover letter optimization.

- ```domain/services```: Contains the service layer to interact with ChatGPT-3 for generating cover letters.

- ```domain/valueobjects```: Contains request value objects used in the application.

- ``` entities``` : Contains database entities, such as the Job entity.

- ```templates``` : Contains HTML templates for rendering web pages.

### Contributing
Contributions to CoverCraft are welcome! If you have ideas for improvements, new features, or bug fixes, please create a pull request. For major changes, please open an issue first to discuss your ideas.

### License
This project is licensed under the MIT License.