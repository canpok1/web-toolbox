import { LogIn, Users } from "lucide-react";

function TopPage() {
    return (
        <div className="max-w-200 mx-auto py-25 text-center">
            <div className="mb-5">
                <h1 className="text-3xl font-bold">プランニングポーカー</h1>
            </div>
            <div className="mb-5">
                <a href="./sessions/create" className="btn min-w-5/6">
                    <LogIn />セッションを作成
                </a>
            </div>
            <div className="justify-center">
                <a href="./sessions/join" className="btn min-w-5/6">
                    <Users />セッションを作成
                </a>
            </div>
        </div>
    )
}

export default TopPage;
