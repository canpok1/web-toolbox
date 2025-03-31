import { LogIn, Users } from "lucide-react";
import { Link } from "react-router-dom";

function TopPage() {
    return (
        <div className="max-w-2xl mx-auto py-25 text-center">
            <div className="mb-5">
                <h1 className="text-3xl font-bold">プランニングポーカー</h1>
            </div>
            <div className="mb-5">
                <Link to="./sessions/create" className="btn min-w-5/6" aria-label="セッションを作成">
                    <LogIn />セッションを作成
                </Link>
            </div>
            <div className="justify-center">
                <Link to="./sessions/join" className="btn min-w-5/6" aria-label="セッションに参加">
                    <Users />セッションに参加
                </Link>
            </div>
        </div>
    )
}

export default TopPage;
