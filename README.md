# ms-mpeg-master-stream

ms-mpeg-master-stream is a personal Go project showcasing a robust video encoder that efficiently converts MP4 files into the versatile MPEG-DASH format. This project emphasizes simplicity, efficiency, and adaptability, making it an ideal choice for seamless integration into various applications.

## Key Features

- **MP4 to MPEG-DASH Conversion:** Easily transform MP4 videos into MPEG-DASH for adaptive streaming.
- **Queue-Based Decoding:** Efficiently handle video conversion requests using a reliable queue system.
- **Web Interface:** User-friendly interface for uploading and initiating the conversion process.
- **Adaptive Streaming:** Enhance the viewing experience with adaptive bitrate streaming using MPEG-DASH manifests.
- **Dockerized Deployment:** Easily deploy and scale the encoder using Docker containers for flexibility.

## Usage

1. **Installation:** Clone the repository and install the necessary dependencies.

   ```bash
   git clone https://github.com/your-username/ms-mpeg-master-stream.git
   cd ms-mpeg-master-stream
   go install
   ```

2. **Run the Encoder:**

   ```bash
   ms-mpeg-master-stream
   ```

3. **Web Interface:**

   - Access the web interface at [http://localhost:8080](http://localhost:8080).
   - Upload your MP4 files and initiate the conversion process.

4. **Queue-Based Decoding:**

   - View the efficient handling of video conversion requests through the queue system.

5. **Adaptive Streaming:**
   - Retrieve the converted MPEG-DASH files for seamless adaptive streaming.

## Dockerized Deployment

To deploy using Docker, use the following steps:

1. **Build Docker Image:**

   ```bash
   docker build -t ms-mpeg-master-stream .
   ```

2. **Run Docker Container:**

   ```bash
   docker run -p 8080:8080 ms-mpeg-master-stream
   ```

## Contributing

Feel free to contribute by opening issues, providing suggestions, or submitting pull requests. Your contributions are highly valued!

## License

This project is licensed under the [MIT License](LICENSE.md).
