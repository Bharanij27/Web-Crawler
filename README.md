# Web Crawler

A simple web crawler application build using Golang. This CLI based application is suitable to fetch the detail from site and get metadata from it's child sites


# Local Development

### Prerequisites
Before you begin, ensure you have the following installed:

- **Go (1.18 or later)**: [Download Go](https://golang.org/dl/)


### Steps

1. Navigate to the project directory
   ```shell
   git clone https://github.com/BharaniJ27/Web-Crawler.git
   cd Web-Crawler
   ```
2. Install all the dependencies
   ```shell
   go build -o web-crawler
   ```
3. Build the application
    ```shell
        go mod tidy
    ```

# Testing
```shell
web-crawler

Hello, Welcome to web crawler
Please enter a valid URL

https://theuselessweb.com/

Please enter the depth to crawl:
3
```

Sample response
```shell
URL: https://theuselessweb.com/ | Title: The Useless Web | Content: The Useless Web TAKE ME TO A USELESS WEBSITE → PLEASE ← Read About The Sites The Useless Web... 

URL: https://theuselessweb.com//sites | Title: The Useless Web Sites | Content: The Useless Web Sites The Useless Sites of the Useless Web → Discover More Sites Checkbox Race By 

URL: http://tholman.com | Title: Tim Holman - Engineer &  Maker | Content: Tim Holman - Engineer &  Maker Tim Holman Tim Holman I am a veteran front-end engineer who thrives o
```