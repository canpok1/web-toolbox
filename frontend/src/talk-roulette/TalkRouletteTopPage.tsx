import type React from "react";
import { useEffect, useState } from "react";
import FeedbackButtons from "./components/FeedbackButtons";
import GenreSelector from "./components/GenreSelector";
import NewThemeButton from "./components/NewThemeButton";
import NewThemeLink from "./components/NewThemeLink";
import TalkTheme from "./components/TalkTheme";

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

    displayTheme();
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

  const handleNewThemeClick = () => {
    const genreThemes = themes[genre as keyof typeof themes] || themes.all;
    const randomIndex = Math.floor(Math.random() * genreThemes.length);
    setTheme(genreThemes[randomIndex]);
    setLiked(false);
    setDisliked(false);
    setFeedbackMessage("");
  };

  return (
    <div className="flex min-h-screen items-center justify-center p-4">
      <div className="w-full max-w-md rounded-lg bg-white p-8 text-center shadow-xl">
        <h1 className="mb-6 font-semibold text-2xl text-gray-700">
          今日のトークテーマ
        </h1>

        <TalkTheme theme={theme} />

        <div className="mt-6 space-y-4">
          <FeedbackButtons
            liked={liked}
            disliked={disliked}
            handleLikeClick={handleLikeClick}
            handleDislikeClick={handleDislikeClick}
          />
          <p id="feedback-message" className="text-gray-500 text-sm">
            {feedbackMessage}
          </p>

          <GenreSelector genre={genre} handleGenreChange={handleGenreChange} />

          <NewThemeButton onClick={handleNewThemeClick} />

          <NewThemeLink />
          <p className="mt-1 text-gray-500 text-xs">
            良いテーマを思いついたら、ぜひ投稿してください！
          </p>
        </div>
      </div>
    </div>
  );
};

export default TalkRouletteTopPage;
