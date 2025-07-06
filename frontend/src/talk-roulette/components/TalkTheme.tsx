import type React from "react";

interface Props {
  theme: string;
  "data-testid"?: string;
}

const TalkTheme: React.FC<Props> = ({ theme, "data-testid": dataTestId }) => {
  return (
    <div
      id="theme-display"
      className="mb-6 flex min-h-[3rem] items-center justify-center rounded-md bg-blue-50/50 p-4 font-semibold text-blue-800 text-xl transition-all duration-300"
      data-testid={dataTestId}
    >
      {theme}
    </div>
  );
};

export default TalkTheme;
