import { Check, Copy } from "lucide-react";
import { ChevronDown, ChevronUp } from "lucide-react";
import { QRCodeSVG } from "qrcode.react";
import { useState } from "react";
import Alert from "./Alert";
export type SessionInviteProps = {
  sessionId: string;
  className?: string;
};

export default function SessionInvitePanel({
  sessionId,
  className,
}: SessionInviteProps) {
  const [isCopied, setIsCopied] = useState(false);
  const [copyError, setCopyError] = useState<string | null>(null);
  const [isInviteSectionOpen, setIsInviteSectionOpen] = useState(false);

  const joinPageUrl = new URL(
    `/planning-poker/sessions/join?id=${sessionId}`,
    window.location.origin,
  );
  const joinPageUrlString = joinPageUrl.toString();

  const handleCopyClick = async () => {
    setCopyError(null);
    setIsCopied(false);
    try {
      await navigator.clipboard.writeText(joinPageUrlString);
      setIsCopied(true);
      setTimeout(() => setIsCopied(false), 2000);
    } catch (err) {
      console.error("クリップボードへのコピーに失敗しました:", err);
      setCopyError("クリップボードへのコピーに失敗しました。");
      setTimeout(() => setCopyError(null), 3000);
    }
  };

  const toggleInviteSection = () => {
    setIsInviteSectionOpen(!isInviteSectionOpen);
    if (!isInviteSectionOpen) {
      setCopyError(null);
      setIsCopied(false);
    }
  };

  return (
    <div className={`card mx-auto bg-neutral-content shadow-sm ${className}`}>
      <div className="card-body flex flex-col text-left">
        <Alert messages={copyError ? [copyError] : []} />
        <button
          type="button"
          className={`btn btn-sm shrink-0 ${isCopied ? "btn-success" : "btn-primary"}`}
          onClick={handleCopyClick}
          aria-label="参加ページのURLをコピー"
          disabled={isCopied}
        >
          {isCopied ? <Check size={16} /> : <Copy size={16} />}
          {isCopied ? "コピー完了" : "招待URLをコピー"}
        </button>

        <button
          type="button"
          onClick={toggleInviteSection}
          className="btn btn-ghost btn-sm w-full justify-between"
          aria-expanded={isInviteSectionOpen}
        >
          <span>招待URL/QRコード</span>
          {isInviteSectionOpen ? (
            <ChevronUp size={18} />
          ) : (
            <ChevronDown size={18} />
          )}
        </button>

        {isInviteSectionOpen && (
          <>
            <a
              className="break-all text-sm"
              href={joinPageUrlString}
              target="_blank"
              rel="noopener noreferrer"
            >
              {joinPageUrlString}
            </a>
            <div className="flex justify-center">
              <QRCodeSVG value={joinPageUrlString} />
            </div>
          </>
        )}
      </div>
    </div>
  );
}
