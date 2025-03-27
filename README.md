# Employee Management CRUD

This project demonstrates how to implement **CRUD** in a **Go** application with **Fiber**, **GORM**, and **MongoDB** for **Employee Management**. It is based on a Clean Architecture example from the GoFiber documentation, modified to use the **MongoDB Go v2 driver** and the latest versions of **GoFiber** and related libraries.

## Description

This project serves as a learning resource for applying **Clean Architecture** principles in Go. The goal is to structure an application in a way that separates concerns into different layers, making it easier to maintain, scale, and test.

### Key Technologies:

- **GoFiber**: Web framework for building RESTful APIs.
- **GORM**: ORM for Go to interact with MongoDB.
- **MongoDB**: NoSQL database to store employee data.

### Changes Made:

- Replaced the old MongoDB driver with the new v2 MongoDB driver.
- Updated Fiber from version 1.x to 2.x.
- Changed the core entity from `Book` to `Employee`, reflecting a real-world scenario of managing employee data.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [MongoDB](https://www.mongodb.com/try/download/community) v4.x or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `api/`: Contains the HTTP handlers, routes, and presenters for the application.
- `pkg/`: Contains the core business logic, including services and entities.
- `cmd/`: Contains the main application entry point.

## Conclusion

This project demonstrates how to build a simple **GoFiber** application following **Clean Architecture** principles. It can be extended for more complex applications and is a good starting point for anyone interested in structuring their Go projects in a clean, maintainable way.

## References

- [GoFiber Documentation](https://docs.gofiber.io)
- [MongoDB Documentation](https://docs.mongodb.com/)
