# Telegram Bot Starter Template for Gotd

Welcome to the Telegram Bot Starter Template based on the `gotd` library! This template provides you with a solid foundation to build your own Telegram bot using Golang. It includes essential features, integration with PostgreSQL via Prisma, language localization, and encapsulation techniques for maintainable code.

## Table of Contents

- [Features](#features)
- [Installation and Local Launch](#installation-and-local-launch)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Features

- Common middlewares for handling various aspects of bot functionality.
- Level-based access control for bot commands.
- Integration with PostgreSQL using Prisma for efficient database operations.
- Language picker and internationalization support for user-friendly interactions.
- Basic encapsulation techniques to enhance code organization and readability.

## Installation and Local Launch

Follow these steps to get the bot up and running locally:

1. Clone this repository: `git clone https://github.com/jaskaurhello/telegram-bot-gotd-template.git`
2. Launch a local instance of the [PostgreSQL database](https://postgresapp.com).
3. Start a local Redis server.
4. Create a `.env` file with the required environment variables (see [Environment Variables](#environment-variables)).
5. Run `go get -u .` in the root folder of the project.
6. Install Prisma by running: `go get github.com/steebchen/prisma-client-go`
7. Push database schema changes with: `go run github.com/steebchen/prisma-client-go db push`
8. Start the bot with: `make dev`

Once these steps are completed, your bot should be up and running locally, ready for development and testing.

## Environment Variables

Make sure to set the following environment variables in your `.env` file:

- `BOT_TOKEN`: Your Telegram bot token.
- `DATABASE_URL`: URL of the PostgreSQL database.
- `REDIS_URL`: URL of the Redis database.
- `SUDO_ADMINS`: IDs of your bot's admin users.

Don't forget to refer to the provided `.env.sample` for guidance on setting up these variables.

## Contributing

Contributions to this template are welcome! If you find any issues or want to enhance the template's features, feel free to fork this repository and submit pull requests. Let's make building Telegram bots in Golang even better together.

## License

This project is licensed under the MIT License. You're free to use, modify, and distribute the code for any purpose. While not required, it would be appreciated if you acknowledge the original developers by leaving a note in your project. Thank you for using this starter template!

Happy coding! 🚀
