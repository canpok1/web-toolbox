import { Link } from "react-router-dom";

const TopPage = () => {
  return (
    <div className="flex h-screen flex-col items-center justify-center">
      <h1 className="mb-8 font-bold text-4xl">Web Toolbox</h1>
      <div className="space-y-4">
        <Link
          to="/planning-poker"
          className="rounded bg-blue-500 px-4 py-2 font-bold text-white hover:bg-blue-700"
        >
          プランニングポーカー
        </Link>
      </div>
    </div>
  );
};

export default TopPage;
