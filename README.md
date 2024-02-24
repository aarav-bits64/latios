# Latios 🚀

Latios is a versatile and efficient file downloader written in Go, designed to simplify the process of downloading files with a user-friendly interface and progress tracking.

## Features 🌟

- Seamless file downloads from any URL.
- Real-time progress bar for tracking download status.
- Automatic detection of filenames from URLs.
- Estimated time remaining for download completion.

## Getting Started 🚗

1. Clone the Latios repository:

   ```bash
   git clone https://github.com/username/latios.git
   ```

2. Install dependencies:

   ```bash
   go get -u github.com/cheggaaa/pb/v3
   ```

3. Build the Latios executable:

    Using Makefile:
    ```bash
    make build
    ```

    You can also install latios to bin directory
    so that you can use it from any directory by
    entering the following command:
    ```bash
    make install 
    ```

    Or manually by:
   ```bash
   go build -o latios main.go
   ```

4. Run Latios with your desired URL:

   ```bash
   ./latios download https://example.com/samplefile.zip
   ```

## Usage Example 📝

```bash
./latios download https://example.com/samplefile.zip
```

## Author 🧑‍💻

- Aarav Shreshth

## Contribution 🤝

Feel free to contribute by submitting bug reports, feature requests, or pull requests. Let's make Latios even better together!

## License 📜

This project is licensed under the [MIT License](LICENSE).

---

Thank you for using Latios! Happy downloading! 🌐💾