import { Check, Copy } from "lucide-react"; // アイコンをインポート
import { useState } from "react";
import type { Session } from "../types/Session";

function SessionSummary({ session }: { session: Session }) {
  const [isCopied, setIsCopied] = useState(false); // コピー状態を管理するstate
  const names = session.participants.map((p) => p.name);
  const joinPageUrl = new URL(
    // パスを修正: /planning-poker/sessions/join?id=...
    `/planning-poker/sessions/join?id=${session.id}`,
    window.location.origin,
  );
  const joinUrlString = joinPageUrl.toString();

  // クリップボードにコピーする関数
  const handleCopyClick = async () => {
    try {
      await navigator.clipboard.writeText(joinUrlString);
      setIsCopied(true);
      // 2秒後に表示を元に戻す
      setTimeout(() => setIsCopied(false), 2000);
    } catch (err) {
      console.error("クリップボードへのコピーに失敗しました:", err);
      // 必要であればユーザーにエラーメッセージを表示する処理を追加
    }
  };

  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">セッション名: {session.name}</h2>
        {/* 参加ページURLとコピーボタンを横並びにする */}
        <div className="flex items-center gap-2">
          <p className="flex-grow break-all">参加ページ: {joinUrlString}</p>
          <button
            type="button"
            className={`btn btn-sm ${isCopied ? "btn-success" : "btn-ghost"}`}
            onClick={handleCopyClick}
            aria-label="参加ページのURLをコピー"
            disabled={isCopied}
          >
            {isCopied ? <Check size={16} /> : <Copy size={16} />}
            {isCopied ? "コピー完了" : "コピー"}
          </button>
        </div>
        <p>参加者: {names.join(", ")}</p>
      </div>
    </div>
  );
}

export default SessionSummary;
