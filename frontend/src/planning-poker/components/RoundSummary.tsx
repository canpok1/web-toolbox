import type { RoundParticipant } from "../types/Participant";

export type RoundSummaryProps = {
  round: number;
  participants: RoundParticipant[];
};

function RoundSummary({ round, participants }: RoundSummaryProps) {
  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">ラウンド{round}：投票結果</h2>
        <div className="grid grid-cols-5 gap-4">
          {participants.map((p) => (
            <div key={p.id} className="stat place-items-center">
              <div className="stat-title">{p.name}</div>
              <div className="stat-value">{p.vote}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default RoundSummary;
