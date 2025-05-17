import type React from "react";
import { useEffect, useState } from "react";

const TalkRouletteTopPage = () => {
  const themes = {
    all: [
      "最近ハマっていることは？",
      "行ってみたい場所はどこですか？",
      "あなたの好きな映画は何ですか？",
      "子供の頃の夢は何でしたか？",
      "休日の過ごし方について教えてください。",
      "今までで一番嬉しかったことは？",
      "何か一つ、特技を教えてください。",
      "あなたの好きな食べ物は何ですか？",
      "尊敬する人はいますか？",
      "今のあなたを漢字一文字で表すと？",
      "タイムマシンがあったら、過去と未来どっちに行きたい？",
      "宝くじが当たったら何をしたい？",
      "宇宙人と友達になれると思いますか？",
      "もし透明人間になれたら、何をしたい？",
      "10年後の自分にメッセージを送るとしたら、何て書く？",
      "無人島に一つだけ持っていくなら何？",
      "超能力が使えるなら、どんな力が欲しい？",
      "学生時代の一番の思い出は何ですか？",
      "あなたのストレス解消法は？",
      "人生で一番影響を受けた本は？",
      "好きな音楽やアーティストは？",
      "ペットを飼っていますか？どんな動物？",
      "あなたのラッキーカラーは何色ですか？",
      "旅行に行くなら、どんなスタイルが好き？",
      "得意料理を教えてください。",
      "朝型？夜型？",
      "最近感動したことは？",
      "どんな時に幸せを感じますか？",
      "あなたの元気の源は？",
      "座右の銘はありますか？",
    ],
    general: [
      "最近、何か面白いことありました？",
      "最近気になっているニュースは？",
      "今の仕事を選んだ理由は？",
      "学生時代に熱中していたことは？",
      "最近買ったお気に入りのものは？",
      "最近挑戦したことは？",
      "最近嬉しかったことは？",
      "最近悲しかったことは？",
      "最近驚いたことは？",
      "最近考えたことは？",
    ],
    hobby: [
      "あなたの趣味は何ですか？",
      "その趣味を始めたきっかけは？",
      "趣味の魅力を教えてください。",
      "趣味を通して得られたものは？",
      "これから挑戦したい趣味は？",
      "趣味にかける時間は？",
      "趣味の道具でこだわりのものは？",
      "他の人におすすめしたい趣味は？",
      "趣味の集まりには参加してる？",
      "あなたの自慢の趣味グッズは？",
    ],
    food: [
      "好きな食べ物は何ですか？",
      "嫌いな食べ物は何ですか？",
      "得意料理を教えてください。",
      "最近作った料理は？",
      "よく行くお店はありますか？",
      "おすすめのレストランは？",
      "料理をする頻度は？",
      "お弁当は作りますか？",
      "外食する頻度は？",
      "料理でこだわっていることは？",
    ],
    travel: [
      "今までで一番思い出に残っている旅行は？",
      "一番良かった旅行先は？",
      "おすすめの旅行先は？",
      "旅行の計画は自分で立てる？",
      "旅行に行く頻度は？",
      "日帰り旅行で行くならどこ？",
      "国内旅行で行きたい場所は？",
      "海外旅行で行きたい場所は？",
      "旅行のスタイルは？",
      "旅行の必需品は？",
    ],
  };

  const [theme, setTheme] = useState("テーマを読み込み中...");
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [feedbackMessage, setFeedbackMessage] = useState("");
  const [genre, setGenre] = useState("all");

  useEffect(() => {
    const displayTheme = () => {
      const genreThemes = themes[genre as keyof typeof themes] || themes.all;
      const randomIndex = Math.floor(Math.random() * genreThemes.length);
      setTheme(genreThemes[randomIndex]);
      setLiked(false);
      setDisliked(false);
      setFeedbackMessage("");
    };

    const newThemeButton = document.getElementById("new-theme-button");
    newThemeButton?.addEventListener("click", displayTheme);

    displayTheme();

    return () => {
      newThemeButton?.removeEventListener("click", displayTheme);
    };
  }, [genre]);

  const handleLikeClick = () => {
    if (!liked) {
      setLiked(true);
      setDisliked(false);
      setFeedbackMessage("良いテーマですね！");
    } else {
      setLiked(false);
      setFeedbackMessage("");
    }
  };

  const handleDislikeClick = () => {
    if (!disliked) {
      setDisliked(true);
      setLiked(false);
      setFeedbackMessage("テーマを変更しますね。");
    } else {
      setDisliked(false);
      setFeedbackMessage("");
    }
  };

  const handleGenreChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setGenre(e.target.value);
  };

  return (
    <div className="flex min-h-screen items-center justify-center p-4">
      <div className="w-full max-w-md rounded-lg bg-white p-8 text-center shadow-xl">
        <h1 className="mb-6 font-semibold text-2xl text-gray-700">
          今日のトークテーマ
        </h1>

        <div
          id="theme-display"
          className="mb-6 flex min-h-[3rem] items-center justify-center rounded-md bg-blue-50/50 p-4 font-semibold text-blue-800 text-xl transition-all duration-300"
        >
          {theme}
        </div>

        <div className="mt-6 space-y-4">
          <div className="flex items-center justify-center gap-8">
            <button
              type="button"
              id="like-button"
              className={`feedback-button text-gray-600 hover:text-green-500 ${
                liked ? "liked" : ""
              }`}
              aria-label="良いテーマ"
              onClick={handleLikeClick}
            >
              <i className="fas fa-thumbs-up text-2xl" />
            </button>
            <button
              type="button"
              id="dislike-button"
              className={`feedback-button text-gray-600 hover:text-red-500 ${
                disliked ? "disliked" : ""
              }`}
              aria-label="悪いテーマ"
              onClick={handleDislikeClick}
            >
              <i className="fas fa-thumbs-down text-2xl" />
            </button>
          </div>
          <p id="feedback-message" className="text-gray-500 text-sm">
            {feedbackMessage}
          </p>

          <div>
            <label
              htmlFor="genre-select"
              className="mb-2 block font-bold text-gray-700 text-sm"
            >
              ジャンルを選択:
            </label>
            <select
              id="genre-select"
              className="w-full appearance-none rounded border px-4 py-3 text-gray-700 leading-tight shadow focus:shadow-outline focus:outline-none"
              value={genre}
              onChange={handleGenreChange}
            >
              <option value="all">すべて</option>
              <option value="general">一般</option>
              <option value="hobby">趣味</option>
              <option value="food">食べ物</option>
              <option value="travel">旅行</option>
            </select>
          </div>

          <button
            type="button"
            id="new-theme-button"
            className="rounded-full bg-gradient-to-r from-blue-600 to-blue-700 px-6 py-3 font-semibold text-white shadow-md transition duration-300 ease-in-out hover:scale-105 hover:from-blue-700 hover:to-blue-800"
          >
            別のテーマを引く
          </button>

          <a
            href="https://forms.gle/your-google-form-url"
            target="_blank"
            rel="noopener noreferrer"
            className="mt-2 block font-semibold text-blue-600 text-sm transition-colors duration-200 hover:text-blue-800"
          >
            新しいトークテーマを投稿する
          </a>
          <p className="mt-1 text-gray-500 text-xs">
            良いテーマを思いついたら、ぜひ投稿してください！
          </p>
        </div>
      </div>
    </div>
  );
};

export default TalkRouletteTopPage;
