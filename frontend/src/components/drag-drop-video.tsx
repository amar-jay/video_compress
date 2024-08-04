/**
 * v0 by Vercel.
 * @see https://v0.dev/t/v94jporfSux
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { useState, useRef } from "react";
import { CheckIcon } from "lucide-react";
import { Skeleton } from "./ui/skeleton";
import { downloadFile, formatFileSize } from "@/lib/utils";

enum UploadStatus {
  IDLE,
  LOADING,
  UPLOADED,
  COMPRESSED,
}

export default function Component() {
  const file = useRef<File>();
  const [uploadStatus, setUploadStatus] = useState<UploadStatus>(
    UploadStatus.IDLE,
  );
  const [output, setOutput] = useState<string>("");
  const [fileName, setFileName] = useState<string>("");
  const [thumbnail, setThumbnail] = useState<string>();
  const [times, setTimes] = useState<number>(); // set magnitude of reduction in size
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const _file = e.target?.files?.[0];
    if (_file && _file.type.startsWith("video/")) {
      file.current = _file;
      setFileName(_file.name);
      setUploadStatus(UploadStatus.UPLOADED);
    } else {
      alert("Please select a valid video file.");
    }
  };

  const handleCompress = async () => {
    if (!file.current) {
      alert("Please select a video file to upload");
      return;
    }

    if (!!file.current) {
      setUploadStatus(UploadStatus.LOADING);
      const formData = new FormData();
      formData.append("video", file.current);
      try {
        const response = await fetch(
          `http://localhost:5000/change-codec?inputfile=${fileName}`,
          {
            method: "POST",
            body: formData,
          },
        );
        if (response.ok) {
          const { data } = await response.json();
          console.log(data);
          setOutput(data["output_path"]);
          if (data["output_size"] && typeof data["output_size"] == "number") {
            setTimes(data["output_size"] / file.current.size);
          } else {
            alert(
              "output_size isn't a number, type: " +
                typeof data?.["output_size"] || "unknown",
            );
          }
          alert("Video uploaded successfully");
          setUploadStatus(UploadStatus.COMPRESSED);
        }
      } catch (e) {
        alert("Failed to upload video");
        console.error(e);
        file.current = undefined;
        setFileName("");
        setUploadStatus(UploadStatus.IDLE);
      }
    } else {
      setUploadStatus(UploadStatus.IDLE);
      alert("Please select a video file to upload");
    }
  };

  const handleDownload = async () => {
    setUploadStatus(UploadStatus.LOADING);
    await downloadFile(`http://localhost:5000/download?file=${output}`, output)
      .then(() => {
        alert("Download file successfully");
        file.current = undefined;
        setFileName("");
        setUploadStatus(UploadStatus.IDLE);
      })
      .catch((e) => {
        alert("Failed to download video");
        console.error(e);
        setUploadStatus(UploadStatus.COMPRESSED);
      });
  };

  const handleButton = async () => {
    switch (uploadStatus) {
      case UploadStatus.IDLE:
      case UploadStatus.UPLOADED:
        return handleCompress();
      case UploadStatus.COMPRESSED:
        return handleDownload();
      default:
        return () => {
          alert("poor handling of event");
        };
    }
  };
  return (
    <Card className="w-full max-w-xl mx-auto">
      <CardHeader>
        <CardTitle>Upload a Video</CardTitle>
        <CardDescription>
          Compress video size with no compromise on quality
        </CardDescription>
      </CardHeader>
      {uploadStatus === UploadStatus.COMPRESSED && (
        <CardContent className="flex items-center gap-2 text-success">
          {/**
           */}
          <CheckIcon className="w-10 h-10 px-1 py-1 text-green-500 bg-green-100 rounded-full" />
          <div>
            <p className="text-xl">
              <strong className="text-primary"> {times}x </strong> smaller in
              size{" "}
            </p>
            <p className="text-muted-foreground">
              Size compresssion completed successfully.{" "}
            </p>
          </div>
        </CardContent>
      )}
      {uploadStatus === UploadStatus.LOADING && (
        <CardContent className="grid grid-cols-[120px_1fr] gap-4">
          <div className="relative flex flex-col items-center justify-center gap-2 border-2 border-dashed border-muted rounded-md p-8 transition-colors hover:border-primary hover:bg-muted aspect-square">
            <Skeleton className="h-full w-full rounded-md" />
          </div>
          <div className="flex flex-col justify-center">
            <div>
              <Skeleton className="h-4 w-24" />
              <Skeleton className="h-3 w-32 mt-2" />
            </div>
          </div>
        </CardContent>
      )}
      {uploadStatus === UploadStatus.UPLOADED && (
        <UploadedFileCardContent
          fileName={fileName}
          thumbnail={thumbnail}
          fileSize={
            !!file.current && file.current?.size > 0
              ? formatFileSize(file.current.size)
              : "unknown video file size"
          }
        />
      )}

      {uploadStatus === UploadStatus.IDLE && (
        <UploadFileCardContent
          handleFileChange={handleFileChange}
          setFileName={setFileName}
          filename={fileName}
        />
      )}

      <CardFooter className="flex justify-end">
        <Button
          onClick={handleButton}
          type="submit"
          disabled={uploadStatus === UploadStatus.LOADING}
        >
          {
            {
              [UploadStatus.IDLE]: "Upload Video",
              [UploadStatus.LOADING]: "Uploading...",
              [UploadStatus.UPLOADED]: "Compress Video",
              [UploadStatus.COMPRESSED]: "Download Video",
            }[uploadStatus]
          }
        </Button>
      </CardFooter>
    </Card>
  );
}

