# Backend

Functionality:

-   web interface:
    -   CRUD expense-types
    -   CRUD expense
-   telegram bot
    -   CREATE expense (write a message like 1000 and choose the type of expense)
    -   Send token on any message /*

Structure:

-   core
    -   db (sqlite)
    -   services (buisness logic)
-   web interface (templates + htmx + tailwindcss)
    - web-server
-   telegram-bot interface
    - bot

# How to develop

```bash
cp .env.example .env
```

fill in the .env file with your own values

```bash
docker compose up
```

# How to run in production

Basicly, you just need to upload you container to some ecr, set envs and run docker-compose.yml

## Automations

You might need to change compose and scripts files to fit your needs, but here is a general idea of how to run it in production.

```bash
./scripts/build-and-push-image.bash
```

```bash
cp .env.example .env
```

add your ssh key to your server

fill in the .env file with your own values

copy it to your vps

and run

```bash
./scripts/deploy.bash
```

This script will copy everything needed to vps and run the app with docker compose
It also sets up a backup cron job for db

You'll need to manage ssl and ingress on your own, because in my case, it's configured outside of this projects scope.
