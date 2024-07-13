# 🦆 QuackBin

QuackBin is a minimalist, pastebin service that prioritizes privacy and simplicity. Built with Go and powered by DuckDB, QuackBin offers a lightweight, single-binary solution for sharing encrypted text snippets.

## Features

- **Client-Side Encryption**: All encryption and decryption occur in the browser, ensuring zero-knowledge on the server's side
- **Private Pastes**: Option to create password-protected pastes for added security
- **IP Rate Limiting**: Prevents abuse through a mutex-based rate limiting system

![QuackBin Preview](https://github.com/GiorgioBaldelli/quackbin/blob/main/preview.png)

## Tech Stack

- **Backend**: Go
- **Database**: DuckDB
- **Frontend**: HTML, CSS, JavaScript
- **Encryption**: CryptoJS library (AES-256 in CBC mode)

## Getting Started

To get QuackBin up and running on your local machine, follow these simple steps:

1. Clone this repository: `git clone https://github.com/GiorgioBaldelli/quackbin.git`

2. Navigate to the project directory:
`cd quackbin`

3. Build: `docker build -t quackbin .`

4. Run the server: `docker run -p 8080:8080 quackbin`

5. Navigate to `http://localhost:8080`

## Contributing

If you have ideas for improvements, feel free to:

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