function UploadedFileCardContent({
  fileName,
  thumbnail,
  fileSize,
}: {
  fileName: string;
  fileSize: string;
  thumbnail: any;
}) {
  return (
    <CardContent className="grid grid-cols-[120px_1fr] gap-4">
      <div className="relative flex flex-col items-center justify-center border-2 p-2 border-dashed border-muted rounded-md transition-colors hover:border-primary hover:bg-muted aspect-square">
        <img
          src={thumbnail || "/favicon.svg"}
          alt="Thumbnail"
          width={120}
          height={120}
          className="object-cover rounded-md"
          style={{ aspectRatio: "120/120", objectFit: "cover" }}
        />
      </div>
      <div className="flex flex-col justify-center">
        <div>
          <div className="font-medium">{fileName || "unknown video name"}</div>
          <div className="text-muted-foreground text-sm">{fileSize}</div>
        </div>
      </div>
    </CardContent>
  );
}
function UploadFileCardContent({
  handleFileChange,
  setFileName,
  filename,
}: {
  handleFileChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  setFileName: (_: string) => void;
  filename: string;
}) {
  return (
    <CardContent className="grid gap-4">
      <div className="relative flex flex-col items-center justify-center gap-2 border-2 border-dashed border-muted rounded-md p-8 transition-colors hover:border-primary hover:bg-muted">
        <UploadIcon className="w-8 h-8 text-muted-foreground" />
        <div className="text-center text-muted-foreground">
          <span className="mt-2">Drag and drop a video file here or</span>
          <Button
            variant="link"
            className="ml-0 pl-1 underline inline-flex text-md"
          >
            click to select
          </Button>
        </div>
        <input
          type="file"
          accept="video/*"
          onChange={handleFileChange}
          id="video-input"
          className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
        />
      </div>
      <div className="grid gap-2">
        <Label htmlFor="title">Title</Label>
        <Input
          id="name"
          placeholder="Change name"
          onChange={(e) => setFileName(e.target.value)}
          value={filename}
        />
      </div>
    </CardContent>
  );
}

function UploadIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
      <polyline points="17 8 12 3 7 8" />
      <line x1="12" x2="12" y1="3" y2="15" />
    </svg>
  );
}

function XIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>
  );
}
