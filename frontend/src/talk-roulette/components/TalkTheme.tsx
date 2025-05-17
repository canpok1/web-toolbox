import type React from "react";

interface Props {
  theme: string;
}

const TalkTheme: React.FC<Props> = ({ theme }) => {
  return (
    <div
      id="theme-display"
      className="mb-6 flex min-h-[3rem] items-center justify-center rounded-md bg-blue-50/50 p-4 font-semibold text-blue-800 text-xl transition-all duration-300"
    >
      {theme}
    </div>
  );
};

export default TalkTheme;
