import { LogIn } from "lucide-react";
import type React from "react";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { ApiClient } from "../api/ApiClient";

function CreateSessionPage() {
  const [userName, setUserName] = useState<string>("");
  const [scale, setScale] = useState<"fibonacci" | "t-shirt" | "power-of-two">(
    "fibonacci",
  );
  const navigate = useNavigate();

  const client = new ApiClient();

  const handleUserNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUserName(event.target.value);
  };

  const handleScaleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    if (
      event.target.value === "fibonacci" ||
      event.target.value === "t-shirt" ||
      event.target.value === "power-of-two"
    ) {
      setScale(event.target.value);
    }
  };

  const handleSubmit = async () => {
    try {
      console.log("clicked button, userName:%s, scale:%s", userName, scale);
      const resp = await client.createSession({
        sessionName: "xxx",
        hostName: userName,
        scaleType: "fibonacci",
      });
      navigate(`/planning-poker/sessions/${resp.sessionId}?id=${resp.hostId}`);
    } catch (error) {
      console.error(error);
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
          <label className="floating-label mx-auto mb-3 w-full">
            <span>名前</span>
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
