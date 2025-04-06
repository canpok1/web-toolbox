import { CircleX } from "lucide-react";

function Alert({
  messages,
  className,
}: { messages: string[]; className?: string }) {
  if (messages.length === 0) {
    return null;
  }
  return (
    <div className={`flex flex-col gap-2 ${className}`}>
      {messages.map((message) => (
        <div key={message} role="alert" className="alert alert-error">
          <CircleX />
          {message}
        </div>
      ))}
    </div>
  );
}

export default Alert;
