import type { SessionParticipant } from "../types/Participant";
import type { RoundSummary } from "../types/RoundSummary";
import type { Vote } from "../types/Vote";
import VoteResult from "./VoteResult";

export type RoundSummaryProps = {
  participants: SessionParticipant[];
  votes: Vote[];
  revealed: boolean;
  summary?: RoundSummary;
  className?: string;
};

export default function RoundSummaryPanel({
  participants,
  votes,
  revealed,
  summary,
  className,
}: RoundSummaryProps) {
  const title = revealed ? "投票結果" : "投票状況";
  const voteMap = new Map(votes.map((v) => [v.participantId, v]));
  const hasVotes = votes.length > 0; // 投票があるかどうかを判定する変数

  const showSummary = revealed && summary;
  const showNoVote = revealed && !hasVotes;
  const showVote = !revealed || hasVotes;

  // 平均値と中央値を小数点以下1桁にフォーマットする関数 (必要に応じて調整)
  const formatNumber = (num: number | undefined): string => {
    if (num === undefined) return "-";
    return num.toFixed(1);
  };

  return (
    <div className={`card mx-auto shadow-sm ${className}`}>
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">{title}</h2>

        {/* 集計結果の表示 (revealed かつ summary が存在する場合) */}
        {showSummary && (
          <div className="mb-4 pb-4">
            <div className="stats stats-horizontal w-full shadow">
              {/* 平均値 */}
              <div className="stat">
                <div className="stat-title">平均値</div>
                <div className="stat-value text-primary">
                  {formatNumber(summary.average)}
                </div>
              </div>
              {/* 中央値 */}
              <div className="stat">
                <div className="stat-title">中央値</div>
                <div className="stat-value text-primary">
                  {" "}
                  {/* text-secondary から text-primary に変更 (他の値と合わせる) */}
                  {formatNumber(summary.median)}
                </div>
              </div>
              {/* 最大値 (summary に max があれば) */}
              {summary.max !== undefined && (
                <div className="stat">
                  <div className="stat-title">最大値</div>
                  <div className="stat-value text-primary">
                    {formatNumber(summary.max)}
                  </div>
                </div>
              )}
              {/* 最小値 (summary に min があれば) */}
              {summary.min !== undefined && (
                <div className="stat">
                  <div className="stat-title">最小値</div>
                  <div className="stat-value text-primary">
                    {formatNumber(summary.min)}
                  </div>
                </div>
              )}
            </div>
          </div>
        )}

        {/* 投票結果表示エリア */}
        {showNoVote && (
          // revealed が true で投票がない場合
          <div className="py-4 text-center text-gray-500">
            投票はありませんでした。
          </div>
        )}

        {showVote && (
          // revealed が false、または revealed が true で投票がある場合
          <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6">
            {participants.map((p) => (
              <VoteResult
                key={p.id}
                name={p.name}
                // revealed=true で投票がない参加者は VoteResult 側で null を返すので、
                // ここでは revealed=false の場合も考慮して voteMap.get(p.id) を渡す
                vote={voteMap.get(p.id)}
                revealed={revealed}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
