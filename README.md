# Introduction

This is a template project with hot reload using fresh package.
source : https://medium.com/@McMenemy/godorp-docker-compose-for-development-and-production-e37fe0a58d61

# Commands

Create .env file with 3 variables (APP_ENV, POSTGRES_USER, POSTGRES_PASSWORD)

- Launch project in dev mode :

1. Delete APP_ENV from .env file.
2. `docker-compose up --build`

- Launch project in production mode :

1. Add APP_ENV=production to .env file.
2. `docker-compose up --build`
