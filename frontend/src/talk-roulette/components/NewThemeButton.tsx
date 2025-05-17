import type React from "react";

interface Props {
  onClick: () => void;
}

const NewThemeButton: React.FC<Props> = ({ onClick }) => {
  return (
    <button
      type="button"
      id="new-theme-button"
      className="rounded-full bg-gradient-to-r from-blue-600 to-blue-700 px-6 py-3 font-semibold text-white shadow-md transition duration-300 ease-in-out hover:scale-105 hover:from-blue-700 hover:to-blue-800"
      onClick={onClick}
    >
      別のテーマを引く
    </button>
  );
};

export default NewThemeButton;
