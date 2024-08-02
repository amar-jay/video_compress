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

export default function Component() {
    const [fileName, setFileName] = useState('');
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target?.files?.[0];
    if (file && file.type.startsWith('video/')) {
      setFileName(file.name);
    } else {
      alert('Please select a valid video file.');
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
          <Input id="name" placeholder="Change name" />
        </div>
      </CardContent>
      <CardFooter className="flex justify-end">
        <Button type="submit">Upload Video</Button>
      </CardFooter>
    </Card>
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
