import type React from "react";
import { useCallback, useEffect, useState } from "react";
import { TalkRouletteClient } from "../api/TalkRouletteClient";
import FeedbackButtons from "./components/FeedbackButtons";
import GenreSelector from "./components/GenreSelector";
import NewThemeButton from "./components/NewThemeButton";
import NewThemeLink from "./components/NewThemeLink";
import TalkTheme from "./components/TalkTheme";

const talkRouletteClient = new TalkRouletteClient();

const TalkRouletteTopPage = () => {
  const [theme, setTheme] = useState("テーマを読み込み中...");
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [feedbackMessage, setFeedbackMessage] = useState("");
  const [genre, setGenre] = useState("all");

  // APIからテーマを取得する関数
  const fetchTheme = useCallback(async (selectedGenre: string) => {
    setTheme("テーマを読み込み中..."); // ロード中に表示
    setLiked(false);
    setDisliked(false);
    setFeedbackMessage("");
    try {
      const response = await talkRouletteClient.getThemes(
        selectedGenre === "all" ? undefined : selectedGenre,
        1, // 1件だけ取得
      );
      if (response.themes && response.themes.length > 0) {
        setTheme(response.themes[0].theme);
      } else {
        setTheme("テーマが見つかりませんでした。");
      }
    } catch (error) {
      console.error("Failed to fetch theme:", error);
      setTheme("テーマの取得に失敗しました。");
    }
  }, []);

  useEffect(() => {
    fetchTheme(genre);
  }, [genre, fetchTheme]);

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
      fetchTheme(genre); // テーマ変更時に新しいテーマを取得
    } else {
      setDisliked(false);
      setFeedbackMessage("");
    }
  };

  const handleGenreChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setGenre(e.target.value);
  };

  const handleNewThemeClick = () => {
    fetchTheme(genre); // 新しいテーマボタンクリック時に新しいテーマを取得
  };

  return (
    <div className="flex min-h-screen items-center justify-center p-4">
      <div className="w-full max-w-md rounded-lg bg-white p-8 text-center shadow-xl">
        <h1 className="mb-6 font-semibold text-2xl text-gray-700">
          今日のトークテーマ
        </h1>

        <TalkTheme theme={theme} data-testid="talk-theme" />

        <div className="mt-6 space-y-4">
          <FeedbackButtons
            liked={liked}
            disliked={disliked}
            handleLikeClick={handleLikeClick}
            handleDislikeClick={handleDislikeClick}
          />
          <p data-testid="feedback-message" className="text-gray-500 text-sm">
            {feedbackMessage}
          </p>

          <GenreSelector genre={genre} handleGenreChange={handleGenreChange} />

          <NewThemeButton onClick={handleNewThemeClick} />

          <NewThemeLink />
        </div>
      </div>
    </div>
  );
};

export default TalkRouletteTopPage;
