# Go IAM E-Commerce Microservice

This project is a Go-based IAM (Identity and Access Management) E-commerce microservice that provides a complete and scalable e-commerce solution, incorporating user management, product management, shopping cart management, logistics, finance, advertisements, and inventory management. The microservice is also designed to integrate with Amazon's APIs for further expansion and capabilities.

## Features

1. **Identity and Access Management (IAM)**: Securely manage user authentication and authorization, ensuring proper access control to various e-commerce functionalities.
2. **User Management**: Add, update, delete, and retrieve user information, allowing easy administration of customer accounts.
3. **Item Management**: Add, update, delete, and retrieve product information, allowing easy management of the product catalog.
4. **Shopping Cart Management**: Enable users to add, remove, and update items in their shopping cart, providing a seamless shopping experience.
5. **Logistics**: Manage and track shipping and delivery information, ensuring efficient fulfillment of customer orders.
6. **Finance**: Handle payments, refunds, and other financial transactions, providing a secure and streamlined checkout process.
7. **Advertisements**: Showcase featured products, promotions, and discounts to drive sales and customer engagement.
8. **Inventory Management**: Monitor and manage stock levels in real-time, preventing stock-outs and ensuring product availability.
9.  **Amazon API Integration**: Leverage the power of Amazon's APIs to expand the platform's capabilities, such as utilizing Amazon's fulfillment services or product recommendations.

## Getting Started

To get started with the Go IAM E-Commerce Microservice, follow these steps:

1. Clone the repository to your local machine.

```bash
cd {your-project-path}/go-iam-ecommerce-microservice/cmd/apiserver
go build
./apiserver -c {config_path}
```

### Example:

```bash
(base) huanghaitao@huanghaitaodeMacBook-Pro apiserver % go build                                       
(base) huanghaitao@huanghaitaodeMacBook-Pro apiserver % ./apiserver -c ../../configs/iam-apiserver.yaml
2023-05-17 14:11:44.570 INFO    app/app.go:293  ==> WorkingDir: /Users/huanghaitao/go/src/go-iam-ecommerce-microservice/cmd/apiserver
2023-05-17 14:11:44.574 INFO    app/app.go:252  ==> Starting IAM API Server ...
2023-05-17 14:11:44.574 INFO    app/app.go:254  ==> Version: `{"gitVersion":"v0.0.0-master+$Format:%h$","gitCommit":"$Format:%H$","gitTreeState":"","buildDate":"1970-01-01T00:00:00Z","goVersion":"go1.19.1","compiler":"gc","platform":"darwin/amd64"}`
2023-05-17 14:11:44.574 INFO    app/app.go:257  ==> Config file used: `../../configs/iam-apiserver.yaml`
2023-05-17 14:11:44.574 INFO    app/app.go:285  ==> Config: `{"server":{"mode":"debug","healthz":true,"middlewares":["recovery","logger","secure","nocache","cors","dump"]},"grpc":{"bind-address":"127.0.0.1","bind-port":8081,"max-msg-size":4194304},"insecure":{"bind-address":"127.0.0.1","bind-port":8883},"secure":{"bind-address":"127.0.0.1","bind-port":8443,"Required":true,"tls":{"cert-key":{"cert-file":"/Users/huanghaitao/iam/etc/cert/iam-apiserver.pem","private-key-file":"/Users/huanghaitao/iam/etc/cert/iam-apiserver-key.pem"},"cert-dir":"/var/run/iam","pair-name":"iam"}},"mysql":{"host":"127.0.0.1:3306","username":"root","database":"iam","max-idle-connections":100,"max-open-connections":100,"max-connection-life-time":10000000000,"log-level":4},"redis":{"host":"127.0.0.1","port":6379,"addrs":[],"username":"","password":"MyN3wP4ssw0rd","database":0,"master-name":"","optimisation-max-idle":2000,"optimisation-max-active":4000,"timeout":0,"enable-cluster":false,"use-ssl":false,"ssl-insecure-skip-verify":false},"jwt":{"realm":"JWT","key":"dfVpOK8LZeJLZHYmHdb1VdyRrACKpqoo","timeout":86400000000000,"max-refresh":86400000000000},"log":{"output-paths":["/Users/huanghaitao/iam/log/iam-apiserver.log","stdout"],"error-output-paths":["/Users/huanghaitao/iam/log//iam-apiserver.error.log"],"level":"debug","format":"console","disable-caller":false,"disable-stacktrace":false,"enable-color":true,"development":true,"name":"apiserver"},"feature":{"profiling":true,"enable-metrics":true}}`
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 ...
```

