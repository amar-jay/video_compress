import { useState, useEffect } from "react";

export const useVideoThumbnail = (videoFile: Blob, seekTo = 1) => {
  const [thumbnail, setThumbnail] = useState("");

  useEffect(() => {
    if (!videoFile) {
      setThumbnail("");
      return;
    }

    const generateThumbnail = async () => {
      try {
        const videoUrl = URL.createObjectURL(videoFile);
        const video = document.createElement("video");

        const seekPromise = new Promise((resolve, reject) => {
          video.onloadedmetadata = () => {
            if (video.duration < seekTo) {
              reject(new Error("Video duration is less than seekTo time"));
            } else {
              video.currentTime = seekTo;
            }
          };
          video.onseeked = resolve;
          video.onerror = reject;
        });

        video.src = videoUrl;
        await seekPromise;

        const canvas = document.createElement("canvas");
        canvas.width = video.videoWidth;
        canvas.height = video.videoHeight;
        const ctx = canvas.getContext("2d");
        if (ctx == null) {
          alert("unknown context");
          return;
        }

        ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

        const thumbnailUrl = canvas.toDataURL();
        setThumbnail(thumbnailUrl);

        URL.revokeObjectURL(videoUrl);
      } catch (err: any) {
        alert(err.message);
      }
    };

    generateThumbnail();
  }, [videoFile, seekTo]);

  return { thumbnail };
};
