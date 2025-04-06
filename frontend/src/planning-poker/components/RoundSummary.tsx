import type { RoundParticipant } from "../types/Participant";
import VoteResult from "./VoteResult";

export type RoundSummaryProps = {
  participants: RoundParticipant[];
  revealed: boolean;
};

function RoundSummary({ participants, revealed }: RoundSummaryProps) {
  const title = revealed ? "投票結果" : "投票状況";
  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">{title}</h2>
        <div className="grid grid-cols-3 gap-4 md:grid-cols-3 lg:grid-cols-6">
          {participants.map((p) => (
            <VoteResult key={p.id} participant={p} revealed={revealed} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default RoundSummary;
