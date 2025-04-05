import { LogIn, Users } from "lucide-react";
import { Link } from "react-router-dom";

function PlanningPokerTopPage() {
  return (
    <div className="mx-auto max-w-2xl py-25 px-5 text-center">
      <div className="mb-5">
        <h1 className="font-bold text-3xl">プランニングポーカー</h1>
      </div>
      <div className="mb-5">
        <Link
          to="./sessions/create"
          className="btn btn-primary w-full"
          aria-label="セッションを作成"
        >
          <LogIn />
          セッションを作成
        </Link>
      </div>
      <div className="justify-center">
        <Link
          to="./sessions/join"
          className="btn btn-primary w-full"
          aria-label="セッションに参加"
        >
          <Users />
          セッションに参加
        </Link>
      </div>
    </div>
  );
}

export default PlanningPokerTopPage;