# Item Management API Documentation

## Overview

The Item API is designed to provide an interface for managing inventory items in an e-commerce system. It is structured around a three-tier architecture: Controller, Service, and Store. These layers work together to provide a robust and flexible solution for managing your inventory items.

## Controller Layer

The Controller layer is the entry point for the API, responsible for handling HTTP requests and responses. It validates input, calls the appropriate Service functions based on the HTTP method (GET, POST, PUT, DELETE), and formats the response to be sent back to the client.

## Service Layer

The Service layer provides the business logic for our API. It contains the core functions for creating, reading, updating, and deleting (CRUD) items. The Service layer handles transactions, ensures data integrity, and calls the appropriate Store methods to interact with the database.

## Store Layer

The Store layer is responsible for data storage and retrieval. It communicates directly with the database to execute queries. This layer abstracts the database-specific details, making it possible to switch databases with minimal changes to the rest of the code.

## Field Search

The Item API supports advanced filtering using a field search feature. This allows users to filter results based on the value of specific fields in the database. The fields that can be searched are:

- `ASIN`
- `SKU`
- `Brand`
- `Title`
- `ProductGroup`
- `ProductType`

To use field search, include a `fieldSelector` parameter in your request with the format `field=value`. You can include multiple field-value pairs, separated by commas, to filter on multiple fields at once. For example, `fieldSelector=brand=Kuggini,title=headphones` would return items where the `brand` is `Kuggini` and the `title` contains `headphones`.

## Examples

Here's an example of a curl command to get items with a specific brand and title:

### Item list and advanced filter

```bash
curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer your_auth_token' "http://127.0.0.1:8883/v2/items?fieldSelector=brand%3DKuggini%2Ctitle%3Dheadphones&offset=0&limit=10"
```

In this example, `%3D` is the URL encoded equivalent of the `=` character, and `%2C` is the URL encoded equivalent of the `,` character. This is necessary because these characters have special meanings in URLs. The `fieldSelector` is `brand%3DKuggini%2Ctitle%3Dheadphones`, which is equivalent to `brand=Kuggini,title=headphones` when URL decoded.

### Create Item
```bash
curl -X POST -H'Content-Type: application/json' -H'Authorization: Bearer your_auth_token' -d '{"asin":"B0B3HXK001","title":"Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX6 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 000000001","brand":"Kuggini","product_group":"","product_type":""}' http://127.0.0.1:8883/v2/items
```

### Update Item
``` bash
curl -X PUT -H'Content-Type: application/json' -H'Authorization: Bearer your_auth_token' -d '{"asin":"B0B3HXK000","title":"Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX7 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 000-0001","sku":"LU-US-10000","brand":"Kuggini","product_group":"electronics","product_type":"headphones"}' http://127.0.0.1:8883/v2/items/10000
```

The Item API is a powerful tool for managing your e-commerce inventory. With its three-tier architecture and advanced field search feature, it provides a flexible and robust solution for all your inventory management needs.

# Documentation

Detailed documentation for each feature can be found in the /docs directory of this repository. This includes information on how to interact with each microservice, API endpoints, and sample code.

## To-Do List

Here are some features and improvements that we plan to implement in the future:

