import { Check, ChevronDown, ChevronUp, Copy } from "lucide-react";
import { QRCodeSVG } from "qrcode.react";
import { useMemo, useState } from "react"; // useMemo をインポート
import type { Session } from "../types/Session";
import Alert from "./Alert";

export type SessionSummaryProps = {
  session: Session;
  currentParticipantId: string | null;
};

function SessionSummary({
  session,
  currentParticipantId,
}: SessionSummaryProps) {
  const [isCopied, setIsCopied] = useState(false);
  const [isInviteSectionOpen, setIsInviteSectionOpen] = useState(false);
  const [copyError, setCopyError] = useState<string | null>(null);

  const currentUserName = useMemo(() => {
    if (!currentParticipantId) {
      return null;
    }
    const currentUser = session.participants.find(
      (p) => p.id === currentParticipantId,
    );
    return currentUser ? currentUser.name : null;
  }, [session.participants, currentParticipantId]);

  const participantNames = session.participants.map((p) => p.name);
  const joinPageUrl = new URL(
    `/planning-poker/sessions/join?id=${session.sessionId}`,
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
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        {currentUserName && (
          <p className="font-semibold">あなたの名前: {currentUserName}</p>
        )}
        <p>
          参加者 ({participantNames.length}名): {participantNames.join(", ")}
        </p>
        <div>
          <button
            type="button"
            onClick={toggleInviteSection}
            className="btn btn-ghost btn-sm mt-2 flex w-full justify-between"
            aria-expanded={isInviteSectionOpen}
            aria-controls="invite-section"
          >
            <span>参加用URL/QRコードを表示</span>
            {isInviteSectionOpen ? (
              <ChevronUp size={18} />
            ) : (
              <ChevronDown size={18} />
            )}
          </button>

          {isInviteSectionOpen && (
            <div id="invite-section" className="mt-2 space-y-4">
              <Alert messages={copyError ? [copyError] : []} />

              <div className="flex items-center gap-2">
                <p className="flex-grow break-all text-sm">
                  {joinPageUrlString}
                </p>
                <button
                  type="button"
                  className={`btn btn-sm shrink-0 ${isCopied ? "btn-success" : "btn-ghost"}`}
                  onClick={handleCopyClick}
                  aria-label="参加ページのURLをコピー"
                  disabled={isCopied}
                >
                  {isCopied ? <Check size={16} /> : <Copy size={16} />}
                  {isCopied ? "コピー完了" : "コピー"}
                </button>
              </div>

              <div className="flex justify-center">
                <QRCodeSVG value={joinPageUrlString} />
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default SessionSummary;
