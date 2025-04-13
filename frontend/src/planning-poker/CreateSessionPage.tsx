import { LogIn } from "lucide-react";
import { type ChangeEvent, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { ApiClient } from "../api/ApiClient";
import Alert from "./components/Alert";
import { isScaleType } from "./types/ScaleType";
import { ExtractErrorMessage } from "./utils/error";

function CreateSessionPage() {
  const [userName, setUserName] = useState<string>("");
  const [scale, setScale] = useState<string>("fibonacci");
  const [errorMessages, setErrorMessages] = useState<string[]>([]);

  const navigate = useNavigate();

  const client = new ApiClient();
  const shouldSubmit = userName !== "";

  const handleUserNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    setUserName(event.target.value);
  };

  const handleScaleChange = (event: ChangeEvent<HTMLSelectElement>) => {
    setScale(event.target.value);
  };

  const handleSubmit = async () => {
    try {
      if (!isScaleType(scale)) {
        console.error("invalid scale: %s", scale);
        return;
      }

      console.log("clicked button, userName:%s, scale:%s", userName, scale);
      const resp = await client.createSession({
        hostName: userName,
        scaleType: scale,
      });
      navigate(`/planning-poker/sessions/${resp.sessionId}?id=${resp.hostId}`);
    } catch (error) {
      console.error(error);
      setErrorMessages([ExtractErrorMessage(error)]);
    }
  };

  return (
    <div className="mx-auto max-w-2xl px-5 py-25 text-center">
      <div className="mb-5">
        <h1 className="font-bold text-3xl">プランニングポーカー</h1>
      </div>
      <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
        <div className="card-body bg-neutral-content text-left">
          <h2 className="card-title">セッションを作成</h2>
          <p className="mb-5">ホストとしてセッションを開始します。</p>
          <Alert messages={errorMessages} className="mb-3" />
          <label className="floating-label mx-auto mb-3 w-full">
            <span>あなたの名前</span>
            <input
              className="input w-full"
              type="text"
              placeholder="あなたの名前"
              value={userName}
              onChange={handleUserNameChange}
            />
          </label>
          <label className="floating-label">
            <span>見積スケール</span>
            <select
              className="select mx-auto mb-3 w-full"
              value={scale}
              onChange={handleScaleChange}
            >
              <option value="fibonacci">
                フィボナッチ（0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ?）
              </option>
              <option value="t-shirt">
                Tシャツサイズ(XS, S, M, L, XL, XXL, ?)
              </option>
              <option value="power-of-two">
                2の累乗(1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, ?)
              </option>
            </select>
          </label>
          <button
            type="button"
            className="btn btn-primary w-full"
            aria-label="セッションを作成"
            onClick={handleSubmit}
            disabled={!shouldSubmit}
          >
            <LogIn />
            セッションを作成
          </button>
        </div>
      </div>
      <Link
        to="/planning-poker"
        className="btn btn-secondary min-w-full"
        aria-label="戻る"
      >
        戻る
      </Link>
    </div>
  );
}

export default CreateSessionPage;