- [ ] Enhance user management with password recovery and email verification.
- [ ] Implement advanced product search and filtering capabilities.
- [ ] Add support for multiple currencies and languages.
- [ ] Improve logistics management with automated shipping label generation and tracking updates.
- [ ] Integrate third-party payment gateways for a more diverse set of payment options.
- [ ] Develop a customizable reporting dashboard to monitor sales and user activity.
- [ ] Incorporate machine learning algorithms for personalized product recommendations.
- [ ] Implement a fully responsive web-based admin panel for managing the e-commerce platform.
- [ ] Add APIs for third-party integrations and plugins.
- [ ] Improve performance and scalability with advanced caching and database optimization techniques.

Feel free to contribute to the project or suggest new features and improvements by creating an issue or submitting a pull request.

# Architecture

Below is an overview of the project architecture, describing the purpose of each directory:

- `CHANGELOG`: Contains a log of notable changes made to the project over time.
- `CONTRIBUTING.md`: Provides guidelines for contributors to the project.
- `LICENSE`: Contains the project's software license information.
- `Makefile`: A file containing build and other automation recipes for the project.
- `README.md`: Provides a high-level overview and documentation of the project.
- `api`: Contains API-related files, such as API specifications or generated client code.
- `build`: Stores build-related files, such as Dockerfiles, build scripts, or CI/CD pipeline configurations.
- `cmd`: Contains the main entry points for the project's executables, such as `apiserver` and `authzserver`.
- `configs`: Holds configuration files for various components of the project.
- `deployments`: Contains deployment-related files, such as Kubernetes manifests or Helm charts.
- `docs`: Stores project documentation, including guides, tutorials, or architectural diagrams.
- `githooks`: Contains Git hooks used by the project.
- `go.mod` and `go.sum`: Go module-related files that manage the project's dependencies.
- `init`: Stores initialization files or scripts for setting up the project environment.
- `internal`: Contains internal packages and components that are not meant to be used directly by external consumers.
    - `apiserver`: Houses the API server's core components, such as app, auth, config, controller, grpc, item, options, router, run, server, service, and store.
    - `authzserver`: Contains components related to the authorization server, such as analytics, app, authorization, config, controller, jwt, load, options, router, run, server, and store.
    - `iamctl`: Holds command-line interface (CLI) components for the IAM control tool.
    - `pkg`: Stores utility packages and shared components, such as code, logger, middleware, options, server, util, and validation.
    - `pump`: Contains components for the pump service, which is responsible for data processing and analytics.
    - `watcher`: Houses components for the watcher service, which monitors changes and updates to resources.
- `pkg`: Contains external packages and components that can be used by other projects or services.
    - `app`: Provides application-level components and utilities.
    - `cli`: Contains CLI-related packages and utilities.
    - `db`: Holds database-related components and utilities, such as MySQL and plugin.
    - `log`: Provides logging-related packages and utilities.
    - `shutdown`: Contains components for managing graceful shutdowns of applications and services.
    - `storage`: Houses storage-related components and utilities, such as Redis cluster and storage.
    - `util`: Contains various utility packages and components.
    - `validator`: Provides validation-related components and utilities.
- `test`: Contains test-related files, such as test data or test utilities.
- `third_party`: Stores third-party dependencies or libraries used by the project.
- `tools`: Contains tools or utilities used for development or management of the project.

This project structure follows best practices for organizing and managing a Go-based project, with separate directories for different components, utilities, and services, making it easy to navigate and understand.

# Contributing

Contributions to the Go IAM E-Commerce Microservice are welcomed and appreciated. Please follow these steps to contribute:

Fork the repository.

Create a new branch for your feature or bug fix.

Make your changes and commit them to the branch.

Create a pull request, detailing the changes made and the purpose of the changes.

# License

This project is licensed under the MIT License. See the LICENSE file for more information.

The Go IAM E-Commerce Microservice is designed to provide a comprehensive and scalable e-commerce solution, complete with essential features and Amazon API integration. By using this platform, you can create a user-friendly and efficient e-commerce experience for your customers.