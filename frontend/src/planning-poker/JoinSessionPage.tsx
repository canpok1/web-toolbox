import { Users } from "lucide-react";
import { type ChangeEvent, useState } from "react";
import { Link } from "react-router-dom";

function JoinSessionPage() {
  const [sessionId, setSessionId] = useState<string>("");
  const [userName, setUserName] = useState<string>("");

  const handleSessionIdChange = (event: ChangeEvent<HTMLInputElement>) => {
    setSessionId(event.target.value);
  };

  const handleUserNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    setUserName(event.target.value);
  };

  const handleSubmit = () => {
    console.log(
      "clicked button, sessionId:%s, userName:%s",
      sessionId,
      userName,
    );
  };

  return (
    <div className="mx-auto max-w-2xl px-5 py-25 text-center">
      <div className="mb-5">
        <h1 className="font-bold text-3xl">プランニングポーカー</h1>
      </div>
      <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
        <div className="card-body bg-neutral-content text-left">
          <h2 className="card-title">セッションに参加</h2>
          <p className="mb-5">既存のセッションに参加します。</p>
          <label className="floating-label mx-auto mb-3 w-full">
            <span>セッションID</span>
            <input
              className="input w-full"
              type="text"
              value={sessionId}
              placeholder="セッションID"
              onChange={handleSessionIdChange}
            />
          </label>
          <label className="floating-label mx-auto mb-3 w-full">
            <span>名前</span>
            <input
              className="input w-full"
              type="text"
              value={userName}
              placeholder="あなたの名前"
              onChange={handleUserNameChange}
            />
          </label>
          <button
            type="button"
            className="btn btn-primary w-full"
            aria-label="セッションに参加"
            onClick={handleSubmit}
          >
            <Users />
            セッションに参加
          </button>
        </div>
      </div>
      <Link
        to="/planning-poker"
        className="btn btn-secondary w-full"
        aria-label="戻る"
      >
        戻る
      </Link>
    </div>
  );
}

export default JoinSessionPage;
