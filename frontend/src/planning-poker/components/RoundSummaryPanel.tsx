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
  const hasVotes = votes.length > 0;
  const hasSummary =
    summary?.average !== undefined &&
    summary?.median !== undefined &&
    summary?.max !== undefined &&
    summary?.min !== undefined;
  const hasCount = summary && summary.voteCounts.length > 0;

  const showSummary = revealed && hasSummary;
  const showCount = revealed && hasCount;
  const showNoVote = revealed && !hasVotes;
  const showVote = !revealed || hasVotes;

  // 平均値と中央値を小数点以下1桁にフォーマットする関数
  const formatNumber = (num: number | undefined): string => {
    if (num === undefined) return "-";
    return num.toFixed(1);
  };

  return (
    <div className={`card mx-auto shadow-sm ${className}`}>
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">{title}</h2>

        {showSummary && (
          <div className="mb-4 pb-4">
            <div className="stats stats-horizontal w-full shadow">
              <div className="stat place-items-center">
                <div className="stat-title">平均値</div>
                <div className="stat-value text-primary">
                  {formatNumber(summary?.average)}
                </div>
              </div>
              <div className="stat place-items-center">
                <div className="stat-title">中央値</div>
                <div className="stat-value text-primary">
                  {formatNumber(summary?.median)}
                </div>
              </div>
              <div className="stat place-items-center">
                <div className="stat-title">最大値</div>
                <div className="stat-value text-primary">
                  {formatNumber(summary?.max)}
                </div>
              </div>
              <div className="stat place-items-center">
                <div className="stat-title">最小値</div>
                <div className="stat-value text-primary">
                  {formatNumber(summary?.min)}
                </div>
              </div>
            </div>
          </div>
        )}

        {showCount && (
          <div className="mb-4 pb-4">
            <div className="stats stats-horizontal w-full shadow">
              {summary.voteCounts.map((voteCount) => (
                <div className="stat place-items-center" key={voteCount.value}>
                  <div className="stat-title">{voteCount.value}</div>
                  <div className="stat-value text-primary">
                    {voteCount.count}票
                  </div>
                  <div className="stat-desc max-w-sm whitespace-pre-line">
                    {voteCount.participants.map((p) => (
                      <span key={p.participantId} className="badge badge-sm">
                        {p.name}
                      </span>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}

        {showNoVote && (
          <div className="py-4 text-center text-gray-500">
            投票はありませんでした。
          </div>
        )}

        {showVote && (
          <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6">
            {participants.map((p) => (
              <VoteResult
                key={p.id}
                name={p.name}
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
