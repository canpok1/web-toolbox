import type React from "react";

interface Props {
  genre: string;
  handleGenreChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
}

const GenreSelector: React.FC<Props> = ({ genre, handleGenreChange }) => {
  return (
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
  );
};

export default GenreSelector;
