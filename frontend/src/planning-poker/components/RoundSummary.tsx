import type { RoundParticipant } from "../types/Participant";
import VoteResult from "./VoteResult";

export type RoundSummaryProps = {
  round: number;
  participants: RoundParticipant[];
  revealed: boolean;
};

function RoundSummary({ round, participants, revealed }: RoundSummaryProps) {
  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">ラウンド{round}：投票結果</h2>
        <div className="grid grid-cols-5 gap-4">
          {participants.map((p) => (
            <VoteResult key={p.id} participant={p} revealed={revealed} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default RoundSummary;
