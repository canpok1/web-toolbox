import { Users } from "lucide-react";
import { type ChangeEvent, useState } from "react";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import { ApiClient } from "../api/ApiClient";
import Alert from "./components/Alert";
import { ExtractErrorMessage } from "./utils/error";

function JoinSessionPage() {
  const [searchParams] = useSearchParams();
  const initialSessionId = searchParams.get("id") ?? "";
  const hasInitialSessionId = initialSessionId !== "";

  const [sessionId, setSessionId] = useState<string>(initialSessionId);
  const [userName, setUserName] = useState<string>("");
  const [errorMessages, setErrorMessages] = useState<string[]>([]);

  const navigate = useNavigate();

  const client = new ApiClient();
  const shouldSubmit = sessionId.trim() !== "" && userName.trim() !== "";

  const handleSessionIdChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (hasInitialSessionId) {
      return;
    }
    setSessionId(event.target.value);
  };

  const handleUserNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    setUserName(event.target.value);
  };

  const handleSubmit = async () => {
    setErrorMessages([]);

    try {
      console.log(
        "clicked button, sessionId:%s, userName:%s",
        sessionId,
        userName,
      );
      const resp = await client.joinSession(sessionId, {
        name: userName,
      });
      navigate(
        `/planning-poker/sessions/${sessionId}?id=${resp.participantId}`,
      );
    } catch (error) {
      console.error(error);
      setErrorMessages([ExtractErrorMessage(error)]);
    }
  };

  return (
    <div className="mx-auto max-w-2xl px-5 py-5 text-center">
      <div className="mb-5">
        <h1 className="font-bold text-3xl">プランニングポーカー</h1>
      </div>
      <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
        <div className="card-body bg-neutral-content text-left">
          <h2 className="card-title">セッションに参加</h2>
          <p className="mb-5">既存のセッションに参加します。</p>
          <Alert messages={errorMessages} className="mb-3" />
          <label className="floating-label mx-auto mb-3 w-full">
            <span>セッションID</span>
            <input
              readOnly={hasInitialSessionId}
              className={`input w-full ${hasInitialSessionId ? "input-disabled cursor-not-allowed bg-base-200 text-opacity-70" : ""}`}
              type="text"
              value={sessionId}
              placeholder="セッションID"
              onChange={handleSessionIdChange}
              aria-readonly={hasInitialSessionId}
              title={
                hasInitialSessionId ? "セッションIDは変更できません" : undefined
              }
            />
          </label>
          <label className="floating-label mx-auto mb-3 w-full">
            <span>名前</span>
            <input
              className="input w-full"
              type="text"
              value={userName}
              maxLength={10}
              placeholder="あなたの名前"
              aria-label="あなたの名前"
              onChange={handleUserNameChange}
            />
          </label>
          <button
            type="button"
            className="btn btn-primary w-full"
            aria-label="セッションに参加"
            onClick={handleSubmit}
            disabled={!shouldSubmit}
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
