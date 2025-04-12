import { Check, ChevronDown, ChevronUp, Copy } from "lucide-react"; // Chevron アイコンを追加
import { useState } from "react";
import type { Session } from "../types/Session";

function SessionSummary({ session }: { session: Session }) {
  const [isCopied, setIsCopied] = useState(false);
  const [isInviteSectionOpen, setIsInviteSectionOpen] = useState(false); // 開閉状態を管理するstate
  const names = session.participants.map((p) => p.name);
  const joinPageUrl = new URL(
    `/planning-poker/sessions/join?id=${session.id}`,
    window.location.origin,
  );
  const joinUrlString = joinPageUrl.toString();

  const handleCopyClick = async () => {
    try {
      await navigator.clipboard.writeText(joinUrlString);
      setIsCopied(true);
      setTimeout(() => setIsCopied(false), 2000);
    } catch (err) {
      console.error("クリップボードへのコピーに失敗しました:", err);
    }
  };

  // 開閉状態を切り替える関数
  const toggleInviteSection = () => {
    setIsInviteSectionOpen(!isInviteSectionOpen);
  };

  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">セッション名: {session.name}</h2>
        <p className="mt-4">参加者: {names.join(", ")}</p>
        <div>
          <button
            type="button"
            onClick={toggleInviteSection}
            className="btn btn-ghost btn-sm flex w-full justify-between"
            aria-expanded={isInviteSectionOpen}
            aria-controls="invite-section"
          >
            <span>参加用URLを表示</span>
            {isInviteSectionOpen ? (
              <ChevronUp size={18} />
            ) : (
              <ChevronDown size={18} />
            )}
          </button>

          {/* 開閉されるコンテンツ */}
          {isInviteSectionOpen && (
            <div id="invite-section" className="mt-2 flex items-center gap-2">
              <p className="flex-grow break-all text-sm"> {joinUrlString}</p>
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
          )}
        </div>
      </div>
    </div>
  );
}

export default SessionSummary;
