# **📡 RSS Feed Aggregator**

Welcome to the RSS Feed Aggregator project! This is a robust Go-based web application designed to fetch and display content from various RSS and Atom feeds. I built this to provide a clean, dynamic, and easy-to-use interface for keeping up with news and articles from multiple sources, all in one place.

With support for concurrent fetching and a snappy, modern frontend powered by HTMX, this aggregator offers a seamless reading experience. Whether you're a developer, a news enthusiast, or just looking for a better way to consume your favorite content, this project aims to simplify your daily information intake. 🚀

## Installation

Getting the RSS Feed Aggregator up and running on your local machine is straightforward. Follow these steps:

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/samueltuoyo15/Rss-Feed-Aggregator.git
    cd Rss-Feed-Aggregator
    ```

2.  **Build the Application**:
    This command will compile the Go source code and create an executable binary named `main` in your project root.
    ```bash
    go build -o main ./cmd/main.go
    ```

3.  **Run the Application**:
    Execute the compiled binary. The server will start and listen on port `8080`.
    ```bash
    ./main
    ```

4.  **Access in Browser**:
    Open your web browser and navigate to `http://localhost:8080`. You should see the aggregator fetching and displaying feeds.

### 🐳 Running with Docker

For a containerized setup, you can use the provided `Dockerfile`:

1.  **Build the Docker Image**:
    ```bash
    docker build -t rss-aggregator .
    ```

2.  **Run the Docker Container**:
    ```bash
    docker run -p 8080:8080 rss-aggregator
    ```
    The application will now be accessible at `http://localhost:8080`.

## Usage

Once the server is running, navigating to `http://localhost:8080` in your web browser will load the main interface.

The application dynamically fetches and displays the latest articles from the configured RSS feeds. Here's what you can expect:

*   **Homepage (`/`)**: This serves the `index.html` page, which includes HTMX to automatically trigger a request to `/api/feeds` upon loading. This ensures that the latest feed items are immediately displayed without a full page refresh.
*   **Dynamic Content Loading**: As you interact with the page, or when it first loads, HTMX will make an `HX-Request` to the server. The server, recognizing this, will render and return only the `feed_items.html` partial, which HTMX then seamlessly inserts into the page, providing a fast and responsive user experience.
*   **JSON API (`/api/feeds`)**: You can also access the aggregated feed data directly as JSON by visiting `http://localhost:8080/api/feeds`. This endpoint provides all the fetched feed items in a structured format, which can be useful for other applications or debugging.
*   **Configuring Feeds**: The list of RSS feeds is loaded from `config/feeds.yaml`. To add or remove feeds, simply edit this file. The application will fetch from all URLs listed there.
*   **Refreshing Content**: The current setup fetches feeds once on page load or when triggered by HTMX. For continuous updates, you might consider implementing a refresh mechanism (e.g., a refresh button or a periodic polling using HTMX).

## Features

This RSS Feed Aggregator comes packed with features designed for efficiency and ease of use:

*   **Concurrent Feed Fetching**: Utilizes Go's concurrency primitives (goroutines and wait groups) to fetch multiple RSS feeds simultaneously, ensuring quick loading times.
*   **YAML-based Configuration**: Easily configure your desired RSS feeds by modifying the `config/feeds.yaml` file, making it simple to add or remove sources.
*   **Dynamic Frontend with HTMX**: Leverages HTMX for light-weight, dynamic content updates without writing complex JavaScript, resulting in a highly responsive user interface.
*   **Clean HTML Templating**: Uses Go's `html/template` package to render server-side HTML, integrating custom functions for date formatting and content truncation.
*   **Docker Support**: Includes a `Dockerfile` for easy containerization, enabling consistent deployment across different environments.
*   **Structured Data Models**: Defines clear Go structs (`Feed` and `FeedItem`) to represent feed configurations and parsed articles, promoting clean and maintainable code.
*   **Robust Error Handling**: Incorporates error checking during feed loading and fetching to provide informative messages and maintain application stability.
*   **Flexible RSS Parsing**: Employs the `gofeed` library to reliably parse various RSS and Atom feed formats, extracting essential details like title, link, published date, description, categories, and author.
*   **Thumbnail Extraction**: Automatically attempts to find and display relevant image thumbnails for feed items from various sources within the feed data.

## Technologies Used

This project is built using a modern stack, combining the power of Go on the backend with a lightweight, dynamic frontend approach.

| Technology      | Description                                                    | Links                                                                 |
| :-------------- | :------------------------------------------------------------- | :-------------------------------------------------------------------- |
| **Go**          | Primary language for the backend, known for performance and concurrency. | [Go Official Website](https://golang.org/)                            |
| **HTMX**        | A lightweight JavaScript library for dynamic HTML updates without heavy JS. | [HTMX Official Website](https://htmx.org/)                            |
| **Tailwind CSS**| A utility-first CSS framework for rapidly building custom designs. | [Tailwind CSS Official Website](https://tailwindcss.com/)             |
| **GoFeed**      | A Go library for parsing RSS and Atom feeds.                   | [GoFeed GitHub](https://github.com/mmcdole/gofeed)                    |
| **YAML**        | Used for simple and human-readable configuration of RSS feeds. | [YAML Official Website](https://yaml.org/)                            |
| **Docker**      | For containerizing the application, ensuring consistent environments. | [Docker Official Website](https://www.docker.com/)                    |
| **Render**      | Cloud platform for deploying web services (configured via `render.yaml`). | [Render Official Website](https://render.com/)                        |

## Contributing

I'd be thrilled if you'd like to contribute to the RSS Feed Aggregator! Whether it's a bug fix, a new feature, or an improvement to the documentation, your input is highly valued.

👉 **Here's how you can get started:**

*   **Fork the repository**.
*   **Create a new branch** for your feature or bug fix: `git checkout -b feature/your-feature-name` or `git checkout -b bugfix/issue-description`.
*   **Make your changes**, ensuring they align with the project's coding style and conventions.
*   **Write clear, concise commit messages**.
*   **Push your branch** to your forked repository.
*   **Open a Pull Request** to the `main` branch of this repository, describing your changes in detail.

I appreciate your efforts in making this project even better! ✨

## License

This project is licensed under the MIT License.

## Author Info

👋 Hi, I'm Samuel Tuoyo, the creator of this RSS Feed Aggregator!

I'm passionate about building robust and efficient web applications using Go and exploring modern frontend patterns. Connect with me on:

*   **GitHub**: [@samueltuoyo15](https://github.com/samueltuoyo15)
*   **LinkedIn**: [Your LinkedIn Profile](https://linkedin.com/in/your-profile)
*   **Portfolio**: [Your Portfolio Website](https://your-portfolio.com)

---

[![Go Version](https://img.shields.io/github/go-mod/go-version/samueltuoyo15/Rss-Feed-Aggregator?style=flat-square&label=Go)](https://golang.org/)
[![Powered by HTMX](https://img.shields.io/badge/Powered%20by-HTMX-blueviolet?style=flat-square)](https://htmx.org/)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)