# Video Compress 

A simple web app to convert videos from H.264 to H.265 encoding, significantly reducing file size while maintaining quality.

## Why H.265?

H.265 (HEVC) is the successor to H.264 (AVC). It offers better compression, allowing for smaller file sizes or higher quality at the same bitrate. This project uses FFmpeg to convert videos from H.264 to H.265.

## Quick Start
Run in both servers in different windows

### Backend

```bash
cd video-compress/backend
docker build -t video-compress .
docker run -p 5000:5000 video-compress
```

### Frontend


```bash
# Install Bun (for macOS, Linux, and WSL)
curl -fsSL https://bun.sh/install | bash

cd video-compress/frontend
bun install
bun run build
bun run preview
```

## Usage

1. Start the backend and frontend.
2. Open the web app in your browser.
3. Upload an H.264 video.
4. Download the converted H.265 video.

Enjoy!
