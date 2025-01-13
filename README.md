## Table of Contents

- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Development Setup](#development-setup)
- [Development Mode](#development-mode)
- [Production Mode](#production-mode)
- [Access Information](#access-information)
- [Login Information](#login-information)

## Overview

This project is aimed at learning application development using Go for the backend and Next.js for the frontend.  
It is designed as a web service where users can buy and sell image materials.

## Tech Stack

### Backend

- Language: Golang
- Framework: Echo
- ORM: GORM
- DB: MySQL

### Frontend

- Language: TypeScript
- Framework: Next.js
- CSS Framework: ChakraUI

## Development Setup

### Create `.env` file for Backend

```sh
cp backend/app/.env.local.example backend/app/.env.local
```

### Create `.env` file for Frontend

```sh
cp frontend/.env.local.example frontend/.env.local
```

## Development Mode

This mode is used for local development.

### Build

```sh
docker compose build
```

### Start

#### Initialize and start with sample data

```sh
SEEDING=true docker compose up -d
```

#### Normal start

```sh
docker compose up -d
```

### Stop

```sh
docker compose down
```

## Production Mode

This mode is used for production environments. It can also be used to verify operations in an environment equivalent to production during development.

### Build

```sh
docker compose -f docker-compose.yml build
```

### Start

```sh
docker compose -f docker-compose.yml up -d
```

### Stop

```sh
docker compose down
```

## Access Information

In the local environment, you can access it at http://trill.localhost.

## Login Information

The following sample users are provided for testing. Sample users will be created when starting with `SEEDING=true docker compose up -d`.

| Email Address | Password |
| ---- | ---- |
| bob@example.com | uHrdx55u | 
| john@example.com | uHrdx55u | 
| alice@example.com | uHrdx55u | 
| emma@example.com | uHrdx55u | 
