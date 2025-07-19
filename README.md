# **RSS Feed Aggregator: Your Curated News Hub**

Dive into a seamless news experience with the **RSS Feed Aggregator**! This project is a lean, mean, and efficient Go application designed to fetch and display RSS/Atom feeds from various sources in one convenient place. Built with performance in mind and featuring a clean web interface powered by HTMX, it's perfect for staying on top of your favorite content without the clutter. ✨

## Installation

Getting this aggregator up and running on your local machine is straightforward. You have a couple of options: directly using Go or leveraging Docker for containerization.

### Prerequisites

*   **Go**: Version 1.24.2 or newer
*   **Git**: For cloning the repository
*   **Docker (Optional)**: If you prefer to run it in a container

### Clone the Repository

First things first, grab the code:

```bash
git clone https://github.com/samueltuoyo15/Rss-Feed-Aggregator.git
cd Rss-Feed-Aggregator
```

### Option 1: Running with Go

Follow these steps to run the application directly using your Go installation:

1.  📦 **Download Dependencies**:
    ```bash
    go mod download
    ```
2.  🛠️ **Build the Application**:
    ```bash
    go build -o main ./cmd/main.go
    ```
3.  🚀 **Run the Application**:
    ```bash
    ./main
    ```

The server will start on `http://localhost:8080`.

### Option 2: Running with Docker

For a more isolated and portable setup, Docker is your friend:

1.  🐳 **Build the Docker Image**:
    ```bash
    docker build -t rss-feed-aggregator .
    ```
2.  🚀 **Run the Docker Container**:
    ```bash
    docker run -p 8080:8080 rss-feed-aggregator
    ```

The application will now be accessible via `http://localhost:8080` in your browser.

## Usage

Once the server is running (either via Go or Docker), you can interact with the RSS Feed Aggregator in a couple of ways:

### Web Interface

Open your web browser and navigate to `http://localhost:8080`. You'll be greeted by the main page, which automatically fetches and displays the latest articles from the configured RSS feeds. The UI is designed to be responsive, providing a clean reading experience across different devices.

The feed items are loaded using HTMX, providing a snappy, dynamic feel without full page reloads.

### Configuration

The RSS feeds are configured via a YAML file located at `config/feeds.yaml`. You can easily add or remove feeds by editing this file.

Here's an example of the configuration structure:

```yaml
feeds:
  - name: "TechCrunch"
    url: "https://techcrunch.com/feed/"
  - name: "BBC News"
    url: "http://feeds.bbci.co.uk/news/rss.xml"
  - name: "BBC Sport - Football"
    url: "http://feeds.bbci.co.uk/sport/football/rss.xml"
```

After modifying `config/feeds.yaml`, you'll need to restart the application for the changes to take effect.

### API Endpoint

The aggregator also exposes a simple API endpoint for programmatically accessing the aggregated feeds:

*   **GET `/api/feeds`**: Returns a JSON array of all fetched feeds and their respective items. This is what the web interface uses under the hood for initial data and subsequent updates.

Example cURL request:
```bash
curl http://localhost:8080/api/feeds
```

## Features

*   📰 **Multi-Source Feed Aggregation**: Consolidates articles from multiple RSS/Atom feeds into a single, unified view.
*   🚀 **Efficient Fetching**: Utilizes Go's concurrency features to fetch feeds in parallel, ensuring quick updates.
*   🖼️ **Intelligent Thumbnail Extraction**: Beyond standard RSS enclosures, it includes custom logic to extract featured images from linked articles (e.g., TechCrunch).
*   ⚙️ **YAML-based Configuration**: Easily add, remove, or modify feed sources by editing a simple `feeds.yaml` file.
*   🌐 **Modern Web Interface**: A clean and responsive user interface built with HTML templates, HTMX for dynamic content loading, and Tailwind CSS for styling.
*   💡 **Custom Template Functions**: Includes handy functions for date formatting and content truncation to present information neatly.
*   📦 **Docker Support**: Provided `Dockerfile` for easy containerization and deployment.
*   📡 **RESTful API**: Exposes a `/api/feeds` endpoint for consuming the aggregated data programmatically.

## Technologies Used

| Category   | Technology     | Description                                                                     |
| :--------- | :------------- | :------------------------------------------------------------------------------ |
| **Backend**  | [Go](https://go.dev/)          | The primary language for the server-side logic, chosen for its concurrency and performance. |
| **Libraries**| [gofeed](https://github.com/mmcdole/gofeed) | Robust RSS/Atom feed parsing library.                                           |
|            | [goquery](https://github.com/PuerkitoBio/goquery) | A Go library that brings jQuery-like syntax to Go for HTML parsing (used for thumbnail extraction). |
|            | [yaml.v3](https://github.com/go-yaml/yaml) | Handles YAML file parsing for feed configuration.                               |
| **Frontend** | [HTMX](https://htmx.org/)      | A lightweight JavaScript library for dynamic UI updates via HTML attributes.    |
|            | [Tailwind CSS](https://tailwindcss.com/) | A utility-first CSS framework for rapid and consistent styling.                 |
| **Containerization**| [Docker](https://www.docker.com/) | For building and running the application in isolated containers.              |

## License

This project is open-source and released under an [MIT License](https://opensource.org/licenses/MIT).

## Author

**Samuel Tuoyo**

A passionate developer with a keen interest in building efficient and scalable applications.

* [X](https://x.com/TuoyoS26091)

---

[![Go Version](https://img.shields.io/github/go-mod/go-version/samueltuoyo15/Rss-Feed-Aggregator?style=flat-square&color=00ADD8)](https://go.dev/)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/samueltuoyo15/rss-feed-aggregator/latest?style=flat-square&label=Docker%20Image&color=0db7ed)](https://hub.docker.com/r/samueltuoyo15/rss-feed-aggregator)
[![Repository Stars](https://img.shields.io/github/stars/samueltuoyo15/Rss-Feed-Aggregator?style=flat-square&color=blue)](https://github.com/samueltuoyo15/Rss-Feed-Aggregator/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
