import { CircleCheckBig, CircleSlash } from "lucide-react";
import type { RoundParticipant } from "../types/Participant";

export type VoteResultProps = {
  participant: RoundParticipant;
  revealed: boolean;
};

function VoteResult({ participant, revealed }: VoteResultProps) {
  return (
    <div className="stats shadow">
      <div className="stat place-items-center">
        <div className="stat-title">{participant.name}</div>
        {revealed && <div className="stat-value">{participant.vote ?? ""}</div>}
        {!revealed && participant.vote && (
          <CircleCheckBig className="text-success-content" />
        )}
        {!revealed && !participant.vote && (
          <CircleSlash className="text-error" />
        )}
      </div>
    </div>
  );
}

export default VoteResult;
