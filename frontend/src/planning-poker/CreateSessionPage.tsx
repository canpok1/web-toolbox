import { LogIn } from "lucide-react";
import type React from "react";
import { useState } from "react";
import { Link } from "react-router-dom";

function CreateSessionPage() {
  const [userName, setUserName] = useState<string>("");
  const [scale, setScale] = useState<string>("");

  const handleUserNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUserName(event.target.value);
  };

  const handleScaleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setScale(event.target.value);
  };

  const handleSubmit = () => {
    console.log("clicked button, userName:%s, scale:%s", userName, scale);
  };

  return (
    <div className="max-w-2xl mx-auto py-25 text-center">
      <div className="mb-5">
        <h1 className="text-3xl font-bold">プランニングポーカー</h1>
      </div>
      <div className="card max-w-2xl mx-auto shadow-sm mb-5">
        <div className="card-body text-left bg-neutral-content">
          <h2 className="card-title">セッションを作成</h2>
          <p className="mb-5">ホストとしてセッションを開始します。</p>
          <label className="input w-full mx-auto">
            <span className="名前">名前</span>
            <input
              type="text"
              placeholder="あなたの名前"
              value={userName}
              onChange={handleUserNameChange}
            />
          </label>
          <select
            className="select w-full mx-auto"
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
